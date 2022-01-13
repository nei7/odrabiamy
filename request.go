package main

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

	return res, nil
}
