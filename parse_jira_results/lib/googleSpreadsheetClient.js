/**
 * Creates the google spreadsheet client.
 */
define(['googleapis',
        'google-spreadsheet'],
    function (googleapis, GoogleSpreadsheet) {
        var googleSpreadsheetClient,
            createClient;

        createClient = function (configuration, callback) {
            var jwtClient;
            jwtClient = new googleapis.auth.JWT(
                configuration.clientEmail,
                configuration.clientPemFilePath,
                null,
                // scope from https://developers.google.com/google-apps/spreadsheets/#authorizing_requests_with_oauth_20
                'https://spreadsheets.google.com/feeds'
            );
            callback(null, jwtClient);
        };

        googleSpreadsheetClient = {
            createClient : createClient
        };
        return googleSpreadsheetClient;
    }
);
