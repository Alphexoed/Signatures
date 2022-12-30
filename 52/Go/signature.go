package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func generateSignature(data string) string {
	key, err := hex.DecodeString("EAB4F1B9E3340CD1631EDE3B587CC3EBEDF1AFA9")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(data))
	signature := append([]byte{'R'}, mac.Sum(nil)...)
	return base64.StdEncoding.EncodeToString(signature)
}

func main() {
	fmt.Println(generateSignature("teste"))
}