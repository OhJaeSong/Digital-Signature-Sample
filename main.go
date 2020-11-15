package main

import (
	"fmt"
)

const version = byte(0x00)
const addressChecksumLen = 4

/* 임의의 문자열 데이터 */
var procData = []byte("abcdefghijklmnopqrstuvwxyz")

func main() {
	/* ECDSA기반key pair (private, public key) 생성 */
	fmt.Printf("%s\n", "--------------------------")
	fmt.Printf("%s\n", "1. ECDSA기반key pair (private, public key) 생성")

	keyPair := NewKey()

	fmt.Printf("PrivateKey: %s\n", keyPair.PrivateKey)
	fmt.Printf("PublicKey: %s\n", keyPair.PublicKey)




	/* 비트코인의 주소 체계와 같은 주소 생성 */
	fmt.Printf("\n")
	fmt.Printf("%s\n", "--------------------------")
	fmt.Printf("%s\n", "2. 비트코인의 주소 체계와 같은 주소 생성")

	address := Address(keyPair.PublicKey)

	fmt.Printf("address: %s\n", address)





	/* 임의의 데이터(문자열)을 private key로 signature 생성 */
	fmt.Printf("\n")
	fmt.Printf("%s\n", "--------------------------")
	fmt.Printf("%s\n", "3. 임의의 데이터(문자열)을 private key로 signature 생성")

	signature, hashData := Signature(keyPair, procData)

	fmt.Printf("signature: %s\n", signature)






	/* Public key를 이용하여 검증 */
	fmt.Printf("\n")
	fmt.Printf("%s\n", "--------------------------")
	fmt.Printf("%s\n", "4. Public key를 이용하여 검증")

	result := Verify(signature, keyPair.PublicKey, hashData)

	fmt.Printf("result: %t\n", result)
}