package business_ru_api_integration_go

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	ApiPath = "/api/rest/"
)

type ApiBuilder interface {
	setAddress(Address string) ApiBuilder
	setAppID(AppID string) ApiBuilder
	setSecretKey(SecretKey string) ApiBuilder
	setModel(Model string) ApiBuilder
	setAction() ApiBuilder
	setParams() ApiBuilder
}

type CommandBuilder struct {
	Address   string
	AppID     string
	SecretKey string
	Model     string
	Action    string
	Token     string
	Params    struct{}
}

func New() {
	return &ApiBuilder()
}

func (b *CommandBuilder) setAddress(Address string) {
	b.Address = Address
}

func (b *CommandBuilder) setAppID(AppID string) {
	b.AppID = AppID
}

func (b *CommandBuilder) setSecretKey(SecretKey string) {
	b.SecretKey = SecretKey
}

func (b *CommandBuilder) setToken(Token string) {
	b.Token = Token
}

func (b *CommandBuilder) getToken() string {
	return b.Token
}

// Обновление токена
func RefreshToken() {
	CommandBuilder{}.setToken(GetRefreshToken())
}

// Полеучение нового токена
func GetRefreshToken() string {

	u := GetURL("repair")
	uq := u.Query()
	uq.Set("app_id", CommandBuilder{}.SecretKey)
	uq.Set("app_psw", GetMD5Hash(CommandBuilder{}.SecretKey+uq.Encode()))

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

func Execute(Action string, Model string, Params interface{}) string {

	if Params == nil {
		fmt.Println("Params is nil")
	}

	if &CommandBuilder.getToken() == "" {
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

	var body = ParseResponseBody(resp.Body)

	log.Println(GetResponseBody(body))

	ExecutionResultString = GetResponseBody(body)

	TokenRenew(body)

	return ExecutionResultString
}

func getResultAsString() string {
	return ExecutionResultString
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

func TokenRenew(Body []byte) {

	var s = new(TokenResponse)

	err := json.Unmarshal(Body, &s)

	if err != nil {
		log.Fatalln("UNMARSHAL RESPONSE BODY " + err.Error())
	}

	Token = s.Token
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
