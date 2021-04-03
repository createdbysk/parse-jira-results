package operator

import (
	"bytes"
	"encoding/json"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func jiraSearchResultsIssuesFixture(maxResults int) []jira.Issue {
	var m map[string]interface{}
	var issues []jira.Issue
	results := jiraSearchResultsJsonFixture(maxResults)

	// Extract the issues field out of the json
	// Then convert that issues field into []jira.Issue.
	json.Unmarshal(results, &m)
	i, _ := json.Marshal(m["issues"])
	json.Unmarshal(i, &issues)
	return issues
}

func jiraSearchResultsJsonFixture(maxResults int) []byte {
	// Extract the issues field out of the json
	// Get maxResults number of items from that slice.
	// Marshal it back to json with maxResults issues.
	var m map[string]interface{}
	results := jiraAllResultsJsonFixture()
	json.Unmarshal(results, &m)
	m["issues"] = m["issues"].([]interface{})[:maxResults]
	response, _ := json.Marshal(m)
	return response
}

func jiraAllResultsJsonFixture() []byte {
	return bytes.NewBufferString(`{
		"expand": "schema,names",
		"startAt": 0,
		"maxResults": 3,
		"total": 26,
		"issues": [
			{
				"expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
				"id": "10602",
				"self": "http://localhost:58080/rest/api/2/issue/10602",
				"key": "TEST-26",
				"fields": {
					"issuetype": {
						"self": "http://localhost:58080/rest/api/2/issuetype/10000",
						"id": "10000",
						"description": "A task that needs to be done.",
						"iconUrl": "http://localhost:58080/secure/viewavatar?size=xsmall&avatarId=10318&avatarType=issuetype",
						"name": "Task",
						"subtask": false,
						"avatarId": 10318
					},
					"components": [],
					"timespent": null,
					"timeoriginalestimate": null,
					"description": "Test 1",
					"project": {
						"self": "http://localhost:58080/rest/api/2/project/10001",
						"id": "10001",
						"key": "TEST",
						"name": "Test",
						"projectTypeKey": "business",
						"avatarUrls": {
							"48x48": "http://localhost:58080/secure/projectavatar?avatarId=10324",
							"24x24": "http://localhost:58080/secure/projectavatar?size=small&avatarId=10324",
							"16x16": "http://localhost:58080/secure/projectavatar?size=xsmall&avatarId=10324",
							"32x32": "http://localhost:58080/secure/projectavatar?size=medium&avatarId=10324"
						}
					},
					"fixVersions": [],
					"customfield_10110": null,
					"customfield_10111": null,
					"aggregatetimespent": null,
					"resolution": {
						"self": "http://localhost:58080/rest/api/2/resolution/10000",
						"id": "10000",
						"description": "Work has been completed on this issue.",
						"name": "Done"
					},
					"customfield_10104": null,
					"customfield_10105": "0|i0005j:",
					"customfield_10106": "{summaryBean=com.atlassian.jira.plugin.devstatus.rest.SummaryBean@7f413ea8[summary={pullrequest=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@1cb81cd5[overall=PullRequestOverallBean{stateCount=0, state='OPEN', details=PullRequestOverallDetails{openCount=0, mergedCount=0, declinedCount=0}},byInstanceType={}], build=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@8a2e558[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BuildOverallBean@54a00e2d[failedBuildCount=0,successfulBuildCount=0,unknownBuildCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], review=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@16fb2561[overall=com.atlassian.jira.plugin.devstatus.summary.beans.ReviewsOverallBean@69e665cc[stateCount=0,state=<null>,dueDate=<null>,overDue=false,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], deployment-environment=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@7671d58d[overall=com.atlassian.jira.plugin.devstatus.summary.beans.DeploymentOverallBean@22f47c08[topEnvironments=[],showProjects=false,successfulCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], repository=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@6d6b951a[overall=com.atlassian.jira.plugin.devstatus.summary.beans.CommitOverallBean@3779457e[count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], branch=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@4fdc23a1[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BranchOverallBean@4ec9026e[count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}]},errors=[],configErrors=[]], devSummaryJson={\"cachedValue\":{\"errors\":[],\"configErrors\":[],\"summary\":{\"pullrequest\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"stateCount\":0,\"state\":\"OPEN\",\"details\":{\"openCount\":0,\"mergedCount\":0,\"declinedCount\":0,\"total\":0},\"open\":true},\"byInstanceType\":{}},\"build\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"failedBuildCount\":0,\"successfulBuildCount\":0,\"unknownBuildCount\":0},\"byInstanceType\":{}},\"review\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"stateCount\":0,\"state\":null,\"dueDate\":null,\"overDue\":false,\"completed\":false},\"byInstanceType\":{}},\"deployment-environment\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"topEnvironments\":[],\"showProjects\":false,\"successfulCount\":0},\"byInstanceType\":{}},\"repository\":{\"overall\":{\"count\":0,\"lastUpdated\":null},\"byInstanceType\":{}},\"branch\":{\"overall\":{\"count\":0,\"lastUpdated\":null},\"byInstanceType\":{}}}},\"isStale\":false}}",
					"customfield_10107": null,
					"customfield_10108": null,
					"aggregatetimeestimate": null,
					"customfield_10109": null,
					"resolutiondate": "2021-04-04T23:47:35.000+0000",
					"workratio": -1,
					"summary": "Test 3",
					"lastViewed": "2021-04-05T00:00:19.714+0000",
					"watches": {
						"self": "http://localhost:58080/rest/api/2/issue/TEST-26/watchers",
						"watchCount": 1,
						"isWatching": true
					},
					"creator": {
						"self": "http://localhost:58080/rest/api/2/user?username=admin",
						"name": "admin",
						"key": "JIRAUSER10000",
						"emailAddress": "skumar@datto.com",
						"avatarUrls": {
							"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
							"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
							"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
							"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
						},
						"displayName": "SATISH KUMAR",
						"active": true,
						"timeZone": "America/Chicago"
					},
					"subtasks": [],
					"created": "2021-04-04T23:47:26.000+0000",
					"reporter": {
						"self": "http://localhost:58080/rest/api/2/user?username=admin",
						"name": "admin",
						"key": "JIRAUSER10000",
						"emailAddress": "skumar@datto.com",
						"avatarUrls": {
							"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
							"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
							"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
							"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
						},
						"displayName": "SATISH KUMAR",
						"active": true,
						"timeZone": "America/Chicago"
					},
					"aggregateprogress": {
						"progress": 0,
						"total": 0
					},
					"priority": {
						"self": "http://localhost:58080/rest/api/2/priority/2",
						"iconUrl": "http://localhost:58080/images/icons/priorities/high.svg",
						"name": "High",
						"id": "2"
					},
					"customfield_10001": null,
					"customfield_10100": null,
					"labels": [],
					"environment": null,
					"timeestimate": null,
					"aggregatetimeoriginalestimate": null,
					"versions": [],
					"duedate": null,
					"progress": {
						"progress": 0,
						"total": 0
					},
					"issuelinks": [
						{
							"id": "10001",
							"self": "http://localhost:58080/rest/api/2/issueLink/10001",
							"type": {
								"id": "10001",
								"name": "Cloners",
								"inward": "is cloned by",
								"outward": "clones",
								"self": "http://localhost:58080/rest/api/2/issueLinkType/10001"
							},
							"outwardIssue": {
								"id": "10601",
								"key": "TEST-25",
								"self": "http://localhost:58080/rest/api/2/issue/10601",
								"fields": {
									"summary": "Test 2",
									"status": {
										"self": "http://localhost:58080/rest/api/2/status/3",
										"description": "This issue is being actively worked on at the moment by the assignee.",
										"iconUrl": "http://localhost:58080/images/icons/statuses/inprogress.png",
										"name": "In Progress",
										"id": "3",
										"statusCategory": {
											"self": "http://localhost:58080/rest/api/2/statuscategory/4",
											"id": 4,
											"key": "indeterminate",
											"colorName": "yellow",
											"name": "In Progress"
										}
									},
									"priority": {
										"self": "http://localhost:58080/rest/api/2/priority/3",
										"iconUrl": "http://localhost:58080/images/icons/priorities/medium.svg",
										"name": "Medium",
										"id": "3"
									},
									"issuetype": {
										"self": "http://localhost:58080/rest/api/2/issuetype/10000",
										"id": "10000",
										"description": "A task that needs to be done.",
										"iconUrl": "http://localhost:58080/secure/viewavatar?size=xsmall&avatarId=10318&avatarType=issuetype",
										"name": "Task",
										"subtask": false,
										"avatarId": 10318
									}
								}
							}
						}
					],
					"votes": {
						"self": "http://localhost:58080/rest/api/2/issue/TEST-26/votes",
						"votes": 0,
						"hasVoted": false
					},
					"assignee": {
						"self": "http://localhost:58080/rest/api/2/user?username=admin",
						"name": "admin",
						"key": "JIRAUSER10000",
						"emailAddress": "skumar@datto.com",
						"avatarUrls": {
							"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
							"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
							"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
							"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
						},
						"displayName": "SATISH KUMAR",
						"active": true,
						"timeZone": "America/Chicago"
					},
					"updated": "2021-04-05T00:00:19.000+0000",
					"status": {
						"self": "http://localhost:58080/rest/api/2/status/10001",
						"description": "",
						"iconUrl": "http://localhost:58080/images/icons/status_generic.gif",
						"name": "Done",
						"id": "10001",
						"statusCategory": {
							"self": "http://localhost:58080/rest/api/2/statuscategory/3",
							"id": 3,
							"key": "done",
							"colorName": "green",
							"name": "Done"
						}
					}
				},
				"changelog": {
					"startAt": 0,
					"maxResults": 4,
					"total": 4,
					"histories": [
						{
							"id": "10203",
							"author": {
								"self": "http://localhost:58080/rest/api/2/user?username=admin",
								"name": "admin",
								"key": "JIRAUSER10000",
								"emailAddress": "skumar@datto.com",
								"avatarUrls": {
									"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
									"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
									"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
									"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
								},
								"displayName": "SATISH KUMAR",
								"active": true,
								"timeZone": "America/Chicago"
							},
							"created": "2021-04-04T23:47:26.394+0000",
							"items": [
								{
									"field": "Link",
									"fieldtype": "jira",
									"from": null,
									"fromString": null,
									"to": "TEST-25",
									"toString": "This issue clones TEST-25"
								}
							]
						},
						{
							"id": "10205",
							"author": {
								"self": "http://localhost:58080/rest/api/2/user?username=admin",
								"name": "admin",
								"key": "JIRAUSER10000",
								"emailAddress": "skumar@datto.com",
								"avatarUrls": {
									"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
									"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
									"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
									"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
								},
								"displayName": "SATISH KUMAR",
								"active": true,
								"timeZone": "America/Chicago"
							},
							"created": "2021-04-04T23:47:30.467+0000",
							"items": [
								{
									"field": "status",
									"fieldtype": "jira",
									"from": "10000",
									"fromString": "To Do",
									"to": "3",
									"toString": "In Progress"
								}
							]
						},
						{
							"id": "10206",
							"author": {
								"self": "http://localhost:58080/rest/api/2/user?username=admin",
								"name": "admin",
								"key": "JIRAUSER10000",
								"emailAddress": "skumar@datto.com",
								"avatarUrls": {
									"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
									"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
									"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
									"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
								},
								"displayName": "SATISH KUMAR",
								"active": true,
								"timeZone": "America/Chicago"
							},
							"created": "2021-04-04T23:47:35.372+0000",
							"items": [
								{
									"field": "resolution",
									"fieldtype": "jira",
									"from": null,
									"fromString": null,
									"to": "10000",
									"toString": "Done"
								},
								{
									"field": "status",
									"fieldtype": "jira",
									"from": "3",
									"fromString": "In Progress",
									"to": "10001",
									"toString": "Done"
								}
							]
						},
						{
							"id": "10207",
							"author": {
								"self": "http://localhost:58080/rest/api/2/user?username=admin",
								"name": "admin",
								"key": "JIRAUSER10000",
								"emailAddress": "skumar@datto.com",
								"avatarUrls": {
									"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
									"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
									"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
									"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
								},
								"displayName": "SATISH KUMAR",
								"active": true,
								"timeZone": "America/Chicago"
							},
							"created": "2021-04-05T00:00:19.592+0000",
							"items": [
								{
									"field": "priority",
									"fieldtype": "jira",
									"from": "4",
									"fromString": "Low",
									"to": "2",
									"toString": "High"
								}
							]
						}
					]
				}
			},
			{
				"expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
				"id": "10601",
				"self": "http://localhost:58080/rest/api/2/issue/10601",
				"key": "TEST-25",
				"fields": {
					"issuetype": {
						"self": "http://localhost:58080/rest/api/2/issuetype/10001",
						"id": "10001",
						"description": "A story.",
						"iconUrl": "http://localhost:58080/secure/viewavatar?size=xsmall&avatarId=10318&avatarType=issuetype",
						"name": "Story",
						"subtask": false,
						"avatarId": 10318
					},
					"components": [],
					"timespent": null,
					"timeoriginalestimate": null,
					"description": "Test 1",
					"project": {
						"self": "http://localhost:58080/rest/api/2/project/10001",
						"id": "10001",
						"key": "TEST",
						"name": "Test",
						"projectTypeKey": "business",
						"avatarUrls": {
							"48x48": "http://localhost:58080/secure/projectavatar?avatarId=10324",
							"24x24": "http://localhost:58080/secure/projectavatar?size=small&avatarId=10324",
							"16x16": "http://localhost:58080/secure/projectavatar?size=xsmall&avatarId=10324",
							"32x32": "http://localhost:58080/secure/projectavatar?size=medium&avatarId=10324"
						}
					},
					"fixVersions": [],
					"customfield_10110": null,
					"customfield_10111": null,
					"aggregatetimespent": null,
					"resolution": null,
					"customfield_10104": null,
					"customfield_10105": "0|i0005b:",
					"customfield_10106": "{summaryBean=com.atlassian.jira.plugin.devstatus.rest.SummaryBean@7e603e25[summary={pullrequest=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@4b527a0f[overall=PullRequestOverallBean{stateCount=0, state='OPEN', details=PullRequestOverallDetails{openCount=0, mergedCount=0, declinedCount=0}},byInstanceType={}], build=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@718bfcda[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BuildOverallBean@29d172e6[failedBuildCount=0,successfulBuildCount=0,unknownBuildCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], review=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@2d0610d3[overall=com.atlassian.jira.plugin.devstatus.summary.beans.ReviewsOverallBean@897f304[stateCount=0,state=<null>,dueDate=<null>,overDue=false,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], deployment-environment=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@26d4b3e0[overall=com.atlassian.jira.plugin.devstatus.summary.beans.DeploymentOverallBean@4b0b39e[topEnvironments=[],showProjects=false,successfulCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], repository=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@192dbb53[overall=com.atlassian.jira.plugin.devstatus.summary.beans.CommitOverallBean@2e9564a1[count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], branch=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@2d5d8920[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BranchOverallBean@2cf0e356[count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}]},errors=[],configErrors=[]], devSummaryJson={\"cachedValue\":{\"errors\":[],\"configErrors\":[],\"summary\":{\"pullrequest\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"stateCount\":0,\"state\":\"OPEN\",\"details\":{\"openCount\":0,\"mergedCount\":0,\"declinedCount\":0,\"total\":0},\"open\":true},\"byInstanceType\":{}},\"build\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"failedBuildCount\":0,\"successfulBuildCount\":0,\"unknownBuildCount\":0},\"byInstanceType\":{}},\"review\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"stateCount\":0,\"state\":null,\"dueDate\":null,\"overDue\":false,\"completed\":false},\"byInstanceType\":{}},\"deployment-environment\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"topEnvironments\":[],\"showProjects\":false,\"successfulCount\":0},\"byInstanceType\":{}},\"repository\":{\"overall\":{\"count\":0,\"lastUpdated\":null},\"byInstanceType\":{}},\"branch\":{\"overall\":{\"count\":0,\"lastUpdated\":null},\"byInstanceType\":{}}}},\"isStale\":false}}",
					"customfield_10107": null,
					"customfield_10108": null,
					"aggregatetimeestimate": null,
					"customfield_10109": null,
					"resolutiondate": null,
					"workratio": -1,
					"summary": "Test 2",
					"lastViewed": "2021-04-05T00:00:28.005+0000",
					"watches": {
						"self": "http://localhost:58080/rest/api/2/issue/TEST-25/watchers",
						"watchCount": 1,
						"isWatching": true
					},
					"creator": {
						"self": "http://localhost:58080/rest/api/2/user?username=admin",
						"name": "admin",
						"key": "JIRAUSER10000",
						"emailAddress": "skumar@datto.com",
						"avatarUrls": {
							"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
							"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
							"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
							"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
						},
						"displayName": "SATISH KUMAR",
						"active": true,
						"timeZone": "America/Chicago"
					},
					"subtasks": [],
					"created": "2021-04-04T23:47:13.000+0000",
					"reporter": {
						"self": "http://localhost:58080/rest/api/2/user?username=admin",
						"name": "admin",
						"key": "JIRAUSER10000",
						"emailAddress": "skumar@datto.com",
						"avatarUrls": {
							"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
							"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
							"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
							"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
						},
						"displayName": "SATISH KUMAR",
						"active": true,
						"timeZone": "America/Chicago"
					},
					"aggregateprogress": {
						"progress": 0,
						"total": 0
					},
					"priority": {
						"self": "http://localhost:58080/rest/api/2/priority/3",
						"iconUrl": "http://localhost:58080/images/icons/priorities/medium.svg",
						"name": "Medium",
						"id": "3"
					},
					"customfield_10001": null,
					"customfield_10100": null,
					"labels": [],
					"environment": null,
					"timeestimate": null,
					"aggregatetimeoriginalestimate": null,
					"versions": [],
					"duedate": null,
					"progress": {
						"progress": 0,
						"total": 0
					},
					"issuelinks": [
						{
							"id": "10000",
							"self": "http://localhost:58080/rest/api/2/issueLink/10000",
							"type": {
								"id": "10001",
								"name": "Cloners",
								"inward": "is cloned by",
								"outward": "clones",
								"self": "http://localhost:58080/rest/api/2/issueLinkType/10001"
							},
							"outwardIssue": {
								"id": "10600",
								"key": "TEST-24",
								"self": "http://localhost:58080/rest/api/2/issue/10600",
								"fields": {
									"summary": "Test 1",
									"status": {
										"self": "http://localhost:58080/rest/api/2/status/10000",
										"description": "",
										"iconUrl": "http://localhost:58080/images/icons/status_generic.gif",
										"name": "To Do",
										"id": "10000",
										"statusCategory": {
											"self": "http://localhost:58080/rest/api/2/statuscategory/2",
											"id": 2,
											"key": "new",
											"colorName": "blue-gray",
											"name": "To Do"
										}
									},
									"priority": {
										"self": "http://localhost:58080/rest/api/2/priority/4",
										"iconUrl": "http://localhost:58080/images/icons/priorities/low.svg",
										"name": "Low",
										"id": "4"
									},
									"issuetype": {
										"self": "http://localhost:58080/rest/api/2/issuetype/10000",
										"id": "10000",
										"description": "A task that needs to be done.",
										"iconUrl": "http://localhost:58080/secure/viewavatar?size=xsmall&avatarId=10318&avatarType=issuetype",
										"name": "Task",
										"subtask": false,
										"avatarId": 10318
									}
								}
							}
						},
						{
							"id": "10001",
							"self": "http://localhost:58080/rest/api/2/issueLink/10001",
							"type": {
								"id": "10001",
								"name": "Cloners",
								"inward": "is cloned by",
								"outward": "clones",
								"self": "http://localhost:58080/rest/api/2/issueLinkType/10001"
							},
							"inwardIssue": {
								"id": "10602",
								"key": "TEST-26",
								"self": "http://localhost:58080/rest/api/2/issue/10602",
								"fields": {
									"summary": "Test 3",
									"status": {
										"self": "http://localhost:58080/rest/api/2/status/10001",
										"description": "",
										"iconUrl": "http://localhost:58080/images/icons/status_generic.gif",
										"name": "Done",
										"id": "10001",
										"statusCategory": {
											"self": "http://localhost:58080/rest/api/2/statuscategory/3",
											"id": 3,
											"key": "done",
											"colorName": "green",
											"name": "Done"
										}
									},
									"priority": {
										"self": "http://localhost:58080/rest/api/2/priority/2",
										"iconUrl": "http://localhost:58080/images/icons/priorities/high.svg",
										"name": "High",
										"id": "2"
									},
									"issuetype": {
										"self": "http://localhost:58080/rest/api/2/issuetype/10000",
										"id": "10000",
										"description": "A task that needs to be done.",
										"iconUrl": "http://localhost:58080/secure/viewavatar?size=xsmall&avatarId=10318&avatarType=issuetype",
										"name": "Task",
										"subtask": false,
										"avatarId": 10318
									}
								}
							}
						}
					],
					"votes": {
						"self": "http://localhost:58080/rest/api/2/issue/TEST-25/votes",
						"votes": 0,
						"hasVoted": false
					},
					"assignee": {
						"self": "http://localhost:58080/rest/api/2/user?username=admin",
						"name": "admin",
						"key": "JIRAUSER10000",
						"emailAddress": "skumar@datto.com",
						"avatarUrls": {
							"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
							"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
							"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
							"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
						},
						"displayName": "SATISH KUMAR",
						"active": true,
						"timeZone": "America/Chicago"
					},
					"updated": "2021-04-05T00:00:27.000+0000",
					"status": {
						"self": "http://localhost:58080/rest/api/2/status/3",
						"description": "This issue is being actively worked on at the moment by the assignee.",
						"iconUrl": "http://localhost:58080/images/icons/statuses/inprogress.png",
						"name": "In Progress",
						"id": "3",
						"statusCategory": {
							"self": "http://localhost:58080/rest/api/2/statuscategory/4",
							"id": 4,
							"key": "indeterminate",
							"colorName": "yellow",
							"name": "In Progress"
						}
					}
				},
				"changelog": {
					"startAt": 0,
					"maxResults": 4,
					"total": 4,
					"histories": [
						{
							"id": "10200",
							"author": {
								"self": "http://localhost:58080/rest/api/2/user?username=admin",
								"name": "admin",
								"key": "JIRAUSER10000",
								"emailAddress": "skumar@datto.com",
								"avatarUrls": {
									"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
									"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
									"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
									"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
								},
								"displayName": "SATISH KUMAR",
								"active": true,
								"timeZone": "America/Chicago"
							},
							"created": "2021-04-04T23:47:13.521+0000",
							"items": [
								{
									"field": "Link",
									"fieldtype": "jira",
									"from": null,
									"fromString": null,
									"to": "TEST-24",
									"toString": "This issue clones TEST-24"
								}
							]
						},
						{
							"id": "10202",
							"author": {
								"self": "http://localhost:58080/rest/api/2/user?username=admin",
								"name": "admin",
								"key": "JIRAUSER10000",
								"emailAddress": "skumar@datto.com",
								"avatarUrls": {
									"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
									"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
									"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
									"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
								},
								"displayName": "SATISH KUMAR",
								"active": true,
								"timeZone": "America/Chicago"
							},
							"created": "2021-04-04T23:47:16.637+0000",
							"items": [
								{
									"field": "assignee",
									"fieldtype": "jira",
									"from": null,
									"fromString": null,
									"to": "JIRAUSER10000",
									"toString": "SATISH KUMAR"
								},
								{
									"field": "status",
									"fieldtype": "jira",
									"from": "10000",
									"fromString": "To Do",
									"to": "3",
									"toString": "In Progress"
								}
							]
						},
						{
							"id": "10204",
							"author": {
								"self": "http://localhost:58080/rest/api/2/user?username=admin",
								"name": "admin",
								"key": "JIRAUSER10000",
								"emailAddress": "skumar@datto.com",
								"avatarUrls": {
									"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
									"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
									"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
									"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
								},
								"displayName": "SATISH KUMAR",
								"active": true,
								"timeZone": "America/Chicago"
							},
							"created": "2021-04-04T23:47:26.396+0000",
							"items": [
								{
									"field": "Link",
									"fieldtype": "jira",
									"from": null,
									"fromString": null,
									"to": "TEST-26",
									"toString": "This issue is cloned by TEST-26"
								}
							]
						},
						{
							"id": "10208",
							"author": {
								"self": "http://localhost:58080/rest/api/2/user?username=admin",
								"name": "admin",
								"key": "JIRAUSER10000",
								"emailAddress": "skumar@datto.com",
								"avatarUrls": {
									"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
									"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
									"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
									"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
								},
								"displayName": "SATISH KUMAR",
								"active": true,
								"timeZone": "America/Chicago"
							},
							"created": "2021-04-05T00:00:27.970+0000",
							"items": [
								{
									"field": "priority",
									"fieldtype": "jira",
									"from": "4",
									"fromString": "Low",
									"to": "3",
									"toString": "Medium"
								}
							]
						}
					]
				}
			},
			{
				"expand": "operations,versionedRepresentations,editmeta,changelog,renderedFields",
				"id": "10600",
				"self": "http://localhost:58080/rest/api/2/issue/10600",
				"key": "TEST-24",
				"fields": {
					"issuetype": {
						"self": "http://localhost:58080/rest/api/2/issuetype/10000",
						"id": "10000",
						"description": "A task that needs to be done.",
						"iconUrl": "http://localhost:58080/secure/viewavatar?size=xsmall&avatarId=10318&avatarType=issuetype",
						"name": "Task",
						"subtask": false,
						"avatarId": 10318
					},
					"components": [],
					"timespent": null,
					"timeoriginalestimate": null,
					"description": "Test 1",
					"project": {
						"self": "http://localhost:58080/rest/api/2/project/10001",
						"id": "10001",
						"key": "TEST",
						"name": "Test",
						"projectTypeKey": "business",
						"avatarUrls": {
							"48x48": "http://localhost:58080/secure/projectavatar?avatarId=10324",
							"24x24": "http://localhost:58080/secure/projectavatar?size=small&avatarId=10324",
							"16x16": "http://localhost:58080/secure/projectavatar?size=xsmall&avatarId=10324",
							"32x32": "http://localhost:58080/secure/projectavatar?size=medium&avatarId=10324"
						}
					},
					"fixVersions": [],
					"customfield_10110": null,
					"customfield_10111": null,
					"aggregatetimespent": null,
					"resolution": null,
					"customfield_10104": null,
					"customfield_10105": "0|i00053:",
					"customfield_10106": "{summaryBean=com.atlassian.jira.plugin.devstatus.rest.SummaryBean@11301747[summary={pullrequest=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@12c4c654[overall=PullRequestOverallBean{stateCount=0, state='OPEN', details=PullRequestOverallDetails{openCount=0, mergedCount=0, declinedCount=0}},byInstanceType={}], build=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@b59d6de[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BuildOverallBean@3f1b4787[failedBuildCount=0,successfulBuildCount=0,unknownBuildCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], review=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@3b9e9778[overall=com.atlassian.jira.plugin.devstatus.summary.beans.ReviewsOverallBean@32524565[stateCount=0,state=<null>,dueDate=<null>,overDue=false,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], deployment-environment=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@77789ad7[overall=com.atlassian.jira.plugin.devstatus.summary.beans.DeploymentOverallBean@123accc0[topEnvironments=[],showProjects=false,successfulCount=0,count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], repository=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@5dd0bca8[overall=com.atlassian.jira.plugin.devstatus.summary.beans.CommitOverallBean@7d89b31a[count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}], branch=com.atlassian.jira.plugin.devstatus.rest.SummaryItemBean@4e94205e[overall=com.atlassian.jira.plugin.devstatus.summary.beans.BranchOverallBean@a0cb980[count=0,lastUpdated=<null>,lastUpdatedTimestamp=<null>],byInstanceType={}]},errors=[],configErrors=[]], devSummaryJson={\"cachedValue\":{\"errors\":[],\"configErrors\":[],\"summary\":{\"pullrequest\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"stateCount\":0,\"state\":\"OPEN\",\"details\":{\"openCount\":0,\"mergedCount\":0,\"declinedCount\":0,\"total\":0},\"open\":true},\"byInstanceType\":{}},\"build\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"failedBuildCount\":0,\"successfulBuildCount\":0,\"unknownBuildCount\":0},\"byInstanceType\":{}},\"review\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"stateCount\":0,\"state\":null,\"dueDate\":null,\"overDue\":false,\"completed\":false},\"byInstanceType\":{}},\"deployment-environment\":{\"overall\":{\"count\":0,\"lastUpdated\":null,\"topEnvironments\":[],\"showProjects\":false,\"successfulCount\":0},\"byInstanceType\":{}},\"repository\":{\"overall\":{\"count\":0,\"lastUpdated\":null},\"byInstanceType\":{}},\"branch\":{\"overall\":{\"count\":0,\"lastUpdated\":null},\"byInstanceType\":{}}}},\"isStale\":false}}",
					"customfield_10107": null,
					"customfield_10108": null,
					"aggregatetimeestimate": null,
					"customfield_10109": null,
					"resolutiondate": null,
					"workratio": -1,
					"summary": "Test 1",
					"lastViewed": "2021-04-05T00:01:09.605+0000",
					"watches": {
						"self": "http://localhost:58080/rest/api/2/issue/TEST-24/watchers",
						"watchCount": 1,
						"isWatching": true
					},
					"creator": {
						"self": "http://localhost:58080/rest/api/2/user?username=admin",
						"name": "admin",
						"key": "JIRAUSER10000",
						"emailAddress": "skumar@datto.com",
						"avatarUrls": {
							"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
							"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
							"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
							"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
						},
						"displayName": "SATISH KUMAR",
						"active": true,
						"timeZone": "America/Chicago"
					},
					"subtasks": [],
					"created": "2021-04-04T23:46:56.000+0000",
					"reporter": {
						"self": "http://localhost:58080/rest/api/2/user?username=admin",
						"name": "admin",
						"key": "JIRAUSER10000",
						"emailAddress": "skumar@datto.com",
						"avatarUrls": {
							"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
							"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
							"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
							"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
						},
						"displayName": "SATISH KUMAR",
						"active": true,
						"timeZone": "America/Chicago"
					},
					"aggregateprogress": {
						"progress": 0,
						"total": 0
					},
					"priority": {
						"self": "http://localhost:58080/rest/api/2/priority/4",
						"iconUrl": "http://localhost:58080/images/icons/priorities/low.svg",
						"name": "Low",
						"id": "4"
					},
					"customfield_10001": null,
					"customfield_10100": null,
					"labels": [],
					"environment": null,
					"timeestimate": null,
					"aggregatetimeoriginalestimate": null,
					"versions": [],
					"duedate": null,
					"progress": {
						"progress": 0,
						"total": 0
					},
					"issuelinks": [
						{
							"id": "10000",
							"self": "http://localhost:58080/rest/api/2/issueLink/10000",
							"type": {
								"id": "10001",
								"name": "Cloners",
								"inward": "is cloned by",
								"outward": "clones",
								"self": "http://localhost:58080/rest/api/2/issueLinkType/10001"
							},
							"inwardIssue": {
								"id": "10601",
								"key": "TEST-25",
								"self": "http://localhost:58080/rest/api/2/issue/10601",
								"fields": {
									"summary": "Test 2",
									"status": {
										"self": "http://localhost:58080/rest/api/2/status/3",
										"description": "This issue is being actively worked on at the moment by the assignee.",
										"iconUrl": "http://localhost:58080/images/icons/statuses/inprogress.png",
										"name": "In Progress",
										"id": "3",
										"statusCategory": {
											"self": "http://localhost:58080/rest/api/2/statuscategory/4",
											"id": 4,
											"key": "indeterminate",
											"colorName": "yellow",
											"name": "In Progress"
										}
									},
									"priority": {
										"self": "http://localhost:58080/rest/api/2/priority/3",
										"iconUrl": "http://localhost:58080/images/icons/priorities/medium.svg",
										"name": "Medium",
										"id": "3"
									},
									"issuetype": {
										"self": "http://localhost:58080/rest/api/2/issuetype/10000",
										"id": "10000",
										"description": "A task that needs to be done.",
										"iconUrl": "http://localhost:58080/secure/viewavatar?size=xsmall&avatarId=10318&avatarType=issuetype",
										"name": "Task",
										"subtask": false,
										"avatarId": 10318
									}
								}
							}
						}
					],
					"votes": {
						"self": "http://localhost:58080/rest/api/2/issue/TEST-24/votes",
						"votes": 0,
						"hasVoted": false
					},
					"assignee": null,
					"updated": "2021-04-04T23:47:13.000+0000",
					"status": {
						"self": "http://localhost:58080/rest/api/2/status/10000",
						"description": "",
						"iconUrl": "http://localhost:58080/images/icons/status_generic.gif",
						"name": "To Do",
						"id": "10000",
						"statusCategory": {
							"self": "http://localhost:58080/rest/api/2/statuscategory/2",
							"id": 2,
							"key": "new",
							"colorName": "blue-gray",
							"name": "To Do"
						}
					}
				},
				"changelog": {
					"startAt": 0,
					"maxResults": 1,
					"total": 1,
					"histories": [
						{
							"id": "10201",
							"author": {
								"self": "http://localhost:58080/rest/api/2/user?username=admin",
								"name": "admin",
								"key": "JIRAUSER10000",
								"emailAddress": "skumar@datto.com",
								"avatarUrls": {
									"48x48": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=48",
									"24x24": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=24",
									"16x16": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=16",
									"32x32": "https://www.gravatar.com/avatar/023ca12cf6ce8f148e7790d7a801472f?d=mm&s=32"
								},
								"displayName": "SATISH KUMAR",
								"active": true,
								"timeZone": "America/Chicago"
							},
							"created": "2021-04-04T23:47:13.587+0000",
							"items": [
								{
									"field": "Link",
									"fieldtype": "jira",
									"from": null,
									"fromString": null,
									"to": "TEST-25",
									"toString": "This issue is cloned by TEST-25"
								}
							]
						}
					]
				}
			}
		]
	}`).Bytes()
}
