package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// process handles HTTP requests to the "/process" endpoint
func process(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second) //timeout context of 5 seconds duration
	defer cancel()

	ch := make(chan string) //channel to send result

	go func() {
		for i := 1; i <= 10; i++ {
			select {
			// context is done (timeout or client cancel)
			case <-ctx.Done():
				log.Println("Worker stopped:", ctx.Err())
				return
			// Simulate work by waiting 1 second per iteration
			case <-time.After(1 * time.Second):
				log.Println("Working..", i)
			}
		}
		ch <- "work completed" //sending result to channel after work completes
	}()

	// Main select waits either for the worker result or context cancellation
	select {
	// If context is done before worker finishes
	case <-ctx.Done():
		err := ctx.Err()
		if err == context.DeadlineExceeded {
			log.Println("Context timed out") // Timeout reached
		} else if err == context.Canceled {
			log.Println("Context cancelled due to client") // Client disconnected
		}
		fmt.Fprintf(w, "Received cancellation signal, stopping work: %v\n", ctx.Err())

	// If worker finishes before timeout
	case res, ok := <-ch:
		if ok {
			log.Println(res)
			fmt.Fprintln(w, res)
		}
	}
}

func main() {
	// Register HTTP handler for "/process" endpoint
	http.HandleFunc("/process", process)
	fmt.Println("Server starting on 8080:")

	// Start HTTP server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server")
	}
}
