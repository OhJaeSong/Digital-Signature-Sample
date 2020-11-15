package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"math/big"
)

func Hash(data []byte) []byte {
	var hash [32]byte
	hash = sha256.Sum256(data)

	return hash[:]
}

func Verify(signature, publicKey, data []byte) bool {
	curve := elliptic.P256()

	r := big.Int{}
	s := big.Int{}
	sigLen := len(signature)
	r.SetBytes(signature[:(sigLen / 2)])
	s.SetBytes(signature[(sigLen / 2):])

	x := big.Int{}
	y := big.Int{}
	keyLen := len(publicKey)
	x.SetBytes(publicKey[:(keyLen / 2)])
	y.SetBytes(publicKey[(keyLen / 2):])

	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
	if ecdsa.Verify(&rawPubKey, data, &r, &s) == false {
		return false
	}

	return true
}