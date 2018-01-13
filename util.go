package main

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
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
