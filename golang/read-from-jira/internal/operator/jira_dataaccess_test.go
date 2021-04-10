package operator

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func jqlFixture() string {
	return "project=Test"
}

type mockHandler struct {
	jql string
}

func (h *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// http://localhost:58080/rest/api/2/search?jql=project=TEST&expand=changelog&maxResults=x
	jql := r.URL.Query()["jql"][0]
	maxResults, _ := strconv.Atoi(r.URL.Query()["maxResults"][0])
	expand := r.URL.Query()["expand"][0]
	if jql == h.jql && strings.ToLower(expand) == "changelog" {
		results := jiraSearchResultsJsonFixture(maxResults)
		w.Write(results)
	}
}

func newMockServer() *httptest.Server {
	// http://localhost:8080/rest/api/2/search?jql=project=TEST&maxResults=x
	jql := jqlFixture()
	mux := http.NewServeMux()
	mux.Handle("/rest/api/2/search", &mockHandler{jql})
	mockServer := httptest.NewServer(mux)
	return mockServer
}

func TestJiraConnection(t *testing.T) {
	// GIVEN
	jql := jqlFixture()
	s := newMockServer()
	query := NewJiraQuery(jql)
	for i := 1; i < 3; i++ {
		connection := NewJiraConnection(s.Client(), s.URL, i)
		expected := jiraSearchResultsIssuesFixture(i)

		// WHEN
		var actual []jira.Issue
		connection.Execute(query, &actual)

		// THEN
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf(
				"TestJiraConnection: expected: %v, actual: %v",
				expected,
				actual,
			)
		}
	}
}
