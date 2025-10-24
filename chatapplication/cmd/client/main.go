package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	pb "chatapplication/pkg/chatapp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatAppClient(conn)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	var username string
	fmt.Scanln(&username)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("client.Chat failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.Chat failed: %v", err)
			}
			fmt.Printf("%s: %s\n", msg.FromUser, msg.Message)
		}
	}()
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		stream.Send(&pb.ChatMessage{FromUser: username, Message: text})
	}
	<-waitc

}
