package operator

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"local.dev/jira/internal/model"
)

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
		w.Write(jiraSearchResultsFixture())
	}
}

func newMockServer() *httptest.Server {
	// http://localhost:8080/rest/api/2/search?jql=project=TEST
	jql := "project=TEST"
	mux := http.NewServeMux()
	mux.Handle("/rest/api/2/search", &mockHandler{jql})
	mockServer := httptest.NewServer(mux)
	return mockServer
}

type mockJiraConnection struct {
	client *http.Client
}

func newMockJiraConnection(s *httptest.Server) Connection {
	client := s.Client()
	return &mockJiraConnection{client}
}

type mockJiraQuery struct {
	jql string
}

func newMockJiraQuery(jql string) Query {
	return &mockJiraQuery{jql}
}

func TestJiraExtractor(t *testing.T) {
	// GIVEN
	var ticket model.Ticket
	jql := "project=TEST"
	mockServer := newMockServer()
	defer mockServer.Close()
	connection := newMockJiraConnection(mockServer)
	query := newMockJiraQuery(jql)

	expected := ticketsFixture()

	// WHEN
	extractor, errExtractor := NewExtractor(connection)
	if errExtractor != nil {
		t.Fatalf("NewExtractor() failed with error: %v", errExtractor)
	}
	actual := make([]model.Ticket, 0, len(expected))
	for it := extractor.Extract(query); it.Next(&ticket); {
		actual = append(actual, ticket)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"NewJiraExtractor() expected: %v actual: %v",
			expected, actual,
		)
	}

}
