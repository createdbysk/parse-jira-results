package operator

import (
	"bytes"
	"context"
	"net/http"
	"reflect"
	"testing"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"local.dev/sheetsLoader/internal/testutils"
)

type mockHttpClientFactoryGetter struct {
	called         bool
	ctx            context.Context
	mockHttpClient *http.Client
}

func (c *mockHttpClientFactoryGetter) get(config *jwt.Config) HttpClientFactory {
	c.called = true
	return c.httpClientFactory
}

func (c *mockHttpClientFactoryGetter) httpClientFactory(ctx context.Context) *http.Client {
	c.ctx = ctx
	return c.mockHttpClient
}

type mockConfigFactory struct {
	mockJWTConfig   *jwt.Config
	credentialsJSON string
	scopes          []string
}

func (params *mockConfigFactory) create(credentials []byte, scope ...string) (*jwt.Config, error) {
	params.credentialsJSON = string(credentials)
	params.scopes = scope
	return params.mockJWTConfig, nil
}

func TestNewGoogleConnection(t *testing.T) {
	// GIVEN
	credentialsJSON := `{"fake": "Credentials"}`
	credentials := bytes.NewBufferString(credentialsJSON).Bytes()
	scope := []string{
		"http://www.testscope.com/test1",
		"http://www.testscope.com/test2",
	}
	jwtConfig := jwt.Config{}
	httpClient := &http.Client{}
	httpClientFactoryGetter := &mockHttpClientFactoryGetter{mockHttpClient: httpClient}
	ctx := context.TODO()
	configFactory := &mockConfigFactory{
		mockJWTConfig: &jwtConfig,
	}
	context := GoogleContext{
		ConfigFactory:        configFactory.create,
		GetHttpClientFactory: httpClientFactoryGetter.get,
		Context:              ctx,
	}

	expected := map[string]interface{}{
		"credentialsJSON": credentialsJSON,
		"scope":           scope,
		"called":          true,
		"ctx":             ctx,
		"http.Client":     httpClient,
	}

	// WHEN
	cn, err := NewGoogleConnection(&context, credentials, scope...)
	if err != nil {
		t.Fatalf("TestGoogleDataAccess: Error %v", err)
	}
	actual := map[string]interface{}{
		"credentialsJSON": configFactory.credentialsJSON,
		"scope":           configFactory.scopes,
		"called":          httpClientFactoryGetter.called,
		"ctx":             httpClientFactoryGetter.ctx,
		"http.Client":     cn.(*connection).client,
	}
	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestGoogleDataAccess: expected %v, actual %v",
			expected,
			actual,
		)
	}
}

func TestGoogleContext(t *testing.T) {
	// GIVEN
	jwtConfigFactoryPtr := testutils.GetFnPtr(google.JWTConfigFromJSON)
	getHttpClientFactoryFnPtr := testutils.GetFnPtr(GetHttpClientFactory)
	ctx := context.Background()

	expected := map[string]interface{}{
		"ConfigFactory":        jwtConfigFactoryPtr,
		"GetHttpClientFactory": getHttpClientFactoryFnPtr,
		"ctx":                  ctx,
	}

	// WHEN
	googleContext := NewGoogleContext()

	actual := map[string]interface{}{
		"ConfigFactory": testutils.GetFnPtr(googleContext.ConfigFactory),

		"GetHttpClientFactory": testutils.GetFnPtr(googleContext.GetHttpClientFactory),
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
