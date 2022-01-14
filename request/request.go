package request

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/nei7/odrabiamy/config"
)

func Request(method, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range config.Config.Headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Config.Token))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	okStatusCodes := map[int]bool{
		200: true,
		201: true,
	}

	if !okStatusCodes[res.StatusCode] {
		return nil, fmt.Errorf("%s %s - %s", method, url, res.Status)
	}

	return res, nil
}