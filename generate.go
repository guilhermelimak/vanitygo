package main

import (
	"fmt"
	"strings"
)

// Wallet  is used to group a public and private key and an address
type Wallet struct {
	PublicKey  string
	PrivateKey string
	Address    string
}

func _generate() Wallet {
	privateKey, privateKeyBytes := generatePrivateKey()
	publicKey, publicKeyBytes := generatePublicKey(privateKeyBytes)
	address := generateAddress(publicKeyBytes)

	wallet := Wallet{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Address:    address,
	}

	return wallet
}

func generate() {
	wallet := _generate()

	printKeys(wallet)
}

func _generateVanity(expected string, cSuccess chan bool, cWallet chan Wallet) {
	wallet := _generate()

	hasVanity := checkAddressVanity(wallet.Address, expected)

	cSuccess <- hasVanity
	cWallet <- wallet
}

func generateVanity(expected string) {
	cSuccess := make(chan bool)
	cWallet := make(chan Wallet)

	tries := 0

	for {
		go _generateVanity(expected, cSuccess, cWallet)

		success := <-cSuccess
		wallet := <-cWallet

		if success {
			printKeys(wallet)
			break
		}

		tries++
		fmt.Printf("\x0cTries: %d\n", tries)
	}
}

func generatePrivateKey() (privateKey string, privKeyBytes []byte) {
	bytes, err := getRandomBytes(32)

	if err != nil {
		panic(err)
	}

	version := []byte{128}
	return bytesToWif(version, bytes), bytes
}

func generatePublicKey(privKeyBytes []byte) (string, []byte) {
	publicKeyBytes := privKeyToPub(privKeyBytes)

	return strings.ToLower(bytesToHex(publicKeyBytes)), publicKeyBytes
}

func generateAddress(publicKey []byte) string {
	version := []byte{0}
	ripHash := hashRipemd160(publicKey)

	return bytesToWif(version, ripHash)
}
