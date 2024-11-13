package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func main() {
	// 生成 Ed25519 密钥对
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Public Key: %x\n", publicKey)
	fmt.Printf("Private Key: %x\n", privateKey)
}
