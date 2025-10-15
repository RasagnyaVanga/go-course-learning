package main

//models used to store the user information and get user ids -

type User struct {
	ID int `json:"id"`
}

type CollectiveUserData struct {
	Users []User `json:"users"`
}

//models used to store the final user informtion fetched based on the userids -

type UserEndPoint struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}
