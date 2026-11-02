package main

import (
	// "fmt"
	// "os"
	// "path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// err := os.Mkdir("subdir", 0755)
	// checkError(err)

	// checkError(os.Mkdir("subdir1", 0755))

	// defer os.RemoveAll("subdir1")

	// os.WriteFile("subdir1/file", []byte(""), 0755)

	// checkError(os.MkdirAll("subdir/parent/child", 0755))
	// checkError(os.MkdirAll("subdir/parent/child1", 0755))
	// checkError(os.MkdirAll("subdir/parent/child2", 0755))

	// os.WriteFile("subdir/parent/file", []byte(""), 0755)
	// os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	// res, err := os.ReadDir("subdir/parent")
	// checkError(err)

	// for _,e := range res {
	// 	fmt.Println(e.Name(), e.Type())
	// }

	// checkError(os.Chdir("subdir/parent/child"))
	// res, err := os.ReadDir(".")
	// checkError(err)

	// for _, e := range res {
	// 	fmt.Println(e.Name(), e.Type())
	// }

	// checkError(os.Chdir("../../.."))
	// dir, err := os.Getwd()
	// checkError(err)
	// fmt.Println(dir)

	// pathfile := "subdir"
	// filepath.WalkDir(pathfile,func(path string, d os.DirEntry, err error) error{
	// 	checkError(err)

	// 	fmt.Println(path)
	// 	return nil
	// })

	// checkError(os.RemoveAll("./subdir"))
	// checkError(os.Remove("subdir1"))
}