package business_ru_api_integration_go

import (
	"log"
	"testing"
)

func TestGetEmployees(t *testing.T) {
	api := newBuilder()
	api.setAppID("")
	api.setAddress("")
	api.setAppSecretKey("")
	api.Execute(actionGet, "", nil)

	log.Println(api.AppToken)
}
