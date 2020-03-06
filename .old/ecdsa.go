package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

func main() {
	signData := "wangkaixuan"
	signature, publicKey := generateSignature(signData)

	r1 := big.Int{}
	s1 := big.Int{}

	r1.SetBytes(signature[:len(signature)/2])
	s1.SetBytes(signature[len(signature)/2:])

	hash := sha256.Sum256([]byte(signData))
	res := ecdsa.Verify(&publicKey, hash[:], &r1, &s1)

	fmt.Printf("signature result:%v\n", res)
}

func generateSignature(signData string) ([]byte, ecdsa.PublicKey) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	publicKey := privateKey.PublicKey
	hash := sha256.Sum256([]byte(signData))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Panic(err)
	}

	signature := append(r.Bytes(), s.Bytes()...)
	return signature, publicKey
}
