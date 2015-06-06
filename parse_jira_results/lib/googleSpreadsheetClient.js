/**
 * Creates the google spreadsheet client.
 */
define(['googleapis',
        'google-spreadsheet'],
    function (googleapis, GoogleSpreadsheet) {
        var googleSpreadsheetClient,
            createClient;

        createClient = function (configuration, callback) {
            var jwtClient,
                client,
                makeGetSpreadsheet;

            jwtClient = new googleapis.auth.JWT(
                configuration.clientEmail,
                configuration.clientPemFilePath,
                null,
                // scope from https://developers.google.com/google-apps/spreadsheets/#authorizing_requests_with_oauth_20
                'https://spreadsheets.google.com/feeds'
            );

            console.log("JWT", jwtClient);

            makeGetSpreadsheet = function (auth) {
                return function (spreadsheetName, callback) {
                    var spreadsheet;
                    spreadsheet = new GoogleSpreadsheet(spreadsheetName, auth);
                    callback(null, spreadsheet);
                }
            }

            jwtClient.authorize(function (err, tokens) {
                var auth;
                if (err) {
                    console.log("AUTHORIZE ERR", err);
                    callback(err);
                }
                else {
                    var client,
                        getSpreadsheet;
                    auth = {
                        type: tokens.token_type,
                        value: tokens.access_token
                    };
                    getSpreadsheet = makeGetSpreadsheet(auth);
                    client = {
                        getSpreadsheet : getSpreadsheet
                    };
                    callback(null, client);
                }
            });
        };

        googleSpreadsheetClient = {
            createClient : createClient
        };
        return googleSpreadsheetClient;
    }
);
