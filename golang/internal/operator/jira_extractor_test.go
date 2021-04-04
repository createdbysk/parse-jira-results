package operator

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"

	"local.dev/jira/internal/model"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func ticketsFixture() []model.Ticket {
	return []model.Ticket{
		{
			"name":        "TICKET-1",
			"priority":    "3-Low",
			"type":        "Task",
			"status":      "Open",
			"resolution":  "Unresolved1",
			"sprint":      "Sprint1",
			"createdDate": "2021-02-28T00:00:00Z",
			"startDate":   "2021-03-01T00:00:00Z",
			"endDate":     "",
		},
		{
			"name":        "TICKET-2",
			"priority":    "2-Medium",
			"type":        "Story",
			"status":      "InProgress",
			"resolution":  "Unresolved2",
			"sprint":      "Sprint2",
			"createdDate": "2021-02-25T00:00:00Z",
			"startDate":   "2021-02-26T00:00:00Z",
			"endDate":     "",
		},
		{
			"name":        "TICKET-3",
			"priority":    "1-High",
			"type":        "Task",
			"status":      "Done",
			"resolution":  "Resolved",
			"sprint":      "Sprint4",
			"createdDate": "2021-02-20T00:00:00Z",
			"startDate":   "2021-02-21T00:00:00Z",
			"endDate":     "2021-02-22T00:00:00Z",
		},
	}
}

type mockHandler struct {
	jql string
}

var (
	jiraTimeFormat = "2006-01-02T15:04:05Z"
)
func (h *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jql = r.URL.Query()["jql"]
	if jql == h.jql {
		issue := []jira.Issue{
			{
				Key: "TICKET-1",
				Fields: &jira.IssueFields{
					Type: jira.IssueType{
						Name: "Task",
					},
				},
				Resolution: &jira.Resolution{
					Name: "Unresolved1",
				},
				Status: &jira.Status{
					Name: "Open",
				},
				Sprint: &jira.Sprint{
					Name: "Sprint1",
				},
				Created: time.Parse(
					jiraTimeFormat,
					"2021-02-28T00:00:00Z",
				),
				ChangeLog: &jira.ChangeLog{
					Histories: []jire.ChangelogHistory{
						{
							Created: time.Parse(
								jiraTimeFormat,
								"2021-03-01T00:00:00Z",
							),
							Items: []jira.ChangelogItems {
								Field: "status",
								FromString: "To Do",
								ToString: "In Progress",
							}
						}
					}
				}
			}
		}
		// {
		// 	"name":        "TICKET-2",
		// 	"priority":    "2-Medium",
		// 	"type":        "Story",
		// 	"status":      "InProgress",
		// 	"resolution":  "Unresolved2",
		// 	"sprint":      "Sprint2",
		// 	"createdDate": "2021-02-25T00:00:00Z",
		// 	"startDate":   "2021-02-26T00:00:00Z",
		// 	"endDate":     "",
		// },
		// {
		// 	"name":        "TICKET-3",
		// 	"priority":    "1-High",
		// 	"type":        "Task",
		// 	"status":      "Done",
		// 	"resolution":  "Resolved",
		// 	"sprint":      "Sprint4",
		// 	"createdDate": "2021-02-20T00:00:00Z",
		// 	"startDate":   "2021-02-21T00:00:00Z",
		// 	"endDate":     "2021-02-22T00:00:00Z",
		// },
	}

}

func mockJiraConnection() {
	// http://localhost:8080/rest/api/2/search?jql=assignee=charlie
	http.HandleFunc("/rest/api/2/search", )
	mockServer := httptest.NewSer
}

func TestJiraExtractor(t *testing.T) {
	// GIVEN
	jql := "project=TEST"

	connection := mockJiraConnection(url, username, password)
	query := mockJiraQuery(jql)

	expected := model.TicketsFixture()

	// WHEN
	extractor, errExtractor := NewExtractor(connection, query)
	if errExtractor != nil {
		t.Errorf("NewExtractor() failed with error: %v", errExtractor)
	}
	actual := make([]model.Ticket, 0, len(expected))
	for ticket, valid := extractor.extract(); valid; {
		actual = append(actual, ticket)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"NewJiraExtractor() expected: %v actual: %v",
			expected, actual,
		)
	}

}
