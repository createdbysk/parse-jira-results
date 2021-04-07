package config

import (
	"errors"
	"flag"
	"net/http"
	"os"

	"local.dev/jira/internal/operator"
)

const (
	defaultMaxResults = 500
)

type jiraCommandlineParams struct {
	MaxResults       int
	TemplateFilename string
	JQL              string
}

func parseJiraCommandlineArgs(args []string) (*jiraCommandlineParams, error) {
	var params jiraCommandlineParams
	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	flagSet.IntVar(&params.MaxResults, "maxResults", defaultMaxResults, "Maximum number of results to return.")

	flagSet.Parse(args)
	if flagSet.NArg() < 2 {
		return nil, errors.New("must provide template filename AND JQL")
	}

	params.TemplateFilename = flagSet.Arg(0)
	params.JQL = flagSet.Arg(1)

	return &params, nil
}

type JiraContextDependencies struct {
	NewHttpClient       func(username string, password string) *http.Client
	NewJiraConnection   func(client *http.Client, url string, maxResults int) operator.Connection
	ReadFile            func(filename string) ([]byte, error)
	NewJiraQuery        func(jql string) operator.Query
	NewTemplateRenderer func(text string) operator.Renderer
}

type JiraContext struct {
	Connection operator.Connection
	Query      operator.Query
	Renderer   operator.Renderer
}

func NewJiraContext(di *JiraContextDependencies, args []string) (*JiraContext, error) {
	url := os.Getenv("JIRA_URL")
	username := os.Getenv("JIRA_USER")
	password := os.Getenv("JIRA_PASSWORD")

	if url == "" || username == "" || password == "" {
		return nil, errors.New("require JIRA_URL, JIRA_USER, and JIRA_PASSWORD environment variales")
	}

	commandlineArgs, err := parseJiraCommandlineArgs(args)

	if err != nil {
		return nil, err
	}

	client := di.NewHttpClient(username, password)

	contents, err := di.ReadFile(commandlineArgs.TemplateFilename)
	if err != nil {
		return nil, err
	}

	templateText := string(contents)
	renderer := di.NewTemplateRenderer(templateText)

	connection := di.NewJiraConnection(
		client,
		url,
		commandlineArgs.MaxResults,
	)

	query := di.NewJiraQuery(commandlineArgs.JQL)

	return &JiraContext{
		Connection: connection,
		Renderer:   renderer,
		Query:      query,
	}, nil
}
