package business_ru_api_integration_go

import (
	"log"
	"net/http"
	"testing"
)

func TestRefreshToken(t *testing.T) {
	RefreshToken()
	log.Println(Token)
}

func TestExecuteSynthetic(t *testing.T) {

	type Tr struct {
		Token       string `json:"token"`
		AppPassword string `json:"app_psw"`
	}

	var s = new(Tr)
	s.AppPassword = "pfasdadsd"
	s.Token = "sdfggaefwfw"

	Execute(http.MethodGet, "deals", nil)
}

func TestGetEmployees(t *testing.T) {

}
