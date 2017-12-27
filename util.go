package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/akamensky/base58"
	secp256k1 "github.com/toxeus/go-secp256k1"
	"golang.org/x/crypto/ripemd160"
)

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

func bytesToWif(version []byte, bytes []byte) string {
	keyWithVersion := append(version, bytes...)

	doubleSha := doubleSha256(keyWithVersion)
	checksum := doubleSha[0:4]

	keyWithChecksum := append(keyWithVersion, checksum...)

	return base58.Encode(keyWithChecksum)
}

func privKeyToPub(privKeySlice []byte) []byte {
	var privKey [32]byte

	copy(privKey[:], privKeySlice[:])

	secp256k1.Start()
	defer secp256k1.Stop()

	isPrivKeyValid := secp256k1.Seckey_verify(privKey)

	if !isPrivKeyValid {
		panic("Invalid private key")
	}

	pubKey, success := secp256k1.Pubkey_create(privKey, false)

	if !success {
		panic("Error creating public key")
	}

	isPubKeyvalid := secp256k1.Pubkey_verify(pubKey)

	if !isPubKeyvalid {
		panic("Invalid public key")
	}

	return pubKey
}

func hashRipemd160(publicKey []byte) []byte {
	r := ripemd160.New()

	pubKeySha := sha256.Sum256(publicKey)

	r.Write(pubKeySha[:])

	return r.Sum(nil)
}

func doubleSha256(bytes []byte) []byte {
	firstSha := sha256.Sum256(bytes)
	doubleSha := sha256.Sum256(firstSha[:])
	return doubleSha[:]
}

func checkAddressVanity(addr string, expected string) bool {
	return strings.Contains(addr, expected)
}

func printKeys(pubKey string, privateKey string, address string) {
	println("================================================================================")
	println()
	fmt.Printf("     Private key:\n       %s\n", privateKey)
	println()
	fmt.Printf("     Public key:\n       %s\n", pubKey)
	println()
	fmt.Printf("     Addresss:\n       %s\n", address)
	println()
	println("================================================================================")
}
