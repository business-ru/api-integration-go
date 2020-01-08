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

func TestExecute(t *testing.T) {
	Execute(http.MethodGet, "deals")
}
