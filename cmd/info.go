package cmd

import (
	"fmt"
	"time"

	"github.com/nei7/odrabiamy/odrabiamy"
	"github.com/urfave/cli/v2"
)

func InfoCommand() *cli.Command {
	return &cli.Command{
		Name:        "info",
		Description: "Get user info",
		Action: func(c *cli.Context) error {
			data, err := odrabiamy.GetTokenInfo()
			if err != nil {
				return err
			}

			fmt.Printf("cookie will expire in %.0f hours \n", time.Unix(int64(data.Exp), 0).Sub(time.Now()).Hours())

			fmt.Printf("access token will expire in %.0f minutes \n", time.Unix(int64(data.AccessTokenExpires), 0).Sub(time.Now()).Minutes())

			return nil
		},
	}
}
