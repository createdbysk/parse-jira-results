package config

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// OptionWithCredentialsFileFactory is the type for a function that creates an option.ClientOption
// from a credentialsFilePath string
type OptionWithCredentialsFileFactory func(credentialsFilePath string) option.ClientOption

// OptionWithScopesFactory is the type for a function that creates an option.ClientOption
// from scopes
type OptionWithScopesFactory func(scope ...string) option.ClientOption

// SheetsServiceFactory is the type for a function that creates the sheets.Service.
type SheetsServiceFactory func(ctx context.Context, option ...option.ClientOption) (*sheets.Service, error)

// GoogleContext is the a context abstraction for Google APIs.
type GoogleContext struct {
	OptionWithCredentialsFileFactory OptionWithCredentialsFileFactory
	OptionWithScopesFactory          OptionWithScopesFactory
	SheetsServiceFactory             SheetsServiceFactory
	Context                          context.Context
}

func NewGoogleContext() *GoogleContext {
	return &GoogleContext{
		SheetsServiceFactory: sheets.NewService,
		Context:              context.Background(),
	}
}
