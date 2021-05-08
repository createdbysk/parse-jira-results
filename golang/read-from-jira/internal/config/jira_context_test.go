package config

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"

	"local.dev/jira/internal/operator"
	"local.dev/jira/internal/testutils"
)

func TestNewJiraBasicAuth(t *testing.T) {
	// GIVEN
	username := "username"
	password := "password"

	// WHEN
	client := NewJiraBasicAuth(username, password)

	// THEN
	if client == nil {
		t.Errorf("NewBasicAuth() did not return valid http client.")
	}
}

func TestNewJiraContextDependencies(t *testing.T) {
	// GIVEN
	expected := map[string]uintptr{
		"NewJiraBasicAuthFnPtr":    testutils.GetFnPtr(NewJiraBasicAuth),
		"NewJiraConnectionFnPtr":   testutils.GetFnPtr(operator.NewJiraConnection),
		"ReadFileFnPtr":            testutils.GetFnPtr(ioutil.ReadFile),
		"NewJiraQueryFnPtr":        testutils.GetFnPtr(operator.NewJiraQuery),
		"NewTemplateRendererFnPtr": testutils.GetFnPtr(operator.NewTemplateRenderer),
	}

	// WHEN
	jiraContextDependencies := NewJiraContextDependencies()
	actual := map[string]uintptr{
		"NewJiraBasicAuthFnPtr":    testutils.GetFnPtr(jiraContextDependencies.NewHttpClient),
		"NewJiraConnectionFnPtr":   testutils.GetFnPtr(jiraContextDependencies.NewJiraConnection),
		"ReadFileFnPtr":            testutils.GetFnPtr(jiraContextDependencies.ReadFile),
		"NewJiraQueryFnPtr":        testutils.GetFnPtr(jiraContextDependencies.NewJiraQuery),
		"NewTemplateRendererFnPtr": testutils.GetFnPtr(jiraContextDependencies.NewTemplateRenderer),
	}

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"%v test: expected %v, actual %v",
			"NewJiraContextDependencies()",
			expected,
			actual,
		)
	}
}

func TestNewJiraContext(t *testing.T) {
	// GIVEN
	var maxResults int
	envVars := map[string]string{
		"JIRA_URL":      "http://jira",
		"JIRA_USERNAME": "user",
		"JIRA_PASSWORD": "password",
	}

	client := http.DefaultClient
	jql := "project=TEST"
	filename := "test.tpl"
	templateText := "Hello"
	query := &operator.MockQuery{}
	renderer := &operator.MockRenderer{}
	connection := &operator.MockConnection{}
	mockClientFactory := func(username string, password string) *http.Client {
		if username == envVars["JIRA_USERNAME"] && password == envVars["JIRA_PASSWORD"] {
			return client
		} else {
			return nil
		}
	}
	mockConnectionFactory := func(c *http.Client, url string, m int) operator.Connection {
		if c == client && url == envVars["JIRA_URL"] {
			maxResults = m
			return connection
		} else {
			return nil
		}
	}
	mockReadFile := func(f string) ([]byte, error) {
		if f == filename {
			return bytes.NewBufferString(templateText).Bytes(), nil
		} else {
			return nil, errors.New("Invalid filename")
		}
	}
	mockQueryFactory := func(q string) operator.Query {
		if q == jql {
			return query
		} else {
			return nil
		}
	}
	mockRendererFactory := func(text string) operator.Renderer {
		if text == templateText {
			return renderer
		} else {
			return nil
		}
	}

	di := &JiraContextDependencies{
		mockClientFactory,
		mockConnectionFactory,
		mockReadFile,
		mockQueryFactory,
		mockRendererFactory,
	}

	testcases := []struct {
		name       string
		args       []string
		envVars    map[string]string
		maxResults int
		hasError   bool
	}{
		{
			"Default MaxResults",
			[]string{
				filename,
				jql,
			},
			envVars,
			defaultMaxResults,
			false,
		},
		{
			"MaxResults=3",
			[]string{
				"-maxResults",
				"3",
				filename,
				jql,
			},
			envVars,
			3,
			false,
		},
	}

	for _, testcase := range testcases {
		t.Run(
			testcase.name,
			func(t *testing.T) {
				// GIVEN
				for key, value := range testcase.envVars {
					os.Setenv(key, value)
				}
				expected := map[string]interface{}{
					"maxResults": testcase.maxResults,
					"context": JiraContext{
						Connection: connection,
						Query:      query,
						Renderer:   renderer,
					},
					"hasError": testcase.hasError,
				}

				// WHEN
				context, err := NewJiraContext(di, testcase.args)
				actual := map[string]interface{}{
					"maxResults": maxResults,
					"context":    *context,
					"hasError":   err != nil,
				}

				// THEN
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(actual, expected) {
					t.Errorf(
						"%v: expected %v, actual %v",
						testcase.name,
						expected,
						actual,
					)
				}
			},
		)
	}
}
