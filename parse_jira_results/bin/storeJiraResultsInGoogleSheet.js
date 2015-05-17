// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('library/configuredRequirejs');

requirejs(['commander',
          'bin/readFileAndIterate',
          'linq',
          'path',
          'library/transformLoader',
          'library/transformer',
          'library/jiraRest'
          ],
    function (
        program,
        readFileAndIterate,
        linq,
        path,
        issueStatusExtractor,
        issuePriorityExtractor,
        transformLoader,
        transformer,
        jiraRest)
    {
        'use strict';
        var extractors,
            processIssue,
            processResults,
            displayResults;

        program
            .version('0.0.1')
            .usage('configuration_file_name')
            .parse(process.argv);

        requirejs(program.args[0], function (configuration) {
            console.log("CONFIGURATION", configuration);
        });
    //     transformLoader.loadModules(program.extractor, function (err, transforms) {
    //         processIssue =
    //             function (issue) {
    //                 var resultsWithExtractedFields;
    //                 transformer(issue, transforms, function (err, extractedFields) {
    //                     resultsWithExtractedFields = extractedFields;
    //                 });
    //                 return resultsWithExtractedFields;
    //             };
    //         processResults = function (results) {
    //             var issuesWithExtractFields;
    //             issuesWithExtractFields =
    //                 linq.from(results.issues)
    //                     .select(processIssue);
    //             return issuesWithExtractFields;
    //         };
    //         displayResults = function (err, allIssuesWithExtractFields) {
    //             linq.from(allIssuesWithExtractFields)
    //                 .forEach(function (issueWithExtractFields) {
    //                     console.log(JSON.stringify(issueWithExtractFields));
    //                 });
    //         };
    //         readFileAndIterate(program.args[0], processResults, displayResults);
    //     });
    }
);
