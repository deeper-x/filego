package main

import (
	"fmt"
	"log"
	"os"
)

// FileManager with core behaviours
type FileManager interface {
	Append(string) (bool, error)
	Close()
}

// ResHandler is the file handler object
type ResHandler struct {
	Path       string
	Resource   *os.File
	Descriptor int
	Mode       os.FileMode
}

// New is the ResHandler builder
func New(path string) ResHandler {
	handler := ResHandler{
		Path:       path,
		Descriptor: os.O_APPEND | os.O_CREATE | os.O_WRONLY,
		Mode:       0644,
	}

	return handler
}

// Open file resource
func (rh *ResHandler) Open() error {
	f, err := os.OpenFile(rh.Path, rh.Descriptor, rh.Mode)

	if err != nil {
		log.Println(err)
		return err
	}

	rh.Resource = f

	return nil
}

// Append line to file
func (rh ResHandler) Append(toWrite string) (bool, error) {
	fullStr := fmt.Sprintf("%v\n", toWrite)
	_, err := rh.Resource.WriteString(fullStr)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}

// Close the file handler
func (rh ResHandler) Close() {
	rh.Resource.Close()
}

func main() {
	o := New("./assets/file_demo.txt")

	o.Open()
	defer o.Close()

	_, err := o.Append("string one")
	if err != nil {
		log.Println(err)
	}

	_, err = o.Append("string two")
	if err != nil {
		log.Println(err)
	}
}
