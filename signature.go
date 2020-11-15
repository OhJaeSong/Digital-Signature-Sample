package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"log"
)

func Signature(keyPair *KeyPair, procData []byte) ([]byte, []byte) {
	hashData := Hash(procData)

	r, s, err := ecdsa.Sign(rand.Reader, &keyPair.PrivateKey, hashData)

	if err != nil {
		log.Panic(err)
	}

	signature := append(r.Bytes(), s.Bytes()...)

	return signature, hashData
}