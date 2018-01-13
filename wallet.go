package main

import "strings"

// Wallet is used to group a public and private key and an address
type Wallet struct {
	PublicKey       string
	PrivateKey      string
	Address         string
	privateKeyBytes []byte
	publicKeyBytes  []byte
}

// GenerateWallet is used to create a new wallet
func GenerateWallet() Wallet {
	wallet := Wallet{}
	w := &wallet

	w.generatePrivateKey()
	w.generatePublicKey()
	w.generateAddress()

	return wallet
}

func (w *Wallet) generatePrivateKey() {
	bytes, err := getRandomBytes(32)

	if err != nil {
		panic(err)
	}

	version := []byte{128}
	w.PrivateKey = bytesToWif(version, bytes)
	w.privateKeyBytes = bytes
}

func (w *Wallet) generatePublicKey() {
	bytes := privKeyToPub(w.privateKeyBytes)
	w.PublicKey = strings.ToLower(bytesToHex(bytes))
	w.publicKeyBytes = bytes
}

func (w *Wallet) generateAddress() {
	version := []byte{0}
	ripHash := hashRipemd160(w.publicKeyBytes)
	w.Address = bytesToWif(version, ripHash)
}
