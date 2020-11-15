package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type KeyPair struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewKey() *KeyPair {
	private, public := newKeyPair()
	keyPair := KeyPair{private, public}

	return &keyPair
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, _ := ecdsa.GenerateKey(curve, rand.Reader)
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}