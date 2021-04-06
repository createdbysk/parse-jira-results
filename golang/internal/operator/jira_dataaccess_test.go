package operator

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	jira "gopkg.in/andygrunwald/go-jira.v1"
	"local.dev/jira/internal/model"
)

func jqlFixture() string {
	return "project=Test"
}

func ticketsFixture() []model.Ticket {
	return []model.Ticket{
		{
			"name":        "TEST-24",
			"priority":    "4-Low",
			"type":        "Task",
			"status":      "TO DO",
			"resolution":  "Unresolved",
			"sprint":      "",
			"createdDate": "2021-04-04T23:46:56Z",
			"startDate":   "",
			"endDate":     "",
		},
		{
			"name":        "TEST-25",
			"priority":    "3-Medium",
			"type":        "Story",
			"status":      "In Progress",
			"resolution":  "Unresolved",
			"sprint":      "",
			"createdDate": "2021-04-04T23:47:13Z",
			"startDate":   "2021-04-04T23:47:16Z",
			"endDate":     "",
		},
		{
			"name":        "TEST-26",
			"priority":    "2-High",
			"type":        "Task",
			"status":      "Done",
			"resolution":  "Done",
			"sprint":      "",
			"createdDate": "2021-04-04T23:47:26Z",
			"startDate":   "2021-04-04T23:47:30Z",
			"endDate":     "2021-04-04T23:47:35Z",
		},
	}
}

type mockHandler struct {
	jql string
}

func (h *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// http://localhost:58080/rest/api/2/search?jql=project=TEST&expand=changelog&maxResults=3
	jql := r.URL.Query()["jql"][0]
	expand := r.URL.Query()["expand"][0]
	if jql == h.jql && strings.ToLower(expand) == "changelog" {
		w.Write(jiraSearchResultsJsonFixture())
	}
}

func newMockServer() *httptest.Server {
	// http://localhost:8080/rest/api/2/search?jql=project=TEST
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
	connection := NewJiraConnection(s.Client(), s.URL, 3)
	query := NewJiraQuery(jql)
	expected := jiraSearchResultsIssuesFixture()

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
