package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

}

// FileManager with core behaviours
type FileManager interface {
	open(int) error
	close() error
	WriteLine(string) error
	DeleteLine(string) error
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

// open file resource
func (rh *ResHandler) open(descr int) error {
	f, err := os.OpenFile(rh.Path, descr, rh.Mode)

	if err != nil {
		log.Println(err)
		return err
	}

	rh.Resource = f

	return nil
}

// close the file handler
func (rh ResHandler) close() error {
	err := rh.Resource.Close()

	return err
}

// append line to file
func (rh ResHandler) append(toWrite string) (bool, error) {
	fullStr := fmt.Sprintf("%v\n", toWrite)
	_, err := rh.Resource.WriteString(fullStr)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}

// readContent read, line by line, file content
func (rh *ResHandler) readContent() []string {
	var res []string

	scanner := bufio.NewScanner(rh.Resource)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}

// getOtherLines erase line containing substring
func (rh *ResHandler) getOtherLines(lines []string, content string) []string {
	var res []string

	for _, v := range lines {
		if !strings.Contains(v, content) {
			res = append(res, v)
		}
	}

	return res
}

// writeLines write file line by line
func (rh *ResHandler) writeLines(lines []string) error {
	dw := bufio.NewWriter(rh.Resource)

	for _, v := range lines {
		_, err := dw.WriteString(v + "\n")
		if err != nil {
			log.Println(err)
			return err
		}
	}

	dw.Flush()
	return nil
}

// getLinesToSave returns lines not to be deleted
func (rh ResHandler) getLinesToSave(toDel string) []string {
	rh.open(os.O_RDONLY)
	allLines := rh.readContent()
	linesToSave := rh.getOtherLines(allLines, toDel)
	rh.close()

	return linesToSave
}

// DeleteLine delete line containing string
func (rh ResHandler) DeleteLine(content string) error {
	// // 4 - delete line and print content
	// o.Open(os.O_RDONLY)
	// lines = o.GetOtherLines(lines, "two")
	// o.WriteLines(lines)
	// o.Close()

	// o.Open(os.O_RDONLY)
	// lines = o.ReadContent()
	// o.Close()

	linesToSave := rh.getLinesToSave(content)
	rh.open(os.O_WRONLY | os.O_CREATE | os.O_TRUNC)

	err := rh.writeLines(linesToSave)

	if err != nil {
		log.Println(err)
	}

	rh.close()

	return nil
}

// WriteLine write line to file
func (rh ResHandler) WriteLine(inLine string) error {
	// // 1 - append
	// o.Open(os.O_APPEND | os.O_CREATE | os.O_WRONLY)

	// _, err := o.Append("string one")
	// if err != nil {
	// 	log.Println(err)
	// }

	// _, err = o.Append("string two")
	// if err != nil {
	// 	log.Println(err)
	// }

	// o.Close()
	rh.open(os.O_WRONLY | os.O_CREATE | os.O_APPEND)

	_, err := rh.append(inLine)
	if err != nil {
		log.Println()
		return err
	}

	rh.close()

	return err
}
