# creating a http server
* created a HTTP server using "net/http" package.
* Started a basic server using http.ListenAndServe(), which listens for incoming requests.
* Defined a handler function, func process(w http.ResponseWriter, r *http.Request).
* Registered the handler to a route using, http.HandleFunc("/process",process).

# managing go routines - using context 
* Each incoming request automatically carries a context accessible via r.Context().
* Added a cancel function to r.Context() with context.WithTimeout and added a duration of 5 seconds.
* Now in go function there is a long task running for 10 seconds. 
- in which we included the two cases,
  - work progress (checking whether the go running for every second).
  - context cancellation via <- ctx.Done().
- if the work completed for 10 seconds, add the workcompleted status to an unbuffered channel.

* Back in the main part of the function process, used a select statement to determine the outcome, 
  - if <-ctx.Done() is triggered, log the cancellation reason using ctx.Err(): 
   - context.Canceled - client disconnected.
   - context.DeadlineExceeded - timeout exceeded.
  - the other case is where we check for the successful completion from the unbuffered channel.