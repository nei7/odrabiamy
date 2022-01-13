package main

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/nei7/odrabiamy/config"
)

type payload struct {
	AccessToken        string `json:"accessToken"`
	RefreshToken       string `json:"refreshToken"`
	AccessTokenExpires int    `json:"accessTokenExpires"`
	User               struct {
		UserId uint   `json:"user_id"`
		Email  string `json:"email"`
		Iat    uint   `json:"iat"`
	} `json:"user"`
	Iat int `json:"iat"`
	Exp int `json:"exp"`
}

func GetTokenInfo() (payload, error) {
	encPayload := strings.Split(config.Config.Token, ".")[1]
	tokenPayload, err := base64.StdEncoding.DecodeString(encPayload)
	if err != nil {
		return payload{}, err
	}

	var data payload
	if err := json.Unmarshal(tokenPayload, &data); err != nil {
		return payload{}, nil
	}

	return data, nil
}
