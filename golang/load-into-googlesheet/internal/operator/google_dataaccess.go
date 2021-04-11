package operator

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// JWTConfigFactory is the type for the function that returns a *jwt.Config.
type JWTConfigFactory func(credentials []byte, scope ...string) (*jwt.Config, error)

// HttpClientFactoryGetter is the type for a function thta returns a function that returns a *http.Client
type HttpClientFactoryGetter func(config *jwt.Config) HttpClientFactory

// HttpClientFactory is the type for a function that returns *http.Client.
type HttpClientFactory func(ctx context.Context) *http.Client

// GoogleContext is the a context abstraction for Google APIs.
type GoogleContext struct {
	ConfigFactory        JWTConfigFactory
	GetHttpClientFactory HttpClientFactoryGetter
	Context              context.Context
}

func NewGoogleContext() *GoogleContext {
	return &GoogleContext{
		ConfigFactory:        google.JWTConfigFromJSON,
		GetHttpClientFactory: GetHttpClientFactory,
		Context:              context.Background(),
	}
}

type connection struct {
	client *http.Client
}

func NewGoogleConnection(googleCtx *GoogleContext, credentials []byte, scope ...string) (Connection, error) {
	config, err := googleCtx.ConfigFactory(credentials, scope...)
	if err != nil {
		return nil, err
	}
	httpClientFactory := googleCtx.GetHttpClientFactory(config)
	httpClient := httpClientFactory(googleCtx.Context)
	c := &connection{httpClient}
	return c, nil
}

func (c *connection) Get(impl interface{}) {
	err := errors.New("not implemented")
	panic(err)
}

func GetHttpClientFactory(config *jwt.Config) HttpClientFactory {
	err := errors.New("not implemented")
	panic(err)
	return nil
}
