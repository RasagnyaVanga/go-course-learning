 Main routine workflow:
 - Using `fetchJSON(url string) ([]byte, error)` and `getUserId(url string) ([]int, error)`:
     - Fetches JSON data from a URL and unmarshals it into the CollectiveUserData struct.
     - Extracts user IDs from the struct and stores them in a slice of integers.

 - Using `getUser(ch chan<- User, url string, id int) error`:
     - For each user ID, concurrently fetches detailed user data from the API.
     - Unmarshals the data into the User struct and sends it to a channel.

 - Consumes the user data from the channel in the main routine.

 - Limits the number of concurrent HTTP requests to 5 using a semaphore to control goroutine execution.