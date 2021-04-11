package operator

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

type mockClientFactory struct {
	called     bool
	mockClient *http.Client
}

func (config *mockClientFactory) Client() *http.Client {
	mockClientFactory.called = true
	return config.mockClient
}

type mockClientFactoryParameters struct {
	mockConfig      *mockClientFactory
	credentialsJSON string
	scope           string
}

func (params *mockClientFactoryParameters) create(credentials []byte, scope string) googleClientFactory {
	params.credentialsJSON = string(credentials)
	params.scope = scope
	return params.mockConfig
}

func TestGoogleDataAccess(t *testing.T) {
	// GIVEN
	credentialsJSON := `{"fake": "Credentials"}`
	credentials := bytes.NewBufferString(credentialsJSON).Bytes()
	scope := "http://www.testscope.com/test"
	mockClient := http.DefaultClient
	mockFactory := mockClientFactory{mockClient: mockClient}
	factoryParams := &mockClientFactoryParameters{
		mockConfig: &mockFactory,
	}
	context := GoogleContext{
		GoogleClientFactoryFactory: factoryParams.create,
	}

	expected := map[string]interface{}{
		"credentialsJSON": credentialsJSON,
		"scope":           scope,
		"called":          true,
		"client":          mockClient,
	}

	// WHEN
	connection := NewGoogleConnection(&context)
	actual := map[string]interface{}{
		"credentialsJSON": factoryParams.credentialsJSON,
		"scope":           factoryParams.scope,
		"called":          mockFactory.called,
		"client":          mockFactory.mockClient,
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
