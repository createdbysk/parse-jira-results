package config

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"reflect"
	"testing"

	"local.dev/jira/internal/operator"
)

// func TestparseJiraCommandLineParser(t *testing.T) {
// 	// GIVEN
// 	testcases := []struct {
// 		name     string
// 		args     []string
// 		params   jiraCommandlineParams
// 		hasError bool
// 	}{
// 		{
// 			"Default MaxResults",
// 			[]string{
// 				"test.tpl",
// 				"project=TEST",
// 			},
// 			jiraCommandlineParams{
// 				MaxResults:       defaultMaxResults,
// 				TemplateFilename: "test.tpl",
// 				JQL:              "project=TEST",
// 			},
// 			false,
// 		},
// 		{
// 			"MaxResults=3",
// 			[]string{
// 				"-maxResults",
// 				"3",
// 				"test.tpl",
// 				"project=TEST",
// 			},
// 			jiraCommandlineParams{
// 				MaxResults:       3,
// 				TemplateFilename: "test.tpl",
// 				JQL:              "project=TEST",
// 			},
// 			false,
// 		},
// 	}

// 	for _, testcase := range testcases {
// 		t.Run(
// 			testcase.name,
// 			func(t *testing.T) {
// 				// GIVEN
// 				// testcase
// 				expected := map[string]interface{}{
// 					"params":   testcase.params,
// 					"hasError": testcase.hasError,
// 				}

// 				// WHEN
// 				params, err := parsejiraCommandlineParams(testcase.args)
// 				actual := map[string]interface{}{
// 					"params":   *params,
// 					"hasError": err != nil,
// 				}
// 				// THEN

// 				if !reflect.DeepEqual(actual, expected) {
// 					t.Errorf(
// 						"%v: expected %v, actual %v",
// 						testcase.name,
// 						expected,
// 						actual,
// 					)
// 				}
// 			},
// 		)
// 	}
// }

func TestNewJiraContext(t *testing.T) {
	// GIVEN
	var maxResults int
	envVars := map[string]string{
		"JIRA_URL":      "http://jira",
		"JIRA_USER":     "user",
		"JIRA_PASSWORD": "password",
	}

	client := http.DefaultClient
	jql := "project=TEST"
	filename := "test.tpl"
	templateText := "Hello"
	query := &operator.MockQuery{}
	renderer := &operator.MockRenderer{}
	connection := &operator.MockConnection{}
	di := &JiraContextDependencies{
		func(username string, password string) *http.Client {
			if username == envVars["JIRA_USER"] && password == envVars["JIRA_PASSWORD"] {
				return client
			} else {
				return nil
			}
		},
		func(c *http.Client, url string, m int) operator.Connection {
			if c == client && url == envVars["JIRA_URL"] {
				maxResults = m
				return connection
			} else {
				return nil
			}
		},
		func(f string) ([]byte, error) {
			if f == filename {
				return bytes.NewBufferString(templateText).Bytes(), nil
			} else {
				return nil, errors.New("Invalid filename")
			}
		},
		func(q string) operator.Query {
			if q == jql {
				return query
			} else {
				return nil
			}
		},
		func(text string) operator.Renderer {
			if text == templateText {
				return renderer
			} else {
				return nil
			}
		},
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
