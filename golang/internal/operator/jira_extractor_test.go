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

type mockConnection struct {
	called bool
}

type mockIterator struct {
	more bool
}

func (it *mockIterator) Next(value interface{}) bool {
	it.more = !it.more
	*(value.(*bool)) = it.more
	return it.more
}

type mockQuery struct {
	it Iterator
}

func (q *mockQuery) Get(value interface{}) error {
	*(value.(*Iterator)) = q.it
	return nil
}

func (c *mockConnection) Execute(q Query) Iterator {
	var it Iterator
	q.Get(&it)
	c.called = true
	return it
}

func TestExtractor(t *testing.T) {
	// GIVEN
	connection := &mockConnection{}
	iterator := &mockIterator{}
	query := &mockQuery{iterator}

	// WHEN
	var b bool
	var once bool
	var itMemo Iterator
	extractor := NewExtractor(connection)

	for it := extractor.Extract(query); it.Next(&b); {
		once = true
		itMemo = it
	}
	if !once || itMemo != iterator {
		t.Errorf(
			"TestExtractor() expected once, iterator: %v, %v, actual once, iterator: %v, %v",
			true, once,
			iterator, itMemo,
		)
	}
}

func TestMockServer(t *testing.T) {
	// GIVEN
	jql := jqlFixture()
	s := newMockServer()
	jiraClient, _ := jira.NewClient(s.Client(), s.URL)
	options := jira.SearchOptions{
		MaxResults: 3,
		Expand:     "changelog",
	}

	expected := jiraSearchResultsIssuesFixture()

	// WHEN
	actual, _, _ := jiraClient.Issue.Search(jql, &options)

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestMockServer: expected: %v, actual: %v",
			expected,
			actual,
		)
	}
}
