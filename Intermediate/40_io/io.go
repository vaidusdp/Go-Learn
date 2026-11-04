package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func readFromReader(r io.Reader){
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		panic(err)
	}
}

func closeResource(c io.Closer){
	err := c.Close()
	if err != nil {
		panic(err)
	}
}

func bufferEx() {
	var buf bytes.Buffer
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderEx() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader("World")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func main() {
	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello Reader!"))

	fmt.Println("=== Write to Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Writer!")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferEx()

	fmt.Println("=== Multi Reader Example ===")
	multiReaderEx()

	fmt.Println("=== Pipe Example ===")
	pipeExample()
}