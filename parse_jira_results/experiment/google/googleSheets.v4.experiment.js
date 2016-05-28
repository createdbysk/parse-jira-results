var google = require('googleapis');

jwtClient = new  google.auth.JWT(
    '297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com',
    'KasasaJira.pem',
    null,
    // scope from https://developers.google.com/google-apps/spreadsheets/#authorizing_requests_with_oauth_20
    'https://www.googleapis.com/auth/spreadsheets'
);


sheets = google.sheets('v4');
// Retrieve the values
sheets.spreadsheets.values.get({
    auth: jwtClient,
    spreadsheetId: '1Y6SCKUNqXg75Rd9Dt0jHbUq7ueODddE-Q32JFwlQR8E',
    range: 'Sheet1!A1:C2'
}, function (err, response) {
    var rows,
        row;
    if (err) {
        console.log("ERROR: ", err);
        return;
    }
    console.log(JSON.stringify(response.values));
    // Add values to the sheet.
    sheets.spreadsheets.values.update({
        auth: jwtClient,
        spreadsheetId: '1Y6SCKUNqXg75Rd9Dt0jHbUq7ueODddE-Q32JFwlQR8E',
        range: 'Sheet1!A1:C',
        valueInputOption: 'RAW',
        resource: {
            values: [
                [
                    'A1', 'B1', 'C1'
                ],
                [
                    'A2', 'B2', 'C2'
                ]
            ]
        }
    }, function (err, response) {
        var rows,
            row;
        if (err) {
            console.log("ERROR: ", err);
            return;
        }
        console.log("ADDED ROWS");
        console.log(JSON.stringify(response));
        // Retrieve the values
        sheets.spreadsheets.values.get({
            auth: jwtClient,
            spreadsheetId: '1Y6SCKUNqXg75Rd9Dt0jHbUq7ueODddE-Q32JFwlQR8E',
            range: 'Sheet1!A1:C2'
        }, function (err, response) {
            var rows,
                row;
            if (err) {
                console.log("ERROR: ", err);
                return;
            }
            console.log(JSON.stringify(response.values));
            // Delete the values
            sheets.spreadsheets.batchUpdate({
                auth: jwtClient,
                spreadsheetId: '1Y6SCKUNqXg75Rd9Dt0jHbUq7ueODddE-Q32JFwlQR8E',
                resource: {
                    requests: [
                        {
                            deleteDimension: {
                                range: {
                                    sheetId: 0,
                                    dimension: 'ROWS',
                                    startIndex: 0,
                                    endIndex: 2
                                }
                            }
                        }
                    ]
                }
            }, function (err, response) {
                if (err) {
                    console.log("ERROR: ", err);
                    return;
                }
                console.log("DELETED ROWS");
                console.log(JSON.stringify(response));
            });
        });
    });
});
