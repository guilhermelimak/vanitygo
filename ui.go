package main

import (
	"fmt"
	"strconv"
	"strings"

	ui "github.com/gizak/termui"
)

func genWalletPar(msg string) *ui.Par {
	p := ui.NewPar(msg)
	p.BorderFg = ui.ColorBlack
	p.PaddingBottom = 1
	p.PaddingLeft = 1
	p.PaddingTop = 1
	p.PaddingRight = 1
	p.Height = 18
	p.Width = 70
	p.Float = ui.AlignCenter
	p.TextFgColor = ui.ColorWhite

	return p
}

func genHashRatePar(message string) *ui.Par {
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

	p := genHashRatePar("")

	ui.Handle("/sys/kbd/q", func(ui.Event) { ui.StopLoop() })
	ui.Handle("/sys/kbd/C-c", func(ui.Event) { ui.StopLoop() })

	ui.Render(p)
	ui.Loop()
}

func drawHashRate(hashRate float64) {
	hr := strconv.FormatFloat(hashRate, 'f', 2, 64)
	p := genHashRatePar(fmt.Sprintf("%s h/s", hr))
	ui.Render(p)
}

func drawVanity(wallet Wallet, vanityString string) *ui.Par {
	formattedVanityString := fmt.Sprintf("[%s](fg-green)", vanityString)
	addr := strings.Replace(wallet.Address, vanityString, formattedVanityString, -1)

	expectedMsg := fmt.Sprintf("[Vanity string:](fg-blue)\n[%s](fg-green)\n", vanityString)
	privMsg := fmt.Sprintf("[Private key:](fg-blue)\n[%s](fg-red)\n", wallet.PrivateKey)
	pubMsg := fmt.Sprintf("[Public key:](fg-blue)\n%s\n", wallet.PublicKey)
	addrMsg := fmt.Sprintf("[Addresss:](fg-blue)\n%s\n", addr)

	msg := "\n" + expectedMsg + "\n" + addrMsg + "\n" + privMsg + "\n" + pubMsg

	p := genWalletPar(msg)

	ui.Render(p)
	return p
}

func drawWallet(wallet Wallet) *ui.Par {
	privMsg := fmt.Sprintf("[Private key:](fg-blue)\n[%s](fg-red)\n", wallet.PrivateKey)
	pubMsg := fmt.Sprintf("[Public key:](fg-blue)\n%s\n", wallet.PublicKey)
	addrMsg := fmt.Sprintf("[Addresss:](fg-blue)\n%s\n", wallet.Address)

	msg := "\n" + addrMsg + "\n" + privMsg + "\n" + pubMsg

	p := genWalletPar(msg)

	ui.Render(p)
	return p
}
