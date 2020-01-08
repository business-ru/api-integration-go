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

var (
	SecretKey = "wDskRiaWuV83wT5H24WDmlFJ3t9UY5ek"
	AppID     = "848593"
	Address   = "https://action_457575.business.ru"
	Token     = ""

	ApiPath = "/api/rest/"
)

// Обновление токена
func RefreshToken() {
	Token = GetRefreshToken()
}

// Полеучение нового токена
func GetRefreshToken() string {

	u := GetURL("repair")
	uq := u.Query()
	uq.Set("app_id", AppID)
	uq.Set("app_psw", GetMD5Hash(SecretKey+uq.Encode()))

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

	return s.Token
}

func Execute(Action string, Model string) {
	if Token == "" {
		RefreshToken()
	}

	u := GetURL(Model)

	uq := u.Query()
	uq.Set("app_id", AppID)
	uq.Set("app_psw", GetMD5Hash(Token+SecretKey+uq.Encode()))
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

	TokenRenew(resp.Body)
}

// Получение MD5-хеша строки
func GetMD5Hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func GetURL(m string) *url.URL {
	var ub strings.Builder

	ub.WriteString(Address)
	ub.WriteString(ApiPath)
	ub.WriteString(m)
	ub.WriteString(".json")

	u, err := url.Parse(ub.String())

	if err != nil {
		log.Fatalln(err.Error())
	}

	return u
}

func TokenRenew(Body io.ReadCloser) {
	body, err := ioutil.ReadAll(Body)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var s = new(TokenResponse)

	err = json.Unmarshal(body, &s)

	if err != nil {
		log.Fatalln(err.Error())
	}

	Token = s.Token
}
