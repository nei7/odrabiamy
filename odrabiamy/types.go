package odrabiamy

type sessionInfo struct {
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

type Exercise struct {
	Content string `json:"content"`
	Page    uint   `json:"page"`
	Number  string `json:"number"`
	Id      int    `json:"id"`
}
