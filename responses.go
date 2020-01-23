package bru_api

type TokenResponse struct {
	Token       string `json:"token"`
	AppPassword string `json:"app_psw"`
}
