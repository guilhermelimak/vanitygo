package main

import (
	"fmt"
	"strconv"
	"strings"

	ui "github.com/gizak/termui"
)

func genPar(message string) *ui.Par {
	p := ui.NewPar(message)
	p.BorderFg = ui.ColorBlack
	p.Height = 3
	p.PaddingLeft = 13
	p.Width = 50
	p.Float = ui.AlignCenter

	p.PaddingLeft = (p.Width-len(message))/2 - 1 // center text inside box

	p.TextFgColor = ui.ColorWhite

	return p
}

func drawUI() {
	err := ui.Init()

	if err != nil {
		panic(err)
	}

	defer ui.Close()
	p := genPar("75.32132 h/s")

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Render(p)
	ui.Loop()
}

func drawHashRate(hashRate float64) {
	hr := strconv.FormatFloat(hashRate, 'f', 2, 64)
	p := genPar(fmt.Sprintf("%s h/s", hr))
	ui.Render(p)
}

func drawWallet(wallet Wallet, expected string) *ui.Par {
	formattedExpected := fmt.Sprintf("[%s](fg-green)", expected)
	addr := strings.Replace(wallet.Address, expected, formattedExpected, -1)

	expectedMsg := fmt.Sprintf("Expected string: [%s](fg-green)\n", expected)
	privMsg := fmt.Sprintf("Private key:\n\n[%s](fg-red)\n", wallet.PrivateKey)
	pubMsg := fmt.Sprintf("Public key:\n\n%s\n", wallet.PublicKey)
	addrMsg := fmt.Sprintf("Addresss:\n\n%s\n", addr)

	msg := privMsg + "\n" + pubMsg + "\n" + expectedMsg + "\n" + addrMsg

	p := ui.NewPar(msg)
	p.BorderFg = ui.ColorBlack
	p.PaddingBottom = 1
	p.PaddingLeft = 1
	p.PaddingTop = 1
	p.PaddingRight = 1
	p.Height = 19
	p.Width = 70
	p.Float = ui.AlignCenter
	p.TextFgColor = ui.ColorWhite

	ui.Render(p)
	return p
}
