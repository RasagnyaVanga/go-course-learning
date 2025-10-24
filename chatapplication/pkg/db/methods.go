package db

import (
	pb "chatapplication/pkg/chatapp"
	"log"
)

func (u *userStreamData) AddClient(username string, stream pb.ChatApp_ChatServer) {
	u.mu.Lock() //multiple clients can access this and can try to modify the map simultaneously - avoid it using locks
	defer u.mu.Unlock()
	u.Clients[username] = stream
	log.Printf("%s joined the chat", username)
}

func (u *userStreamData) RemoveClient(username string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	delete(u.Clients, username)
	log.Printf("%s left the chat", username)
}

// broadcasting a message to all online clients in chatroom
func (u *userStreamData) BroadcastMessage(message *pb.ChatMessage) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	//sending the message to every client's stream
	for _, stream := range u.Clients {
		if err := stream.Send(message); err != nil { //each client has its own stream, we are sending the message to each client to its stream.
			return err
		}
	}
	return nil
}
