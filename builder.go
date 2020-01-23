package bru_api

type СonnectorBuilderProcess interface {
	setAppID(string) *СonnectorBuilder
	setAddress(string) *СonnectorBuilder
	setSecretKey(string) *СonnectorBuilder
	setAppToken(string) *СonnectorBuilder
}

type СonnectorBuilder struct {
	AppID        string
	AppAddress   string
	AppSecretKey string
	AppToken     string
}

func NewBuilder() *СonnectorBuilder {
	return &СonnectorBuilder{}
}

func (b *СonnectorBuilder) SetAppID(id string) {
	b.AppID = id
}

func (b *СonnectorBuilder) SetAddress(ad string) {
	b.AppAddress = ad
}

func (b *СonnectorBuilder) SetAppSecretKey(k string) {
	b.AppSecretKey = k
}

func (b *СonnectorBuilder) SetAppToken(t string) {
	b.AppToken = t
}
