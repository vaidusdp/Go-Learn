package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	relPath := "./data/file.txt"
	// absPath := "/home/user/docs/file.txt"
	absPath := "C:\\Users\\vaidu\\OneDrive\\Desktop\\file.txt"

	joinedPath := filepath.Join("downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println(normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println(file)
	fmt.Println(dir)
	fmt.Println(filepath.Base("/home/user/docs/"))

	fmt.Println(filepath.IsAbs(absPath))
	fmt.Println(filepath.IsAbs(relPath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	fmt.Println(filepath.Abs(relPath))
}