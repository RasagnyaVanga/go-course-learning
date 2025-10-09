package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	ch := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				log.Println("Worker stopped:", ctx.Err())
				return
			case <-time.After(1 * time.Second):
				log.Println("Working..", i)
			}
		}
		ch <- "work completed"
	}()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		if err == context.DeadlineExceeded {
			log.Println("Context timed out")
		} else if err == context.Canceled {
			log.Println("Context cancelled due to client")
		}
		fmt.Fprintf(w, "Received cancellation signal, stopping work: %v\n", ctx.Err())

	case res, ok := <-ch:
		if ok {
			log.Println(res)
			fmt.Fprintln(w, res)
		}
	}
}

func main() {
	http.HandleFunc("/process", process)

	fmt.Println("Server starting on 8080:")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server")
	}
}
