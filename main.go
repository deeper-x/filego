package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// FileManager with core behaviours
type FileManager interface {
        Open(int) error
	Append(string) (bool, error)
        ReadContent() []string
        WriteLines([]string)
	Close() error
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
		Path: path,
		Mode: 0644,
	}

	return handler
}

// Open file resource
func (rh *ResHandler) Open(dsc int) error {
	f, err := os.OpenFile(rh.Path, dsc, rh.Mode)

	if err != nil {
		log.Println(err)
		return err
	}

	rh.Resource = f

	return nil
}

// Close the file handler
func (rh ResHandler) Close() error {
	err := rh.Resource.Close()

	return err
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

// ReadContent read, line by line, file content
func (rh *ResHandler) ReadContent() []string {
	var res []string

	rh.Descriptor = os.O_RDONLY
	scanner := bufio.NewScanner(rh.Resource)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		log.Println(scanner.Text())
		res = append(res, scanner.Text())
	}

	return res
}

// GetOtherLines erase line containing substring
func (rh *ResHandler) GetOtherLines(lines []string, content string) []string {
	var res []string

	for _, v := range lines {
		if !strings.Contains(v, content) {
			res = append(res, v)
		}
	}

	return res
}

// WriteLines write file line by line
func (rh *ResHandler) WriteLines(lines []string) {
	dw := bufio.NewWriter(rh.Resource)

	for _, v := range lines {
		_, _ = dw.WriteString(v + "\n")
	}

	dw.Flush()
}

func main() {
	o := New("./assets/file_demo.txt")

	// 1 - append
	o.Open(os.O_APPEND | os.O_CREATE | os.O_WRONLY)

	_, err := o.Append("string one")
	if err != nil {
		log.Println(err)
	}

	_, err = o.Append("string two")
	if err != nil {
		log.Println(err)
	}

	o.Close()

	// 2 - read content
	o.Open(os.O_RDONLY)
	lines := o.ReadContent()
	o.Close()

	log.Println(lines)

	// 3 - write lines
	o.Open(os.O_APPEND | os.O_CREATE | os.O_WRONLY)
	o.WriteLines([]string{"one", "two", "three", "four"})
	o.Close()

	// 4 - delete line and print content
	o.Open(os.O_RDONLY)
	lines = o.GetOtherLines(lines, "two")
	o.WriteLines(lines)
	o.Close()

	o.Open(os.O_RDONLY)
	lines = o.ReadContent()
	o.Close()

}
