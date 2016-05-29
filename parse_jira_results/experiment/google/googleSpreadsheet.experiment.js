var GoogleSpreadsheet,
    mySheet,
    googleapis,
    jwtClient,
    auth;

GoogleSpreadsheet = require('google-spreadsheet');
googleapis = require('googleapis');
// Follow directions on this page at -https://github.com/google/google-api-nodejs-client
// under Using JWT (Service Tokens)
// Use the following command to generate the .pem file from the
// .p12 file downloaded from the Google Developers Console
// openssl pkcs12 -in key.p12 -nocerts -passin pass:notasecret -nodes -out key.pem
// put the .pem file in the same directory as this script.
jwtClient = new  googleapis.auth.JWT(
    '297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com',
    'KasasaJira.pem',
    null,
    // scope from https://developers.google.com/google-apps/spreadsheets/#authorizing_requests_with_oauth_20
    'https://spreadsheets.google.com/feeds'
);

console.log("JWTclient", jwtClient);
jwtClient.authorize(function (err, tokens) {
    var auth;
    if (err) {
        console.log("ERROR", err);
    }
    else {
        console.log("TOKEN", tokens);
        auth = {
            type: tokens.token_type,
            value: tokens.access_token
        };
        // Access the spreadsheet at
        // https://docs.google.com/spreadsheets/d/1krSj0KrswrtC3hCgohXfyYaItkfvVOf_mMQICSAA2G4/edit?usp=sharing
        // Notice the key below matches the key in the url above.
        mySheet = new GoogleSpreadsheet('1krSj0KrswrtC3hCgohXfyYaItkfvVOf_mMQICSAA2G4', auth);
        // Access the rows of the first worksheet
        mySheet.getRows(1, function (err, rowData) {
            console.log("Error", err);
            console.log("RowData", JSON.stringify(rowData));
        });
        // Add a row to second worksheet.
        mySheet.addRow(1, {This:"25", is:"345", a:"3059", test:'ueteu'}, function (err) {
            console.log("ADD ERR", err);
        });
        // Get info about the worksheets.
        mySheet.getInfo( function(err, sheet_info) {
             if (err) {
                 console.log("ERROR", err);
             }
             else {
                 console.log( 'pulled in '+JSON.stringify(sheet_info));
             }
        });
    }
});
