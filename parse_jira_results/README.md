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

## Store results in google spreadsheet
UNDER CONSTRUCTION

### Setup a new server instance
* For the Administrator
** Generate a new P12 Key
*** Login to console.developers.google.com
*** Navigate to the project with id euphoric-coral-95415
*** Navigate to APIs & auth >> Credentials
*** On that screen, click on the "Generate new P12 key" button.
*** This will download a new .p12 file.
*** Deliver this .p12 file securely to the client machine.
* On the Client machine
** Run the following command to generate the <key>.pem file given a <key>.p12 file.
```
openssl pkcs12 -in <key>.p12 -nocerts -passin pass:notasecret -nodes -out <key>.pem
```
** Store the <key>.pem file in the server home directory.
** Add the following object as one of the keys in the googleConfiguration in the configuration file
```
clientConfiguration: {
    clientEmail: "297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com",
    clientPemFilePath: "<pem filename. The file has to be stored in the
                        server home directory."
}
```
NOTE: 297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com is the email address assigned to this application.

### Share the google sheet to store the extracted JIRA data
* Follow the directions in the Directions sheet of the [Lean Sheets Template](https://docs.google.com/spreadsheets/d/14MfkssorAK9OxNJTOJoljAmkRU-PqazlZfCoOGdYl7o/edit?usp=sharing)
* This will result in a new Google Sheet.
* Share the sheet with 297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com
* Note the spreadsheet key, which is part of the shareable link, to use with the application that uploads the spreadsheet data.
** For example, given a shareable link of the following form
```
https://docs.google.com/spreadsheets/d/14MfkssorAK9OxNJTOJoljAmkRU-PqazlZfCoOGdYl7o/edit?usp=sharing
```
the key is
```
14MfkssorAK9OxNJTOJoljAmkRU-PqazlZfCoOGdYl7o
```

Points to highlight
* Provide an empty sheet that has a header row that matches the extracted fields.
* Remember to empty the worksheet before every call to this utility.

### Example configuration file
```
define({
    // names and modules form a name value pair.
    // Each element in the names array should have a corresponding
    // module in the modules array.
    moduleConfiguration: {
        names: ['name', 'type', 'priority', 'startDate', 'endDate'],
        modules: [
            'lib/issueNameExtractor.js',
            'lib/issueTypeExtractor.js',
            'lib/issuePriorityExtractor.js',
            'lib/issueStartDateExtractor.js',
            'lib/issueEndDateExtractor.js'
        ]
    },
    jiraConfiguration: {
        baseUrl: "<JIRA url>",
        username: '<username>',
        password: '<password>',
        strictSSL: <true or false>
    },
    googleConfiguration: {
        clientConfiguration: {
            clientEmail: "297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com",
            clientPemFilePath: "<pem filename. The file has to be stored in the
                                server home directory."
        },
        numberOfRowsToAddInParallel: <The number of rows that the process will upload
                                      to the spreadsheet at a time, which depends on
                                      the internet bandwidth. Start with 20.
                                      You can increase this number until you see the following error.
                                      ERROR addRow { [Error: connect ETIMEDOUT] code: 'ETIMEDOUT', errno: 'ETIMEDOUT',
                                      syscall: 'connect' }
    }
});
```

# Library use
* Create an experiment under the experiment directory to learn how a library works.

# References
## Links
[Accessing Google Spreadsheets from NodeJs]( http://www.nczonline.net/blog/2014/03/04/accessing-google-spreadsheets-from-node-js/)
[Google Oauth2 playground](https://developers.google.com/oauthplayground/)
[Google Javascript Client API](https://developers.google.com/api-client-library/javascript/start/start-js)
[Google NodeJs Client API](https://github.com/google/google-api-nodejs-client)



DO NOT CHECK IN JIRA RESULTS.
