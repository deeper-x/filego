package main

import "github.com/deeper-x/filego/manager"

func main() {
	// Typical usage:
	o := manager.New("./assets/file_demo.txt")

	// o.WriteLine("demo line 1")
	// o.WriteLine("demo line 2")
	// o.WriteLine("demo line 3")

	o.DeleteLine("3")

}
