package business_ru_api_integration_go

type connectorBuilderProcess interface {
	setAppID(string) *connectorBuilder
	setAddress(string) *connectorBuilder
	setSecretKey(string) *connectorBuilder
	setAppToken(string) *connectorBuilder
}

type connectorBuilder struct {
	AppID        string
	AppAddress   string
	AppSecretKey string
	AppToken     string
}

func newBuilder() *connectorBuilder {
	return &connectorBuilder{}
}

func (b *connectorBuilder) setAppID(id string) {
	b.AppID = id
}

func (b *connectorBuilder) setAddress(ad string) {
	b.AppAddress = ad
}

func (b *connectorBuilder) setAppSecretKey(k string) {
	b.AppSecretKey = k
}

func (b *connectorBuilder) setAppToken(t string) {
	b.AppToken = t
}
