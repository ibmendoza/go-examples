package main

import (
	"fmt"
	"github.com/ereyes01/cryptohelper"
)

func main() {
	messageToEncryptThenMAC := "The big brown fox jumps over the lazy dog"

	secret, _ := cryptohelper.RandomKey()

	msgEncrypted, _ := cryptohelper.SecretboxEncrypt(messageToEncryptThenMAC, secret)

	fmt.Println("Encrypted string")
	fmt.Println(msgEncrypted)

	msgDecrypted, _ := cryptohelper.SecretboxDecrypt(msgEncrypted, secret)

	fmt.Println("Decrypted string")
	fmt.Println(msgDecrypted)
}
