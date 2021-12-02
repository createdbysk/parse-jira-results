# golang

## Usage
### Run query and use template to format output

        JIRA_URL=jira_url JIRA_USERNAME=user JIRA_PASSWORD=password ./read-from-jira/jira -maxResults 50 ./renderer_templates/checkStatusChange.tpl 'project = "TEST"' | tee output.txt

### Upload to Google Sheet

        cat output.txt | CREDENTIALS_FILEPATH=$(echo /path/to/credentials/file) SPREADSHEET_ID={:spreadsheet_id} DELIMITER="|" ./load-into-googlesheet/sheetsLoader 'TestSheet!A2'

* The command-parameter is the Destination Cell Reference. It MUST include the single quotes ('') at least on the bash command-line because ! is the symbol in bash to access history.
* Use echo in $(echo /path/to/credentials/file) to expand to the full path.
* Replace {:spreadsheet_id} with the actual spreadsheet id, which you can find in the google sheets URL of the form

        https://docs.google.com/spreadsheets/d/{:spreadsheet_id}/edit#gid=12345678


### Test with Local JIRA
#### First time

        make run-jira
#### Subsequent times

        make start-jira

