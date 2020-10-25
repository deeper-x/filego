Rationale 

Simplifying file management, simply building a file instance and operating on its content.

```go
fh := filego.New("file.txt")

fh.Open()
defer fh.Close()

fh.Append("last line") // add to the tail
fh.Push("first line") // add to the head
fh.Pop() // remove first line
fm.RemoveLine("some") // remove line containing string "some"

```

