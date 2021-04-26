run-jira:
	-docker volume create --name jiraVolume
	-docker run -v jiraVolume:/var/atlassian/application-data/jira --name="jira" -d -p 58080:8080 atlassian/jira-core

start-jira:
	docker start jira