File handling library - push, append, delete content.
[WIP]

### Rationale 

Simplifying file management, building a file instance, opening and operating on its content.

```go
package main

import "github.com/deeper-x/filego"

func main() {
	fg := filego.New("./assets/file_demo.txt")

	// add lines, one by one
	fg.WriteLine("demo line 1")
	fg.WriteLine("demo line 2")
	fg.WriteLine("demo line 3")

	// Then remove one line, that one containing "2"
	fg.DeleteLine("2")
}
```
Results:

```sh
$ cat assets/file_demo.txt 
demo line 1
demo line 3

```
