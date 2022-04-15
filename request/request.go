package request

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/nei7/odrabiamy/config"
	"github.com/nei7/odrabiamy/logger"
)

func Request(method, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logger.ErrorLogger.Fatalf("HTTP: %v \n", err)
		return nil, err
	}

	for k, v := range config.Config.Headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Config.Token))

	res, err := client.Do(req)
	if err != nil {
		logger.ErrorLogger.Fatalf("HTTP: %v \n", err)
		return nil, err
	}

	okStatusCodes := map[int]bool{
		200: true,
		201: true,
	}

	status := fmt.Sprintf("HTTP: %s %s - %s", method, url, res.Status)

	if !okStatusCodes[res.StatusCode] {
		err := fmt.Errorf(status)

		logger.ErrorLogger.Fatalln(err)
		return nil, err
	}

	logger.InfoLogger.Println(status)
	return res, nil
}
