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
The application is an ETL process that
* extracts the data specified by a jql (JIRA query language) query from the configured JIRA instance
* transforms the json data returned by the extract process to obtain the required fields to store in the input to lean sheets
* stores the values for the fields obtained from the json data in the specified destination

The application uses
* a configuration file that
** specifies the source JIRA instance
** specifies the transforms that extract the required fields
** specifies the parameters that configure the application as a google client application
** Example configuration file
```
define({
    // The JIRA configuration for the source of the extract
    jiraConfiguration: {
        baseUrl: "<JIRA url>",
        username: '<username>',
        password: '<password>',
        strictSSL: <true or false>
    },
    // The modules that transform the extracted data to obtain the fields to
    // load into the destination.
    //
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
    // The google configuration that allows the application to load the values to
    // the destination google sheet.
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
* command-line parameters that specify the parameters of each run. These are
** configuration_file_name  - The full path to the configuration file
** search_query             - The JIRA query to extract the data
** spreadsheet_key          - The key for the destination google sheet
** worksheet_index          - The index of the worksheet within the google sheet

Example command-line use
```
node bin/storeJiraResultsInGoogleSheet 'path/to/configuration_file_name.js' 'project%3DSW' '14MfkssorAK9OxNJTOJoljAmkRU-PqazlZfCoOGdYl7o' 2
```

The following sections detail how to setup the configuration file and
obtain the command-line parameters.

### Setup a new client
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
** Add the following object as the value of the clientConfiguration property in the googleConfiguration in the configuration file
```
{
    clientEmail: "297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com",
    clientPemFilePath: "<pem filename. The file has to be stored in the
                        server home directory."
}
```
NOTE: 297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com is the email address assigned to this application.

For example,
```
{
    clientEmail: "297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com",
    clientPemFilePath: "my_pem_file.pem"
}
```


### Share the google sheet to store the extracted JIRA data
* Follow the directions in the Directions sheet of the [Lean Sheets Template](https://docs.google.com/spreadsheets/d/14MfkssorAK9OxNJTOJoljAmkRU-PqazlZfCoOGdYl7o/edit?usp=sharing)
* This will result in a new Google Sheet.
* Share the sheet with 297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com
* Note the spreadsheet key, which is part of the shareable link, as the spreadsheet_key command-line parameter to the application.
** For example, given a shareable link of the following form
```
https://docs.google.com/spreadsheets/d/14MfkssorAK9OxNJTOJoljAmkRU-PqazlZfCoOGdYl7o/edit?usp=sharing
```
the spreadsheet_key is
```
14MfkssorAK9OxNJTOJoljAmkRU-PqazlZfCoOGdYl7o
```

### Configure JIRA use
Supply the following object as the value for the jiraConfiguration property of the configuration.
```
{
    baseUrl: "<JIRA url>",
    username: '<username>',
    password: '<password>',
    strictSSL: <true or false>
}
```

* Set strictSSL to true if client machine is setup correctly to access JIRA over https.
* Set strictSSL to false if the client machine generates certificate errors with the JIRA.
* Access JIRA through the supplied URL. If the browser generates prompts that indicate that the JIRA instance is unsafe to access, then the client is not configured correctly to access the JIRA instance over SSL.

For example,
```
{
    baseUrl: 'https://jira.dev.my_company_name.com',
    username: 'my_username',
    password: 'my_password',
    strictSSL: true
}
```

### Configure the modules that extract the fields to load
This section describes how to setup the moduleConfiguration property of the configuration file. The format of the data is as follows.
```
{
    names: [<array with names that correspond to names of the columns in the sheet specified by worksheet_index>],
    modules: [<array of modules that extract the corresponding field in the names array>]
}
```
For example,
```
{
    names: ['name', 'type', 'priority', 'startDate', 'endDate'],
    modules: [
        'lib/issueNameExtractor.js',
        'lib/issueTypeExtractor.js',
        'lib/issuePriorityExtractor.js',
        'lib/issueStartDateExtractor.js',
        'lib/issueEndDateExtractor.js'
    ]
}
```

Points to highlight
* Provide an empty sheet that has a header row that matches the extracted fields.
* Remember to empty the worksheet before every call to this utility.

# Library use
* Create an experiment under the experiment directory to learn how a library works.

# References
## Links
[Accessing Google Spreadsheets from NodeJs]( http://www.nczonline.net/blog/2014/03/04/accessing-google-spreadsheets-from-node-js/)
[Google Oauth2 playground](https://developers.google.com/oauthplayground/)
[Google Javascript Client API](https://developers.google.com/api-client-library/javascript/start/start-js)
[Google NodeJs Client API](https://github.com/google/google-api-nodejs-client)



DO NOT CHECK IN JIRA RESULTS.
