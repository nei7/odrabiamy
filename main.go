package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V", "v", "version"},
		Usage:   "print version",
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version: %s simple odrabiamy.pl scraper \n",
			color.HiCyanString("Odrabiamy"),
			color.BlueString(c.App.Version))
	}

	app := cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{},
		},
		Name:    "odrabiamy",
		Version: "v0.0.1",
		Usage:   "",
		Commands: []*cli.Command{
			{
				Name:        "info",
				Description: "Get user info",
				Action: func(c *cli.Context) error {
					data, err := GetTokenInfo()
					if err != nil {
						return err
					}

					fmt.Printf("session will expire in %.0f hours \n", time.Unix(int64(data.Exp), 0).Sub(time.Now()).Hours())

					return nil
				},
			},
			{
				Name:        "generate",
				Description: "Generate new jwt token",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
