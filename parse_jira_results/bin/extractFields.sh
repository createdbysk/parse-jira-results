#!/bin/sh
node bin/extractFields.js -e lib/issueNameExtractor.js,name -e lib/issueTypeExtractor,type -e lib/issueCreatedDateExtractor.js,createdDate -e lib/issueStartDateExtractor,startDate -e lib/issueEndDateExtractor,endDate -e lib/issuePriorityExtractor,priority -e lib/issueStoryPointsExtractor,storyPoints -e lib/issueStatusExtractor.js,statuses $1 
