package operator

import (
	"bytes"
	"context"
	"net/http"
	"reflect"
	"testing"

	"golang.org/x/oauth2/jwt"
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
	scope           string
}

func (params *mockConfigFactory) create(credentials []byte, scope string) *jwt.Config {
	params.credentialsJSON = string(credentials)
	params.scope = scope
	return params.mockJWTConfig
}

func TestNewGoogleConnection(t *testing.T) {
	// GIVEN
	credentialsJSON := `{"fake": "Credentials"}`
	credentials := bytes.NewBufferString(credentialsJSON).Bytes()
	scope := "http://www.testscope.com/test"
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
	cn, err := NewGoogleConnection(&context, credentials, scope)
	if err != nil {
		t.Fatalf("TestGoogleDataAccess: Error %v", err)
	}
	actual := map[string]interface{}{
		"credentialsJSON": configFactory.credentialsJSON,
		"scope":           configFactory.scope,
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
