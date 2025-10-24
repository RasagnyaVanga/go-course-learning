package main

import (
	pb "chatapplication/pkg/chatapp"
	db "chatapplication/pkg/db"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

// ChatServer implements the gRPC ChatAppServer interface.
// It handles bidirectional streaming for a chatroom where multiple clients can join and exchange messages.
type ChatServer struct {
	pb.UnimplementedChatAppServer

	clientStreamData db.ClientStreamData //interface to access the map of users and their streams.
}

// NewChatServer creates an instance to the ChatServer struct
func NewChatServer(clientData db.ClientStreamData) *ChatServer {
	return &ChatServer{
		clientStreamData: clientData,
	}
}

// Chat implements a bidirectional streaming RPC that enables a real-time chatroom.
// Multiple clients can join and send messages concurrently.
// - When a client connects, the first message is used to register their username and add them to the chatroom.
// - When a client disconnects or an error occurs while receiving, the client is removed from the chatroom and a notification is broadcasted.
// - All incoming messages from clients are broadcasted to every connected client, including join/leave notifications.
func (s *ChatServer) Chat(stream pb.ChatApp_ChatServer) error {
	var username string
	for {
		message, err := stream.Recv()
		if err == io.EOF { //client disconnects
			s.clientStreamData.RemoveClient(username)
			leftmsg := &pb.ChatMessage{FromUser: "Server", Message: fmt.Sprintf("%s left the room", username)}
			s.clientStreamData.BroadcastMessage(leftmsg)
			return nil
		}
		if err != nil { //no message being received from client indicating client left.
			s.clientStreamData.RemoveClient(username)
			leftmsg := &pb.ChatMessage{FromUser: "Server", Message: fmt.Sprintf("%s left the room", username)}
			s.clientStreamData.BroadcastMessage(leftmsg)
			return err
		}

		if username == "" { //first message of that particular client
			username = message.FromUser
			s.clientStreamData.AddClient(message.FromUser, stream)

			joinmsg := &pb.ChatMessage{FromUser: "Server", Message: fmt.Sprintf("%s joined the room", username)}
			s.clientStreamData.BroadcastMessage(joinmsg)
			continue
		}
		err1 := s.clientStreamData.BroadcastMessage(message)
		if err1 != nil {
			return err1
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen ")
	}

	grpcServer := grpc.NewServer()
	clientData := db.NewClientStreamData() //interface returned through db
	chatServer := NewChatServer(clientData)
	pb.RegisterChatAppServer(grpcServer, chatServer)

	log.Println("Server listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
