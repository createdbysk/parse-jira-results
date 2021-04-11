package operator

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2/jwt"
)

// JWTConfigFactory is the type for the function that returns a *jwt.Config.
type JWTConfigFactory func(credentials []byte, scope string) *jwt.Config

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

type connection struct {
	client *http.Client
}

func NewGoogleConnection(googleCtx *GoogleContext, credentials []byte, scope string) (Connection, error) {
	config := googleCtx.ConfigFactory(credentials, scope)
	httpClientFactory := googleCtx.GetHttpClientFactory(config)
	httpClient := httpClientFactory(googleCtx.Context)
	c := &connection{httpClient}
	return c, nil
}

func (c *connection) Get(impl interface{}) {
	err := errors.New("not implemented")
	panic(err)
}
