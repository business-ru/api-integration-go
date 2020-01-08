package business_ru_api_integration_go

type TokenResponse struct {
	Token       string `json:"token"`
	AppPassword string `json:"app_psw"`
}

type ExecutionResponse struct {
	Status string `json:"status"`
	Result []Deal `json:"result"`
}

type Deal struct {
	Id         string
	TimeCreate string
}
