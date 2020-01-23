package business_ru_api_integration_go

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Полеучение нового токена
func RefreshToken(b *connectorBuilder) {

	u := GetURL(b, "repair")
	uq := u.Query()
	uq.Set("app_id", b.AppID)
	uq.Set("app_psw", GetMD5Hash(b.AppSecretKey+uq.Encode()))

	u.RawQuery = uq.Encode()

	resp, err := http.Get(u.String())

	if err != nil {
		log.Fatalln(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var s = new(TokenResponse)

	err = json.Unmarshal(body, &s)

	if err != nil {
		log.Fatalln(err.Error())
	}

	b.AppToken = s.Token
}

func (b *connectorBuilder) Execute(Action string, Model string, Params interface{}) {

	if b.AppToken == "" {
		RefreshToken(b)
	}

	u := GetURL(b, Model)

	uq := u.Query()
	uq.Set("app_id", b.AppID)
	uq.Set("app_psw", GetMD5Hash(b.AppToken+b.AppSecretKey+uq.Encode()))
	u.RawQuery = uq.Encode()

	client := &http.Client{}

	req, err := http.NewRequest(Action, u.String(), nil)

	if err != nil {
		log.Fatalln(err.Error())
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var body = ParseResponseBody(resp.Body)

	log.Println(GetResponseBody(body))

	TokenRenew(body, b)

}

// Получение MD5-хеша строки
func GetMD5Hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func GetURL(b *connectorBuilder, m string) *url.URL {
	var ub strings.Builder

	ub.WriteString(b.AppAddress)
	ub.WriteString("/api/rest/")
	ub.WriteString(m)
	ub.WriteString(".json")

	u, err := url.Parse(ub.String())

	if err != nil {
		log.Fatalln(err.Error())
	}

	return u
}

func TokenRenew(Body []byte, b *connectorBuilder) {

	var s = new(TokenResponse)

	err := json.Unmarshal(Body, &s)

	if err != nil {
		log.Fatalln("UNMARSHAL RESPONSE BODY " + err.Error())
	}

	b.AppToken = s.Token
}

func GetResponseBody(Body []byte) string {

	bodyString := string(Body)
	return bodyString
}

func ParseResponseBody(Body io.ReadCloser) []byte {
	bodyBytes, err := ioutil.ReadAll(Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	return bodyBytes
}
