package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func generateDeviceId() string {
	identifier := make([]byte, 20)
	_, err := rand.Read(identifier)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	key, err := hex.DecodeString("E7309ECC0953C6FA60005B2765F99DBBC965C8E9")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	h := hmac.New(sha1.New, key)
	h.Write(append([]byte{0x19}, identifier...))
	return fmt.Sprintf("19%x%x", identifier, h.Sum(nil))
}

func main() {
	fmt.Println(generateDeviceId())
}