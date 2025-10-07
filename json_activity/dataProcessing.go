package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name  string `json:"firstName"`
	Age   int
	Email string //`json:"email"`
	Image string `json:"image"`
}

type CollectiveUserData struct {
	Users []User
}

func fetchJSON(url string) ([]byte, error) {
	res, err := http.Get(url) //fetching data from json
	if err != nil {
		return nil, errors.New("error fetching data from api ")
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body) //body -> bytes
	if err != nil {
		return nil, errors.New("error reading byte data from response")
	}
	return bytes, nil
}

func encode(url string) (string, error) {
	imageresp, err := http.Get(url) //download image
	if err != nil {
		return "", errors.New("error downloading the image")
	}
	defer imageresp.Body.Close()

	imageBytes, err := io.ReadAll(imageresp.Body) //body -> bytes
	if err != nil {
		return "", errors.New("error getting the image bytes")
	}

	base64Encoded := base64.StdEncoding.EncodeToString(imageBytes) //bytes -> base64 string
	return base64Encoded, nil
}

func main() {

	bytes, err := fetchJSON("https://dummyjson.com/users?limit=100")
	if err != nil {
		panic(err)
	}

	var collectiveUserData CollectiveUserData

	if err := json.Unmarshal(bytes, &collectiveUserData); err != nil { //json -> struct
		fmt.Println("Error unmarshalling the data", err)
		os.Exit(1)
	}

	for i := range collectiveUserData.Users {
		user := &collectiveUserData.Users[i]
		encodedString, err := encode(user.Image)
		if err != nil {
			panic(err)
		}
		user.Image = encodedString
	}
	resultJSON, err := json.MarshalIndent(collectiveUserData, "", " ") //struct -> json
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("file_json_images", resultJSON, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("File updated successfully")

}
