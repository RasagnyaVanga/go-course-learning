package main

import (
	"context"
	"fmt"
	"gRPC_filetransfer/pkg/config"
	"gRPC_filetransfer/pkg/file"
	pb "gRPC_filetransfer/pkg/proto"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientService struct {
	cfg *config.Config
}

func NewClientService(cfg *config.Config) *ClientService {
	return &ClientService{cfg: cfg}
}

// uploadFile is a client-side function responsible for uploading a file to the server via a gRPC client-streaming RPC.
// It reads the specified file from disk in fixed-size chunks, sends each chunk to the server through the stream,
// and finally waits for the server’s confirmation once the upload is complete.
func (c *ClientService) uploadFile(ctx context.Context, client pb.FileServiceClient, filePath string, batchSize int) error {

	stream, err := client.Upload(ctx) //RPC call from client
	if err != nil {
		return fmt.Errorf("could not start upload: %v", err)
	}
	defer stream.CloseSend()

	file, err := os.Open(filePath) //opening file
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	buf := make([]byte, batchSize) //dividing into chunks of size batchSize
	batchNumber := 1

	for {
		n, err := file.Read(buf) //number of bytes actually read
		if err == io.EOF {       //reached end of file, nothing more to read from the stream, break out of loop.
			break
		}
		if err != nil { //unexpected error.
			return fmt.Errorf("error reading file: %v", err)
		}

		req := &pb.FileUploadRequest{
			FileName: filePath,
			Chunk:    buf[:n],
		}

		if err := stream.Send(req); err != nil {
			return fmt.Errorf("error sending chunk: %v", err)
		}

		log.Printf("Sent batch #%d (%d bytes)", batchNumber, n)
		batchNumber++
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to receive response: %v", err)
	}

	log.Printf("Upload complete: %s (%d bytes)", res.GetFileName(), res.GetSize())
	return nil
}

// downloadFile is a client-side function that handles downloading a file from the server using a streaming RPC.
// The client initiates the download request with the specified file name,
// then continuously receives file chunks from the server stream.
// The chunks are written sequentially to a local file created in the downloads directory
// until the entire file has been received.
func (c *ClientService) downloadFile(ctx context.Context, client pb.FileServiceClient, fileName string) error {

	file := file.NewFile()

	stream, err := client.Download(ctx, &pb.DownloadRequest{FileName: fileName}) //RPC call from client
	if err != nil {
		return fmt.Errorf("could not start download: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error receiving chunk: %v", err)
		}

		if file.FilePath == "" { //creating a new file instancefor downloading file.
			file.SetFile(fileName, c.cfg.FilesStorage.DownloadsLocation) //filepath

			defer file.Close()

			fmt.Println("Downloading file to: ", file.FilePath)
		}

		err = file.Write(resp.Chunk)
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
		fmt.Printf("received a chunk with size: %d\n", len(resp.Chunk))
	}
	fmt.Printf("Downloaded file: %s at %s", fileName, file.FilePath)
	return nil

}
func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	clientService := NewClientService(cfg)
	address := "localhost" + cfg.Server.Port
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	client := pb.NewFileServiceClient(conn)

	fmt.Println("Enter 1 to upload files")
	fmt.Println("Enter 2 to download files")
	fmt.Println("Enter 3 to get metadata from a directory")
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		fmt.Println("Enter path of file that you want to upload:")
		var filePath string
		fmt.Scan(&filePath)
		batchSize := 1024 * 1024 //1 MB

		if err := clientService.uploadFile(ctx, client, filePath, batchSize); err != nil {
			log.Fatalf("upload failed: %v", err)
		}
	case 2:
		fmt.Println("Enter name of the file that you want to download:")
		var fileName string
		fmt.Scan(&fileName)
		if err := clientService.downloadFile(ctx, client, fileName); err != nil {
			log.Fatalf("download failed: %v", err)
		}
	case 3:
		fmt.Println("Enter the stored folder path ")
		var folderPath string
		fmt.Scan(&folderPath)
		res, err := client.GetMetaData(ctx, &pb.MetaDataRequest{FolderPath: folderPath})
		if err != nil {
			log.Fatalf("getting meta data failed:%v", err)
		}
		for _, result := range res.FilesMetaData {
			fmt.Println(result)
		}
	}

}
