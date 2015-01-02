Implements the ability to parse the results from the JIRA rest API queries.

# Usage
## Extract Fields
Use bin/extractFields.js to extract fields from the raw JSON. Pass the extractors as follows 

node bin/extractFields.js -e lib/issueStatusExtractor,status -e lib/issuePriorityExtractor,priority <naw json filename>

The output is rows of JSON that contain the extracted fields for each issue.

# Library use
* Create an experiment under the experiment directory to learn how a library works.


DO NOT CHECK IN JIRA RESULTS.