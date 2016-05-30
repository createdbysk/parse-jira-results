/**
 * Creates the google spreadsheet client.
 */
define(['googleapis',
        'google-spreadsheet'],
    function (google, GoogleSpreadsheet) {
        var googleSpreadsheetClient,
            createClient;

        createLegacyClient = function (configuration, callback) {
            var jwtClient,
                client,
                makeGetSpreadsheet;

            jwtClient = new google.auth.JWT(
                configuration.clientEmail,
                configuration.clientPemFilePath,
                null,
                // scope from https://developers.google.com/google-apps/spreadsheets/#authorizing_requests_with_oauth_20
                'https://spreadsheets.google.com/feeds'
            );

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

        createV4Client = function (configuration, callback) {
            var jwtClient;

            jwtClient = new google.auth.JWT(
                configuration.clientEmail,
                configuration.clientPemFilePath,
                null,
                'https://www.googleapis.com/auth/spreadsheets'
            );


            jwtClient.authorize(function (err) {
                var v4Client;
                if (err) {
                    console.log("AUTHORIZE ERR", err);
                    callback(err);
                }
                else {
                    var sheets,
                        getSheetProperties,
                        deleteRowsRange;

                    sheets = google.sheets('v4');
                    /**
                     * Get the properties of the sheet.
                     *
                     */
                    getSheetProperties = function (spreadsheetId, sheetName, callback) {
                        // First get the information about the spreadsheet.
                        sheets.spreadsheets.get({
                            auth: jwtClient,
                            spreadsheetId: spreadsheetId,
                            ranges: sheetName
                        }, function (err, response) {
                            var properties,
                                propertiesInResponse;
                            if (err) {
                                console.log("getSheetProperties: sheets.spreadsheets.get " + spreadsheetId + " " + sheetName + " failed with ERROR: ", err);
                                callback(err);
                            }
                            propertiesInResponse = response.sheets[0].properties;
                            properties = {
                                sheetId: propertiesInResponse.sheetId,
                                sheetIndex: propertiesInResponse.index,
                                rowCount: propertiesInResponse.gridProperties.rowCount
                            }
                            callback(null, properties);
                        });
                    }

                    /**
                     * Delete the given range of rows from the given sheet.
                     */
                    deleteRowsRange = function (spreadsheetId, sheetId, startIndex, endIndex, callback) {
                        // Delete the values
                        sheets.spreadsheets.batchUpdate({
                            auth: jwtClient,
                            spreadsheetId: spreadsheetId,
                            resource: {
                                requests: [
                                    {
                                        deleteDimension: {
                                            range: {
                                                sheetId: sheetId,
                                                dimension: 'ROWS',
                                                startIndex: startIndex,
                                                endIndex: endIndex
                                            }
                                        }
                                    }
                                ]
                            }
                        }, function (err) {
                            if (err) {
                                console.log("deleteRowsRange(" + spreadsheetId + " " + sheetId + " " + startIndex + " " + endIndex + "): sheets.spreadsheets.batchUpdate failed with ERROR: ", err);
                                callback(err);
                            }
                            callback(null);
                        });
                    }

                    insertRowsRange = function (spreadsheetId, sheetId, startIndex, endIndex, callback) {
                        // Delete the values
                        sheets.spreadsheets.batchUpdate({
                            auth: jwtClient,
                            spreadsheetId: spreadsheetId,
                            resource: {
                                requests: [
                                    {
                                        insertDimension: {
                                            range: {
                                                sheetId: sheetId,
                                                dimension: 'ROWS',
                                                startIndex: startIndex,
                                                endIndex: endIndex
                                            }
                                        }
                                    }
                                ]
                            }
                        }, function (err) {
                            if (err) {
                                console.log("insertRowsRange: sheets.spreadsheets.batchUpdate failed with ERROR: ", err);
                                callback(err);
                            }
                            callback(null);
                        });
                    }
                    v4Client = {
                        getSheetProperties: getSheetProperties,
                        deleteRowsRange: deleteRowsRange,
                        insertRowsRange: insertRowsRange
                    }

                    callback(null, v4Client);
                }
            });
        }

        googleSpreadsheetClient = {
            createLegacyClient : createLegacyClient,
            createV4Client: createV4Client
        };
        return googleSpreadsheetClient;
    }
);
