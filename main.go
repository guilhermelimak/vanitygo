package main

import (
	"strings"
)

func main() {
	privateKey, privateKeyBytes := getPrivateKey()
	publicKey, publicKeyBytes := getPubKey(privateKeyBytes)
	address := generateAddress(publicKeyBytes)
	printKeys(publicKey, privateKey, address)
}

func getPrivateKey() (privateKey string, privKeyBytes []byte) {
	bytes, err := getRandomBytes(32)

	if err != nil {
		panic(err)
	}

	version := []byte{128}
	return bytesToWif(version, bytes), bytes
}

func getPubKey(privKeyBytes []byte) (string, []byte) {
	publicKeyBytes := privKeyToPub(privKeyBytes)

	return strings.ToLower(bytesToHex(publicKeyBytes)), publicKeyBytes
}

func generateAddress(publicKey []byte) string {
	// Version concatenated with RIPEMD-160(SHA-256(public key))
	version := []byte{0}
	ripHash := hashRipemd160(publicKey)

	return bytesToWif(version, ripHash)
}
