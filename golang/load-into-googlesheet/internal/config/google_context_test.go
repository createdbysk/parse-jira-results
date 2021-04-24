package config

import (
	"context"
	"reflect"
	"testing"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
	"local.dev/sheetsLoader/internal/testutils"
)

func TestGoogleContext(t *testing.T) {
	// GIVEN
	jwtConfigFactoryPtr := testutils.GetFnPtr(google.JWTConfigFromJSON)
	getHttpClientFactoryFnPtr := testutils.GetFnPtr(getHttpClientFactory)
	sheetsServiceFactoryFnPtr := testutils.GetFnPtr(sheets.New)
	ctx := context.Background()

	expected := map[string]interface{}{
		"ConfigFactory":        jwtConfigFactoryPtr,
		"GetHttpClientFactory": getHttpClientFactoryFnPtr,
		"ServiceFactory":       sheetsServiceFactoryFnPtr,
		"ctx":                  ctx,
	}

	// WHEN
	googleContext := NewGoogleContext()

	actual := map[string]interface{}{
		"ConfigFactory":        testutils.GetFnPtr(googleContext.ConfigFactory),
		"GetHttpClientFactory": testutils.GetFnPtr(googleContext.GetHttpClientFactory),
		"ServiceFactory":       testutils.GetFnPtr(googleContext.ServiceFactory),
		"ctx":                  googleContext.Context,
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

func TestGetHttpClientFactory(t *testing.T) {
	// GIVEN
	jwtConfig := &jwt.Config{}
	googleContext := NewGoogleContext()

	expected := testutils.GetFnPtr(jwtConfig.Client)

	// WHEN
	actual := testutils.GetFnPtr(googleContext.GetHttpClientFactory(jwtConfig))

	// THEN
	if actual != expected {
		t.Errorf(
			"getHttpClientFactory: expected %v, actual %v",
			expected,
			actual,
		)
	}
}
