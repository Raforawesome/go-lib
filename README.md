# go-lists
A list library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/Raforawesome/go-lists.svg)](https://pkg.go.dev/github.com/Raforawesome/go-lists)


#### Lists
* StrList
* IntList
* Float64List

The list library is functionally similar to dynamic arrays in high level languages, where they appear to have a dynamic length.  Technically, fixed length arrays have a minor performance benefit over this, however unless you're dealing with massive arrays it won't be noticeable.

Usage:
```go
strList := lists.NewStrList()
strList.append("test")
strList.append("test2")
strList.append("test3")
fmt.Println(strList)
// ["test "test2" "test3"]

strList.SliceAt(1)
fmt.Println(strList)
// ["test" "test3"]

strList.SliceElement("test")
fmt.Println(strList)
// ["test3"]
```
Works the same for all types of Lists available 
(`lists.NewStrList()`, `lists.NewIntList()`, `lists.NewFloat64List()`).
