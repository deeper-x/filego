File handling library - push, append, delete content.
[WIP]

### Rationale 
Simplifying file management, building a file instance, opening and operating on its content.

#### Example
Writing a demo log file, appending IP address, then deleting all entries containing given IP

```go
package main

import "github.com/deeper-x/filego"

func main() {
    fg := filego.New("./logs/access.log")

    // add lines, one by one
    fg.WriteLine("100.100.100.100")
    fg.WriteLine("101.101.101.101")
    fg.WriteLine("102.102.102.102")
    fg.WriteLine("101.101.101.101")
    fg.WriteLine("102.102.102.102")
    fg.WriteLine("101.101.101.101")
	
    // Then remove one line, that one containing given ip address
    fg.DeleteLine("101.101.101.101")
}
```
Results:

```sh
$ cat ./logs/access.log 
100.100.100.100
102.102.102.102
102.102.102.102

```

#### Test [WIP]
```sh
$ make test
```
