File handling library - push, append, delete content.
[WIP]

### Rationale 

Simplifying file management, building a file instance, opening and operating on its content.

```go
o := manager.New("./assets/file_demo.txt")

// add lines, one by one
o.WriteLine("demo line 1")
o.WriteLine("demo line 2")
o.WriteLine("demo line 3")

// Then remove one line, that one containing "2"
o.DeleteLine("2")

```

