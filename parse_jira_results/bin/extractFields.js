// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('library/configuredRequirejs');

requirejs(['commander',
          'bin/readFileAndIterate',
          'linq',
          'path',
          'lib/issueStatusExtractor',
          'lib/issuePriorityExtractor',
          'library/transformLoader',
          'library/transformer'
          ],
    function (
        program,
        readFileAndIterate,
        linq,
        path,
        issueStatusExtractor,
        issuePriorityExtractor,
        transformLoader,
        transformer)
    {
        'use strict';
        var formatDate,
            extractors,
            processIssue,
            processResults,
            displayResults;

        transformLoader
            .configureCommander(program)
            .transform('-e --extractor', 'requirejs modules that extract values')
            .version('0.0.1')
            .parse(process.argv);

        transformLoader.loadModules(program.extractor, function (err, transforms) {
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
                issuesWithExtractFields =
                    linq.from(results.issues)
                        .select(processIssue);
                return issuesWithExtractFields;
            };
            displayResults = function (err, allIssuesWithExtractFields) {
                linq.from(allIssuesWithExtractFields)
                    .forEach(function (issueWithExtractFields) {
                        console.log(JSON.stringify(issueWithExtractFields));
                    });
            };
            readFileAndIterate(program.args[0], processResults, displayResults);
        });
    }
);
