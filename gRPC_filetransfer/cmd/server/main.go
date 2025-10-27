package main

import (
	"context"
	"fmt"
	config "gRPC_filetransfer/pkg/config"
	"gRPC_filetransfer/pkg/file"
	pb "gRPC_filetransfer/pkg/proto"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedFileServiceServer
	cfg *config.Config
}

// Upload is a client-streaming RPC method that receives file chunks sent by the client.
// All received chunks belong to a single file. The server creates the file in the directory
// specified in the configuration and writes incoming chunks to it sequentially.
func (s *Server) Upload(stream pb.FileService_UploadServer) error {
	file := file.NewFile() //creating a new file instance
	var totalSize uint32

	for { //receiving chunks from client through stream
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if file.FilePath == "" { //if its path is not set i.e, file is being uploaded for the first time.,
			safeFileName := filepath.Base(req.GetFileName()) // only hello-world.txt - getting the base filename from the local path
			file.SetFile(safeFileName, s.cfg.FilesStorage.UploadsLocation)
			defer func() {
				if err := file.OutputFile.Close(); err != nil {
					fmt.Println("failed to close file:", err)
				}
			}()
		}

		chunk := req.GetChunk()
		if err := file.Write(chunk); err != nil { //writing the chunks to file on disk
			return err
		}

		totalSize += uint32(len(chunk))
		fmt.Printf("received a chunk with size: %d\n", len(chunk))
	}

	fileName := filepath.Base(file.FilePath)

	fmt.Printf("Uploaded file: %s, size: %d\n", fileName, totalSize)

	return stream.SendAndClose(&pb.FileUploadResponse{FileName: fileName, Size: totalSize})
}

// Download is a server-streaming RPC method that handles file download requests from clients.
// The client initiates the download by sending a file name to the server.
// The server locates the requested file in its uploads directory, reads it in fixed-size chunks,
// and streams those chunks sequentially back to the client until the entire file is sent.
func (s *Server) Download(req *pb.DownloadRequest, stream pb.FileService_DownloadServer) error {
	fileName := req.GetFileName()
	if fileName == "" {
		return status.Error(codes.InvalidArgument, "file name is required")
	}
	//checking if the file is present in uploads directory
	filePath := filepath.Join(s.cfg.FilesStorage.UploadsLocation, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist on server:\n %v", err)
		}
		fmt.Println("Failed to open file:")
		return err
	}
	defer file.Close()

	batchSize := 1024 * 1024       //1 MB
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

		if err := stream.Send(&pb.DownloadResponse{Chunk: buf[:n]}); err != nil { //server sending the file from server to client as a response for its download request
			return fmt.Errorf("error sending chunk: %v", err)
		}

		log.Printf("Sent batch #%d (%d bytes)", batchNumber, n)
		batchNumber++
	}
	return nil
}

// GetMetaData is a unary RPC method that retrieves metadata for all files in a given directory.
// The server reads all file entries from the folder path provided in the request,
// extracts details such as file name, size, and last modification time for each file,
// and returns them as a list in the response.
func (s *Server) GetMetaData(ctx context.Context, req *pb.MetaDataRequest) (*pb.MetaDataResponse, error) {
	files, err := os.ReadDir(req.GetFolderPath()) //reading all the file entries from the folder
	var filesMetaData []*pb.MetaData              //slice to store metadata of each file in directory
	if err != nil {
		return nil, fmt.Errorf("error while reading files in the path")
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, fmt.Errorf("error while retrieving the info")
		}
		filesMetaData = append(filesMetaData, &pb.MetaData{
			FileName:         info.Name(),
			Size:             info.Size(),
			ModificationTime: info.ModTime().String(),
		})
	}
	return &pb.MetaDataResponse{FilesMetaData: filesMetaData}, nil
}
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Listen on the configured gRPC port
	lis, err := net.Listen("tcp", cfg.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", cfg.Server.Port, err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register FileService with the server
	fileServer := &Server{cfg: cfg}
	pb.RegisterFileServiceServer(grpcServer, fileServer)

	fmt.Printf("gRPC server listening on %s\n", cfg.Server.Port)
	fmt.Printf("Uploaded files will be saved in server at: %s\n", cfg.FilesStorage.UploadsLocation)
	fmt.Printf("Downloaded files will be saved in client at: %s\n", cfg.FilesStorage.DownloadsLocation)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
