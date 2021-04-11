package operator

import "net/http"

type googleClientFactory interface {
	Client(param interface{}) *http.Client
}

type GoogleContext struct {
	GoogleClientFactoryFactory func(credentials []byte, scope string)
}

func NewGoogleConnection() Connection {

}
