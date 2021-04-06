package operator

import (
	"net/http"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

type connection struct {
	jiraClient *jira.Client
	options    *jira.SearchOptions
}

func NewJiraConnection(client *http.Client, url string, maxResults int) Connection {
	options := jira.SearchOptions{
		MaxResults: 3,
		Expand:     "changelog",
	}
	jiraClient, _ := jira.NewClient(client, url)
	return &connection{jiraClient, &options}
}

type query struct {
	jql string
}

func NewJiraQuery(jql string) Query {
	return &query{jql}
}

func (q *query) Get(result interface{}) error {
	*(result.(*string)) = q.jql
	return nil
}

func (c *connection) Execute(q Query, data interface{}) {
	var jql string
	q.Get(&jql)

	issues, _, _ := c.jiraClient.Issue.Search(jql, c.options)
	*(data.(*[]jira.Issue)) = issues
}
