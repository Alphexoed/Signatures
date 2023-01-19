package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func generateSignature(data string) string {
	key, err := hex.DecodeString("DFA5ED192DDA6E88A12FE12130DC6206B1251E44")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(data))
	signature := append([]byte{0x19}, mac.Sum(nil)...)
	return base64.StdEncoding.EncodeToString(signature)
}

func main() {
	fmt.Println(generateSignature("teste"))
}