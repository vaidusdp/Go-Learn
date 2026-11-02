package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "pass1234"

	// hash := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))

	// fmt.Println(hash)
	// fmt.Println(hash512)
	// fmt.Printf("SHA-256 Hash Hex Value: %x\n", hash)
	// fmt.Printf("SHA-512 Hash Hex Value: %x\n", hash512)

	salt, _ := genSalt()
	 
	hash := hashPass(password, salt)

	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println(saltStr)
	fmt.Println(hash)

	decodedSalt, _ := base64.StdEncoding.DecodeString(saltStr)
	loginHash := hashPass(password, decodedSalt)

	if hash == loginHash {
		fmt.Println("Pass is correct")
	} else {
		fmt.Println("Pass is incorrect!")
	}

}

func genSalt() ([]byte, error){
	salt := make([]byte, 16)
	_, err :=io.ReadFull(rand.Reader, salt)
	if err != nil{
		return nil, err
	}
	return salt, nil;
}

func hashPass(password string, salt []byte) string {
	saltedPass := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPass)
	return base64.StdEncoding.EncodeToString(hash[:])
}