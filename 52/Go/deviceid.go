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

	key, err := hex.DecodeString("AE49550458D8E7C51D566916B04888BFB8B3CA7D")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	h := hmac.New(sha1.New, key)
	h.Write(append([]byte{0x52}, identifier...))
	return fmt.Sprintf("52%x%x", identifier, h.Sum(nil))
}

func main() {
	fmt.Println(generateDeviceId())
}