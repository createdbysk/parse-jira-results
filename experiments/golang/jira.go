package main

import (
	"fmt"
	"log"
	"os"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func main() {
	username := os.Getenv("JIRA_USERNAME")
	password := os.Getenv("JIRA_PASSWORD")
	url := os.Getenv("JIRA_URL")
	filterName := os.Getenv("JIRA_FILTERNAME")
	if username == "" || password == "" || url == "" || filterName == "" {
		log.Fatal("You must define the JIRA_USERNAME, JIRA_PASSWORD, JIRA_URL, AND JIRA_FILTERNAME environment variables.")
	}
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}
	jiraClient, _ := jira.NewClient(tp.Client(), url)
	filters, _, err := jiraClient.Filter.GetFavouriteList()
	if err != nil {
		log.Fatal(err)
	}
	for _, filter := range filters {
		if filter.Name == filterName {
			options := jira.SearchOptions{
				MaxResults: 5,
				Expand:     "changelog",
			}
			issues, _, err := jiraClient.Issue.Search(filter.Jql, &options)
			if err != nil {
				log.Fatalf("Failed %s", err)
			}

			for _, issue := range issues {
				fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
				fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
				fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
				fmt.Println("========================================================")
			}
			return
		}
	}
	log.Printf("Filter %s not found.", filterName)
}
