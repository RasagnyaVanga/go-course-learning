package db

import (
	pb "chatapplication/pkg/chatapp"
	"sync"
)

type ClientStreamData interface {
	AddClient(username string, stream pb.ChatApp_ChatServer)
	RemoveClient(username string)
	BroadcastMessage(message *pb.ChatMessage) error
}
type userStreamData struct {
	Clients map[string]pb.ChatApp_ChatServer //each user- their streams //to store the connected clients
	mu      sync.Mutex
}

func NewClientStreamData() ClientStreamData {
	return &userStreamData{
		Clients: make(map[string]pb.ChatApp_ChatServer),
	}
}
