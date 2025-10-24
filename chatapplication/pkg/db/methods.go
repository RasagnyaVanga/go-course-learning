package db

import (
	pb "chatapplication/pkg/chatapp"
	"log"
)

// AddClient adds a new client to the chatroom.
// It stores the client's gRPC stream in the Clients map using the username as the key.
// Mutex is used to ensure thread-safe access because multiple clients may join simultaneously.
func (u *userStreamData) AddClient(username string, stream pb.ChatApp_ChatServer) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Clients[username] = stream
	log.Printf("%s joined the chat", username)
}

// RemoveClient removes a client from the chatroom.
// The client's entry is deleted from the Clients map safely using a mutex.
func (u *userStreamData) RemoveClient(username string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	delete(u.Clients, username)
	log.Printf("%s left the chat", username)
}

// BroadcastMessage sends a message to all currently connected clients in the chatroom.
// Returns an error if sending to any client fails.
func (u *userStreamData) BroadcastMessage(message *pb.ChatMessage) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	for _, stream := range u.Clients {
		if err := stream.Send(message); err != nil { //Each client has its own stream, we are sending the message to each client to its stream.
			return err
		}
	}
	return nil
}
