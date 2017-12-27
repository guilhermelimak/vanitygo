package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/akamensky/base58"
	"github.com/toxeus/go-secp256k1"
)

func main() {
	_, _, privBytes := getPrivateKey()

	pubKeyStr, _ := seccpBytes(privBytes)

	println(pubKeyStr)
}

func getPrivateKey() (wifKey string, hexKey string, bytes []byte) {
	bytes, err := getRandomBytes(32)

	if err != nil {
		panic(err)
	}

	return bytesToWif(bytes), bytesToHex(bytes), bytes
}

func getRandomBytes(amount int) ([]byte, error) {
	bytes := make([]byte, amount)

	_, err := rand.Read(bytes)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func bytesToHex(bytes []byte) string {
	return strings.ToUpper(hex.EncodeToString(bytes))
}

func bytesToWif(bytes []byte) string {
	keyWithVersion := append([]byte{128}, bytes...)

	firstSHA := sha256.Sum256(keyWithVersion)
	secondSHA := sha256.Sum256(firstSHA[:])

	checksum := secondSHA[0:4]
	keyWithChecksum := append(keyWithVersion, checksum...)

	return base58.Encode(keyWithChecksum)
}

func seccpBytes(privKeySlice []byte) (string, []byte) {
	var privKey [32]byte

	copy(privKey[:], privKeySlice[0:32])

	secp256k1.Start()
	defer secp256k1.Stop()

	isPrivKeyValid := secp256k1.Seckey_verify(privKey)

	if !isPrivKeyValid {
		panic("Invalid private key")
	}

	pubKey, success := secp256k1.Pubkey_create(privKey, true)

	if !success {
		panic("Error creating public key")
	}

	isPubKeyvalid := secp256k1.Pubkey_verify(pubKey)

	if !isPubKeyvalid {
		panic("Invalid public key")
	}

	return strings.ToLower(bytesToHex(pubKey)), pubKey
}
