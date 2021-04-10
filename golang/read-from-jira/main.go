package main

import (
	"log"
	"os"

	"local.dev/jira/internal/config"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func main() {
	jiraContextDependencies := config.NewJiraContextDependencies()
	jiraContext, err := config.NewJiraContext(jiraContextDependencies, os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}
	var issues []jira.Issue
	jiraContext.Connection.Execute(jiraContext.Query, &issues)

	jiraContext.Renderer.Render(os.Stdout, issues)
}
