package file

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// File struct represents a file being uplaoded and written to disk.
// It holds the destination path of the file, buffer, and pointer to the actual output file on disk.
type File struct {
	FilePath   string
	OutputFile *os.File
	reader     io.Reader
}

func NewFile() *File {
	return &File{}
}

// SetFile prepares the file for writing.
// It takes the uploaded file’s name and the storage directory path from config.
// It ensures the directory exists, then creates the actual file on disk.
func (f *File) SetFile(fileName, path string) error {
	err := os.MkdirAll(path, os.ModePerm) //create directories if they don't exist
	if err != nil {
		log.Fatal(err)
	}
	f.FilePath = filepath.Join(path, fileName) //joining directory path and file name to get full file path
	fmt.Println("Saving file to:", f.FilePath) // debug
	file, err := os.Create(f.FilePath)         //creates the file and overwrites if already exists
	if err != nil {
		return err
	}
	f.OutputFile = file //assigning to the struct, as the file can be written later.
	return nil
}

// Write writes a chunk of bytes (from gRPC stream) to the file.
// It’s called repeatedly as chunks arrive from the client.
func (f *File) Write(chunk []byte) error {
	if f.OutputFile == nil { //ensure output file is open
		return nil
	}
	_, err := f.OutputFile.Write(chunk) //write chunk to file.
	return err
}

func (f *File) Close() error {
	return f.OutputFile.Close()
}

func (f *File) Read(chunk []byte) (n int, err error) {
	return f.reader.Read(chunk)
}
