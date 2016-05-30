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
            displayResults,
            configurationFileName,
            searchQuery,
            spreadsheetKey,
            worksheetIndex,
            jiraRest;

        program
            .version('0.0.1')
            .usage('configuration_file_name search_query spreadsheet_key worksheet_index')
            .parse(process.argv);
        configurationFileName = program.args[0];
        searchQuery = program.args[1];
        spreadsheetKey = program.args[2];
        worksheetIndex = parseInt(program.args[3]);
        requirejs([configurationFileName], function (configuration) {
            jiraRest = JiraRest(configuration.jiraConfiguration);
            console.log("Pulling data for query", searchQuery)
            jiraRest.search(searchQuery, function (err, result) {
                if (err) {
                    console.log("ERROR WITH JIRA QUERY", err);
                } else {
                    console.log("JIRA QUERY SUCCEEDED");
                    transformLoader.loadModules(configuration.moduleConfiguration, function (err, transforms) {
                        var fields;
                        processIssue =
                            function (issue) {
                                var resultsWithExtractedFields;
                                transformer(issue, transforms, function (err, extractedFields) {
                                    resultsWithExtractedFields = extractedFields;
                                });
                                return resultsWithExtractedFields;
                            };
                        processResults = function (results) {
                            var issuesWithExtractFields;
                            console.log("Found", result.issues.length, "issues.")
                            issuesWithExtractFields =
                                linq.from(results.issues)
                                    .select(processIssue);
                            return issuesWithExtractFields;
                        };
                        displayResults = function (err, allIssuesWithExtractFields) {
                            googleSpreadsheetClientFactory.createClient(
                                configuration.googleConfiguration.clientConfiguration,
                                function (err, client) {
                                    var spreadsheet;
                                    if (err) {
                                        console.log("ERROR createClient", err);
                                    }
                                    else {
                                        console.log("Loading results into google spreadsheet", spreadsheetKey);
                                        client.getSpreadsheet(spreadsheetKey,
                                            function (err, spreadsheet) {
                                                async.eachLimit(allIssuesWithExtractFields.toArray(),
                                                                configuration.googleConfiguration.numberOfRowsToAddInParallel,
                                                    function (issueWithExtractFields, continuation) {
                                                        spreadsheet.addRow(worksheetIndex, issueWithExtractFields, function (err) {
                                                            if (err) {
                                                                console.log("Add Row Error: ", err, issueWithExtractFields)
                                                            }
                                                            continuation(err);
                                                        });
                                                    },
                                                    function (err) {
                                                        console.log("DONE");
                                                        if (err) {
                                                            console.log("ERROR addRow", err);
                                                        }
                                                    }
                                                );
                                            }
                                        );
                                    }
                                 }
                            );
                        };
                        fields = processResults(result);
                        displayResults(null, fields);
                    });
                }
            });

        });
    }
);
