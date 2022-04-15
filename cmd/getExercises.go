package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/nei7/odrabiamy/config"
	"github.com/nei7/odrabiamy/odrabiamy"
	"github.com/nei7/odrabiamy/s3"
	"github.com/urfave/cli/v2"
)

func GetExercises() *cli.Command {
	return &cli.Command{
		Name: "get",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "book",
				Usage:    "Book id",
				Aliases:  []string{"b"},
				Required: true,
			},
			&cli.UintFlag{
				Name:     "start",
				Usage:    "From which page scraper will start",
				Aliases:  []string{"s"},
				Required: true,
			},
			&cli.UintFlag{
				Name:     "count",
				Usage:    "How many page scraper will download",
				Aliases:  []string{"c"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			start := c.Uint("start")
			count := c.Uint("count")
			book := c.String("book")

			tokenInfo, err := odrabiamy.GetTokenInfo()
			if err != nil {
				return err
			}

			client := odrabiamy.NewClient()

			if time.Now().After(time.Unix(int64(tokenInfo.AccessTokenExpires), 0)) {
				if err := client.GenerateSession(); err != nil {
					return err
				}
			}
			pages, err := client.LoadPages(book)
			if err != nil {
				return err
			}
			startPageIndex := getStartIndex(pages, start)
			if startPageIndex == -1 {
				return errors.New("invalid page")
			}
			if startPageIndex+int(count) > len(pages) {
				return errors.New("you can't get that many pages")
			}

			session := s3.InitSession()

			if _, err := os.Stat(config.Config.Path); err != nil {
				os.MkdirAll(config.Config.Path, 0777)
			}

			bookPath := path.Join(config.Config.Path, book)
			if err := os.Mkdir(bookPath, 0777); err != nil && !errors.Is(err, os.ErrExist) {
				return err
			}

			for i := startPageIndex; i < int(count); i++ {
				page := pages[i]
				exercises, err := client.LoadExercies(page, book)
				if err != nil {
					return err
				}

				if err := os.MkdirAll(path.Join(bookPath, fmt.Sprint(page)), 0777); err != nil && !errors.Is(err, os.ErrExist) {
					return err
				}

				for _, exercise := range exercises {

					path := path.Join(config.Config.Path, book, fmt.Sprint(page), exercise.Number+".html")
					f, err := os.Create(path)
					if err != nil {
						return err
					}
					defer f.Close()

					if _, err := f.WriteString(exercise.Content); err != nil {
						return err
					}

					if err = s3.UploadFile(session, path); err != nil {
						return err
					}
				}
			}
			return nil
		},
	}
}

func getStartIndex(pages []uint, start uint) int {
	for index, page := range pages {
		if page == start {
			return index
		}
	}
	return -1

}
