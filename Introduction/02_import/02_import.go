package main

// Multiple Import is done by brackets and can be named too
import (
	pr "fmt"
	foo "net/http"
)

func main() {
	pr.Println("Standard Library")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		pr.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	pr.Println("Response Status:", resp.Status)
}