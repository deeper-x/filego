File handling library - push, append, delete content.
[WIP]

###Rationale 

Simplifying file management, building a file instance, opening and operating on its content.

```go
fh := filego.New("file.txt")

fh.Open()
defer fh.Close()

// Methods should be something like:
fh.Append("last line") // add to the tail
fh.Push("first line") // add to the head
fh.Pop() // remove first line
fm.RemoveLine("some") // remove line containing string "some"

```

