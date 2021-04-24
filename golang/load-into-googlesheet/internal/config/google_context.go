package config

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

// JWTConfigFactory is the type for the function that returns a *jwt.Config.
type JWTConfigFactory func(credentials []byte, scope ...string) (*jwt.Config, error)

// HttpClientFactoryGetter is the type for a function thta returns a function that returns a *http.Client
type HttpClientFactoryGetter func(config *jwt.Config) HttpClientFactory

// HttpClientFactory is the type for a function that returns *http.Client.
type HttpClientFactory func(ctx context.Context) *http.Client

// SheetsServiceFactory is the type for a function that creates the sheets.Service.
type SheetsServiceFactory func(client *http.Client) (*sheets.Service, error)

// GoogleContext is the a context abstraction for Google APIs.
type GoogleContext struct {
	ConfigFactory        JWTConfigFactory
	GetHttpClientFactory HttpClientFactoryGetter
	ServiceFactory       SheetsServiceFactory
	Context              context.Context
}

func NewGoogleContext() *GoogleContext {
	return &GoogleContext{
		ConfigFactory:        google.JWTConfigFromJSON,
		GetHttpClientFactory: getHttpClientFactory,
		ServiceFactory:       sheets.New,
		Context:              context.Background(),
	}
}

func getHttpClientFactory(config *jwt.Config) HttpClientFactory {
	return config.Client
}
