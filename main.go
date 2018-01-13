package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "vanitygo"
	app.Usage = "Generate vanity bitcoin addresses with golanng"

	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate public and private keys and an address",
			Action: func(c *cli.Context) error {
				go generateSimpleWallet()
				drawUI()
				return nil
			},
		},
		{
			Name:    "vanity",
			Aliases: []string{"v"},
			Usage:   "Generate address until it contains string",
			Action: func(c *cli.Context) error {
				go generateVanity(c.Args().First())
				drawUI()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
