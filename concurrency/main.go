package main

import (
	"fmt"
	"sync"
)

func main() {
	url := "https://dummyjson.com/users"
	var userIds []int
	res, err := getUserId("https://dummyjson.com/users?limit=10")
	if err != nil {
		fmt.Println("Error in getting user ids: ", err)
		return
	}
	userIds = res //userIds slice

	for _, id := range userIds {
		fmt.Println(id)
	}

	ch := make(chan UserEndPoint) //channel to store user structs
	sem := make(chan struct{}, 5) //using semaphore to limit simultaneous requests to 5

	var wg sync.WaitGroup
	wg.Add(len(userIds))
	for _, id := range userIds {

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{} //acquiring slot
			defer func() {
				<-sem //releasing slot
			}()

			if err := getUserEndpoint(ch, url, id); err != nil {
				fmt.Println("Error fetching user:", err)
			}
		}(id)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for user := range ch {
		fmt.Println(user)
	}
}
