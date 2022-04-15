package odrabiamy

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/nei7/odrabiamy/config"
	"github.com/nei7/odrabiamy/request"
)

type client struct{}

func NewClient() *client {
	return &client{}
}

func (c *client) LoadExercies(page uint, book string) ([]Exercise, error) {
	url := fmt.Sprintf("https://odrabiamy.pl/api/v3/books/%s/pages/%d/exercises", book, page)
	res, err := request.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var exercises []Exercise
	if err := json.NewDecoder(res.Body).Decode(&exercises); err != nil {
		return nil, err
	}

	return exercises, nil
}

func (c *client) GenerateSession() error {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, "https://odrabiamy.pl/api/auth/session", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Cookie", " __Secure-next-auth.session-token="+config.Config.Token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var data struct {
		AccessToken string `json:"accessToken"`
	}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	config.Config.Token = data.AccessToken
	return nil
}

func (c *client) LoadPages(book string) ([]uint, error) {
	url := "https://odrabiamy.pl/api/v3/books/" + book
	res, err := request.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pages struct {
		Pages []uint `json:"pages"`
	}

	if json.NewDecoder(res.Body).Decode(&pages); err != nil {
		return nil, err
	}

	return pages.Pages, nil
}

func GetTokenInfo() (sessionInfo, error) {
	encPayload := strings.Split(config.Config.Token, ".")[1]
	tokenPayload, err := base64.StdEncoding.DecodeString(encPayload)
	if err != nil {
		return sessionInfo{}, err
	}

	var data sessionInfo
	if err := json.Unmarshal(tokenPayload, &data); err != nil {
		return sessionInfo{}, nil
	}

	return data, nil
}
