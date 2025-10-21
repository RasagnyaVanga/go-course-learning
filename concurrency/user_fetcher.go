package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// fetchJSON sends an HTTP GET request to the specified URL,
// reads the response body, and returns the raw JSON data as bytes.
// It returns an error if the request fails, the response status is not 200 OK,
// or reading the response body encounters an error.
func fetchJSON(url string) ([]byte, error) {
	res, err := http.Get(url) //fetching data from json
	if err != nil {
		return nil, errors.New("error fetching data from api ")
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response: %d", res.StatusCode)
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body) //body -> bytes
	if err != nil {
		return nil, errors.New("error reading byte data from response")
	}
	return bytes, nil
}

// getUserId fetches JSON data from the given URL,
// parses it, and returns a slice of user IDs.
// It returns an error if fetching the data or unmarshalling the JSON fails.
func getUserId(url string) ([]int, error) {
	bytes, err := fetchJSON(url)
	if err != nil {
		return nil, err
	}

	var data CollectiveUserData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	var userIds []int
	for _, user := range data.Users {
		userIds = append(userIds, user.ID)
	}
	return userIds, nil
}

// getUserEndpoint fetches JSON data for a single user from the specified URL using the given user ID,
// unmarshals it into a UserEndPoint struct, and sends the result to the provided channel.
// It returns an error if fetching the data or unmarshalling the JSON fails.
func getUser(ch chan<- User, url string, id int) error {

	bytes, err := fetchJSON(fmt.Sprintf("%s/%d", url, id))
	if err != nil {
		return err
	}

	var user User
	if err := json.Unmarshal(bytes, &user); err != nil {
		return err
	}
	ch <- user
	return nil

}
