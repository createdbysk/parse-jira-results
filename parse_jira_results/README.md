Implements the ability to parse the results from the JIRA rest API queries.

# Usage
## Extract Fields
Use bin/extractFields.js to extract fields from the raw JSON. Pass the extractors as follows

```
node bin/extractFields.js -e lib/issueStatusExtractor,status -e lib/issuePriorityExtractor,priority <raw json filename>
```

The output is rows of JSON that contain the extracted fields for each issue.

## Calculate Metrics
Use bin/calculateMetrics.js to calculate the metrics based on the extracted fields.

In the example below, bin/extractFields.sh is a convenience script that invokes bin/extractFields.js with a chosen set of fields.

```
bin/extractFields.sh <raw json filename> | node bin/calculateMetrics.js -o <comma separated list of statuses> <name of field to report>
```

The output is a csv.

For example,

```
bin/extractFields.sh <raw json filename> | node bin/calculateMetrics.js -o "Triage,Open,In Design,In Progress,In Review,Ready for Testing" duration
```

# Library use
* Create an experiment under the experiment directory to learn how a library works.

# References
## Links
[Accessing Google Spreadsheets from NodeJs]( http://www.nczonline.net/blog/2014/03/04/accessing-google-spreadsheets-from-node-js/)
[Google Oauth2 playground](https://developers.google.com/oauthplayground/)
[Google Javascript Client API](https://developers.google.com/api-client-library/javascript/start/start-js)
[Google NodeJs Client API](https://github.com/google/google-api-nodejs-client)

## How to generate key file
// Follow directions on this page at -https://github.com/google/google-api-nodejs-client
// under Using JWT (Service Tokens)
// Use the following command to generate the .pem file from the
// .p12 file downloaded from the Google Developers Console
// openssl pkcs12 -in key.p12 -nocerts -passin pass:notasecret -nodes -out key.pem
// put the .pem file in the same directory as this script.
DO NOT CHECK IN JIRA RESULTS.
