package bru_api

import (
	"log"
	"testing"
)

func TestGetEmployees(t *testing.T) {
	api := NewBuilder()
	api.SetAppID("")
	api.SetAddress("")
	api.SetAppSecretKey("")
	api.Execute(ActionGet, "", nil)

	log.Println(api.AppToken)
}
