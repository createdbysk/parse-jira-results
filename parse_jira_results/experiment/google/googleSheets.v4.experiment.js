var google = require('googleapis');

jwtClient = new  google.auth.JWT(
    '297654144845-3mkk4rmp9sbpr0e7gvac3gka7u3484ct@developer.gserviceaccount.com',
    'KasasaJira.pem',
    null,
    // scope from https://developers.google.com/google-apps/spreadsheets/#authorizing_requests_with_oauth_20
    'https://www.googleapis.com/auth/spreadsheets'
);


sheets = google.sheets('v4');
sheets.spreadsheets.values.get({
    auth: jwtClient,
    spreadsheetId: '163Us5x1cLt086NEVLJJtG3zQkUTLEHgOIrfNoaXv3OQ',
    range: 'Raw Data!A3:F5'
}, function (err, response) {
    var rows,
        row;
    if (err) {
        console.log("ERROR: ", err);
        return;
    }
    console.log(JSON.stringify(response.values));
});

sheets.spreadsheets.batchUpdate({
    auth: jwtClient,
    spreadsheetId: '163Us5x1cLt086NEVLJJtG3zQkUTLEHgOIrfNoaXv3OQ',
    resource: {
        requests: [
            {
                deleteDimension: {
                    range: {
                        sheetId: 1,
                        dimension: 'ROWS',
                        startIndex: 2,
                        endIndex: 3
                    }
                }
            }
        ]
    }
});
//
// jwtClient.authorize(function (err, tokens) {
//     var auth,
//         sheets;
//     if (err) {
//         console.log("ERROR", err);
//     }
//     else {
//         console.log("TOKEN", tokens);
//         auth = {
//             type: tokens.token_type,
//             value: tokens.access_token
//         };
//         sheets = google.sheets('v4');
//         sheets.spreadsheets.values.get({
//             auth: jwtClient,
//             spreadsheetId: '163Us5x1cLt086NEVLJJtG3zQkUTLEHgOIrfNoaXv3OQ',
//             range: 'Raw Data!A3:F'
//         }, function (err, response) {
//             var rows,
//                 row;
//             if (err) {
//                 console.log("ERROR: ", err);
//                 return;
//             }
//             console.log(JSON.stringify(response.values));
//         });
//     }
// });