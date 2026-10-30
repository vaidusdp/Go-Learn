package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	// reader := bufio.NewReader(strings.NewReader("Hello Bufio Pacakgeeee!\n"))

	// data := make([]byte, 20)
	// n, err := reader.Read(data)

	// if err != nil {
	// 	fmt.Println("Error Reading:", err)
	// 	return
	// }

	// fmt.Println("Bytes:", n)
	// fmt.Println("Data:", data[:n])
	// fmt.Printf("Data: %d bytes: %s\n",  n, data[:n])

	// s, er := reader.ReadString('\n')
	// if er != nil {
	// 	fmt.Println("Error Reading String:", er)
	// 	return
	// }
	// fmt.Println("Read String:", s)


	writer := bufio.NewWriter(os.Stdout)

	data := []byte("Hello, bufio package!\n")
	n, _ := writer.Write(data)

	fmt.Printf("Wrote %d bytes\n", n)

	er := writer.Flush()
	if er != nil {
		fmt.Println("Error flushing writer:", er)
		return
	}

	str := "fryugiwftb8weceuyrfn9w8y"
	no, _ := writer.WriteString(str)

	fmt.Printf("Wrote %d bytes\n", no)

	err := writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}