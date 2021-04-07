package config

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"local.dev/jira/internal/operator"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

type JiraContext struct {
	Connection operator.Connection
	Query      operator.Query
	Renderer   operator.Renderer
}

func NewJiraContext() (*JiraContext, error) {
	url := os.Getenv("JIRA_URL")
	username := os.Getenv("JIRA_USER")
	password := os.Getenv("JIRA_PASSWORD")

	if url == "" || username == "" || password == "" {
		return nil, errors.New("require JIRA_URL, JIRA_USER, and JIRA_PASSWORD environment variales")
	}

	var maxResults int
	flag.IntVar(&maxResults, "maxResults", 500, "Maximum number of results to return.")
	log.Printf("MaxResults: %v", maxResults)

	flag.Parse()

	if flag.NArg() < 2 {
		return nil, errors.New("must provide template filename AND JQL")
	}

	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	filename := flag.Arg(0)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	templateText := string(contents)
	renderer := operator.NewTemplateRenderer(templateText)

	connection := operator.NewJiraConnection(
		tp.Client(),
		url,
		maxResults,
	)

	jql := flag.Arg(1)

	query := operator.NewJiraQuery(jql)

	return &JiraContext{
		Connection: connection,
		Renderer:   renderer,
		Query:      query,
	}, nil
}
