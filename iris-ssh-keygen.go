//http://stackoverflow.com/questions/21151714/go-generate-an-ssh-public-key
package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln("error generating private key")
	}

	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}
	privateKeyPem := string(pem.EncodeToMemory(&privateKeyBlock))

	publicKey := privateKey.PublicKey
	publicKeyDer, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Fatalln("error generating public key")
	}

	publicKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyDer,
	}
	publicKeyPem := string(pem.EncodeToMemory(&publicKeyBlock))

	f, _ := os.Create("id_rsa")
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(privateKeyPem)
	w.Flush()

	//fmt.Println(privateKeyPem)
	//fmt.Println(publicKeyPem)

	g, _ := os.Create("id_rsa.pub")
	defer g.Close()
	w = bufio.NewWriter(g)
	w.WriteString(publicKeyPem)
	w.Flush()
}
