package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var filePath string
	fmt.Println("Enter the file path: ")
	fmt.Scanf("%s", &filePath)

	fileContent, err := os.ReadFile(filePath) //opens file, reads all file contents into memory of bytes, closes file when done

	//instead of os.ReadFile, we can use os.Open() (opens the file),
	//    then read the file with file.Read()(bytes) or bufio.NewScanner()(line by line) or io.ReadAll()(whole file)
	//    then manually close the file os.Close()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist or Given a wrong path")
			return
		}
		fmt.Println("Error During File Read")
		return
	}
	fmt.Println(string(fileContent))
}
