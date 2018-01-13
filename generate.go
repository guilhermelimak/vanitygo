package main

import (
	"runtime"
	"strings"
	"sync"
	"time"
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

func _generateVanity(expected string) (hasVanity bool, wallet Wallet) {
	const wallet = _generate()
	return checkAddressVanity(wallet.Address, expected), wallet
}

func generateVanity(expected string) {
	runtime.GOMAXPROCS(4)
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)

	tries := 0

	done := false

	for {
		success, wallet := _generateVanity(expected)

		if success {
			drawWallet(wallet, expected)
			defer wg.Done()
			break
		}

		tries++
		time := time.Since(start).Seconds()
		hps := float64(tries) / time
		drawHashRate(hps)
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
