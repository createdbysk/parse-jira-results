// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('library/configuredRequirejs');

requirejs(['commander',
          'bin/readFileAndIterate',
          'linq',
          'async',
          'path',
          'library/transformLoader',
          'library/transformer',
          'library/jiraRest',
          'lib/googleSpreadsheetClient'
          ],
    function (
        program,
        readFileAndIterate,
        linq,
        async,
        path,
        transformLoader,
        transformer,
        JiraRest,
        googleSpreadsheetClientFactory)
    {
        'use strict';
        var extractors,
            processIssue,
            processResults,
            uploadResults,
            configurationFileName,
            searchQuery,
            spreadsheetId,
            sheetName,
            jiraRest,
            firstRowIndex;

        program
            .version('0.0.1')
            .usage('configuration_file_name search_query spreadsheet_key sheetName')
            .parse(process.argv);
        configurationFileName = program.args[0];
        searchQuery = program.args[1];
        spreadsheetId = program.args[2];
        sheetName = program.args[3];
        firstRowIndex = 2;
        requirejs([configurationFileName], function (configuration) {
            jiraRest = JiraRest(configuration.jiraConfiguration);
            console.log("Pulling data for query", searchQuery)
            jiraRest.search(searchQuery, function (err, result) {
                if (err) {
                    console.log("ERROR WITH JIRA QUERY", err);
                    process.exit(1);
                } else {
                    console.log("JIRA QUERY SUCCEEDED");
                    transformLoader.loadModules(configuration.moduleConfiguration, function (err, transforms) {
                        var fields;
                        processIssue =
                            function (issue) {
                                var resultsWithExtractedFields;
                                var issueAndModuleConfiguration;
                                issueAndModuleConfiguration = {
                                  issue: issue,
                                  moduleConfiguration: configuration.moduleConfiguration
                                }
                                transformer(issueAndModuleConfiguration, transforms, function (err, extractedFields) {
                                    resultsWithExtractedFields = extractedFields;
                                });
                                return resultsWithExtractedFields;
                            }
                        processResults = function (results) {
                            var issuesWithExtractFields;
                            console.log("Found", result.issues.length, "issues.")
                            issuesWithExtractFields =
                                linq.from(results.issues)
                                    .select(processIssue);
                            return issuesWithExtractFields;
                        }
                        uploadResults = function (allIssuesWithExtractFields, uploadResultsCallback) {
                            var getSheetProperties,
                                storeResults,
                                insertRows,
                                clearSheet;
                            getSheetProperties = function (v4Client, callback) {
                                console.log("In getSheetProperties");
                                v4Client.getSheetProperties(spreadsheetId, sheetName, callback);
                            }
                            clearSheet = function (v4Client, properties, callback) {
                                console.log("In clearSheet");
                                // If there are more than firstRowIndex rows, then delete all the rows after the firstRowIndex.
                                if (properties.rowCount > firstRowIndex) {
                                    v4Client.deleteRowsRange(spreadsheetId, properties.sheetId, firstRowIndex, properties.rowCount, callback);
                                } else {
                                    callback(null);
                                }
                            }
                            storeResults = function(properties, arrayOfIssues, callback) {
                                console.log("In storeResults", JSON.stringify(arrayOfIssues, null, 4));
                                // Use a Legacy client because it uses the field names to determine the columns to store
                                // the data into.
                                googleSpreadsheetClientFactory.createLegacyClient(
                                    configuration.googleConfiguration.clientConfiguration,
                                    function (err, client) {
                                        var spreadsheet,
                                            worksheetIndex;
                                        // The index used by the legacy client is 1 based while
                                        // the properties returned by the V4 Client are zero based.
                                        worksheetIndex = properties.sheetIndex + 1;
                                        if (err) {
                                            console.log("ERROR createClient", err);
                                        }
                                        else {
                                            console.log("Loading results into google spreadsheet", spreadsheetId, " into sheet with index ", worksheetIndex);
                                            client.getSpreadsheet(spreadsheetId,
                                                function (err, spreadsheet) {
                                                    async.eachLimit(arrayOfIssues,
                                                                    configuration.googleConfiguration.numberOfRowsToAddInParallel,
                                                        function (issueWithExtractFields, continuation) {
                                                            spreadsheet.addRow(worksheetIndex, issueWithExtractFields, function (err) {
                                                                if (err) {
                                                                    console.log("Add Row Error: ", err, issueWithExtractFields)
                                                                }
                                                                continuation(err);
                                                            });
                                                        },
                                                        callback
                                                    );
                                                }
                                            );
                                        }
                                     }
                                );
                            }
                            googleSpreadsheetClientFactory.createV4Client(
                                configuration.googleConfiguration.clientConfiguration,
                                function (err, v4Client) {
                                    if (err) {
                                        uploadResultsCallback(err);
                                    }
                                    getSheetProperties(v4Client, function (err, properties) {
                                        if (err) {
                                            uploadResultsCallback(err);
                                        }
                                        clearSheet(v4Client, properties, function (err) {
                                          var arrayOfIssues;
                                          if (err) {
                                              uploadResultsCallback(err);
                                          }
                                          arrayOfIssues = allIssuesWithExtractFields.toArray();
                                          storeResults(properties, arrayOfIssues, uploadResultsCallback);
                                      });
                                    });
                                }
                            );
                        }
                        fields = processResults(result);
                        uploadResults(fields,
                            function (err) {
                                if (err) {
                                    console.log("uploadResults: ERROR: ", err);
                                    process.exit(1);
                                }
                            }
                        );
                    });
                }
            });

        });
    }
);
