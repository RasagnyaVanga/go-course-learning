package main

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

type CollectiveUserData struct {
	Users []User `json:"users"`
}
