package db

import (
	pb "chatapplication/pkg/chatapp"
	"sync"
)

// ClientStreamData defines the interface for managing active chat clients.
// It provides methods to add/remove clients and broadcast messages to all connected clients.
type ClientStreamData interface {
	AddClient(username string, stream pb.ChatApp_ChatServer)
	RemoveClient(username string)
	BroadcastMessage(message *pb.ChatMessage) error
}

// userStreamData is a concrete implementation of ClientStreamData.
// It stores active clients in a map and uses a mutex to ensure thread-safe access.
type userStreamData struct {
	Clients map[string]pb.ChatApp_ChatServer // Maps usernames to their corresponding gRPC streams
	mu      sync.Mutex
}

// NewClientStreamData creates and returns a new instance of userStreamData with an initialized Clients map.
func NewClientStreamData() ClientStreamData {
	return &userStreamData{
		Clients: make(map[string]pb.ChatApp_ChatServer),
	}
}
