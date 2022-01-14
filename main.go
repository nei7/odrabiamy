package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/nei7/odrabiamy/cmd"
	"github.com/urfave/cli/v2"
)

func main() {

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V", "v"},
		Usage:   "print version",
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version: %s simple odrabiamy.pl scraper \n",
			color.HiCyanString("Odrabiamy"),
			color.BlueString(c.App.Version))
	}

	app := cli.App{
		Name:        "odrabiamy",
		Description: "odrabiamy.pl cli",
		Version:     "v0.0.1",
		Usage:       "",
		Commands: []*cli.Command{
			cmd.InfoCommand(),
			cmd.GetExercises(),
			{
				Name:        "generate",
				Description: "Generate new jwt token",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
