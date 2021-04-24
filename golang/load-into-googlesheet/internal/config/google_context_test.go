package config

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"local.dev/sheetsLoader/internal/testutils"
)

func TestGoogleContext(t *testing.T) {
	// GIVEN
	optionWithCredentialsFileFactoryFnPtr := testutils.GetFnPtr(option.WithCredentialsFile)
	optionWithScopesFactoryFnPtr := testutils.GetFnPtr(option.WithScopes)
	sheetsServiceFactoryFnPtr := testutils.GetFnPtr(sheets.NewService)
	ctx := context.Background()

	expected := map[string]interface{}{
		"OptionWithCredentialsFileFactory": optionWithCredentialsFileFactoryFnPtr,
		"OptionWithScopesFactory":          optionWithScopesFactoryFnPtr,
		"SheetsServiceFactory":             sheetsServiceFactoryFnPtr,
		"Context":                          ctx,
	}

	// WHEN
	googleContext := NewGoogleContext()

	actual := map[string]interface{}{
		"OptionWithCredentialsFileFactory": testutils.GetFnPtr(googleContext.OptionWithCredentialsFileFactory),
		"OptionWithScopesFactory":          testutils.GetFnPtr(googleContext.OptionWithScopesFactory),
		"SheetsServiceFactory":             testutils.GetFnPtr(googleContext.SheetsServiceFactory),
		"Context":                          googleContext.Context,
	}

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestGoogleContext: expected: %v, actual %v",
			expected,
			actual,
		)
	}
}
