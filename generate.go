package main

import (
	"runtime"
	"strings"
	"sync"
	"time"
)

func generateSimpleWallet() {
	drawWallet(GenerateWallet())
}

func generateVanity(expected string) {
	runtime.GOMAXPROCS(4)
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)

	tries := 0

	for {
		wallet := GenerateWallet()
		success := strings.Contains(wallet.Address, expected)

		if success {
			drawVanity(wallet, expected)
			defer wg.Done()
			break
		}

		tries++
		time := time.Since(start).Seconds()
		hps := float64(tries) / time
		drawHashRate(hps)
	}
}
