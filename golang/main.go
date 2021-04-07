package main

import (
	"log"
	"os"

	"local.dev/jira/internal/config"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func main() {
	jiraContext, err := config.NewJiraContext()

	if err != nil {
		log.Fatal(err)
	}
	var issues []jira.Issue
	jiraContext.Connection.Execute(jiraContext.Query, &issues)

	jiraContext.Renderer.Render(os.Stdout, issues)
}
