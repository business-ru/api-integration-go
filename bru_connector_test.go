package bru_api

import (
	"log"
	"testing"
)

func TestGetEmployees(t *testing.T) {
	api := NewBuilder()
	api.SetAppID("848593")
	api.SetAddress("https://action_457575.business.ru")
	api.SetAppSecretKey("wDskRiaWuV83wT5H24WDmlFJ3t9UY5ek")
	response := api.Execute(ActionGet, ModelEmployees, nil)

	log.Println(response.AsString)
}
