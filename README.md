File handling library - push, append, delete content.
[WIP]

### Rationale 
Simplifying file management, building a file instance, opening and operating on its content.

#### Example
Writing a demo log file appending IP address; then deleting all entries containing given IP

```go
package main

import "github.com/deeper-x/filego"

func main() {
    fg := filego.New("./logs/access.log")

    // add lines, one by one
    fg.WriteLine("Access from 100.100.100.100")
    fg.WriteLine("Access from 101.101.101.101")
    fg.WriteLine("Access from 102.102.102.102")
    fg.WriteLine("Access from 101.101.101.101")
    fg.WriteLine("Access from 102.102.102.102")
    fg.WriteLine("Access from 101.101.101.101")
	
    // Then remove all the lines containing given ip address
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

#### Unit Test 
```sh
~/go/src/github.com/deeper-x/filego  (master)$ make test
go test -v -cover ./...
=== RUN   TestWriteLine
--- PASS: TestWriteLine (0.00s)
=== RUN   TestDeleteLine
--- PASS: TestDeleteLine (0.00s)
PASS
coverage: 9.8% of statements
ok  	github.com/deeper-x/filego	(cached)	coverage: 9.8% of statements

```
