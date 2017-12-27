package main

import (
	"crypto/sha256"
	"strings"

	"golang.org/x/crypto/ripemd160"
)

func main() {
	_, _, privBytes := getPrivateKey()
	_, pubKeyBytes := getPubKey(privBytes)
	println(generateAddress(pubKeyBytes))
}

func getPrivateKey() (wifKey string, hexKey string, bytes []byte) {
	bytes, err := getRandomBytes(32)

	if err != nil {
		panic(err)
	}
	version := []byte{128}
	return bytesToWif(version, bytes), bytesToHex(bytes), bytes
}

func getPubKey(privKeyBytes []byte) (string, []byte) {
	pubKeyBytes := privKeyToPub(privKeyBytes)

	return strings.ToLower(bytesToHex(pubKeyBytes)), pubKeyBytes
}

func generateAddress(publicKey []byte) string {
	// Version concatenated with RIPEMD-160(SHA-256(public key))
	version := []byte{0}
	ripemdHash := getRipemd160(publicKey)

	return bytesToWif(version, ripemdHash)
}

func getRipemd160(publicKey []byte) []byte {
	r := ripemd160.New()

	pubKeySha := sha256.Sum256(publicKey)

	r.Write(pubKeySha[:])

	return r.Sum(nil)
}
