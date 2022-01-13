package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type exercise struct {
	Content string `json:"content"`
	Page    uint   `json:"page"`
	Number  string `json:"number"`
	Id      int    `json:"id"`
}

func LoadExercies(page int, book string) ([]exercise, error) {
	url := fmt.Sprintf("https://odrabiamy.pl/api/v3/books/%s/pages/%d/exercises", book, page)
	res, err := Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var exercises []exercise
	if err := json.NewDecoder(res.Body).Decode(&exercises); err != nil {
		return nil, err
	}

	return exercises, nil
}

func LoadPages(book string) ([]uint, error) {
	url := "https://odrabiamy.pl/api/v3/books/" + book
	res, err := Request(http.MethodGet, url, nil)
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
