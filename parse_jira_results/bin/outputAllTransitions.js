// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('../configuredRequirejs.js');

requirejs(['fs', 
          'linq', 
          'moment',
          'lib/issueStatusExtractor', 
          'lib/statusFilter',
          'lib/leadTimeCalculator'], 
    function (fs, linq, moment, issueStatusExtractor, statusFilter, leadTimeCalculator) {
        'use strict';
        fs.readFile('test_input/SW-2000.txt', 'utf-8', function (err, resultsJSON) {
            var allResults,
                leadTimes,
                allTransitions;
            allResults = JSON.parse(resultsJSON);
            console.error("Number of JIRA queries", allResults.length);
            allTransitions = linq.from(allResults)
                .selectMany(function (results) {
                    var allTransitionsForThisSet;
                    issueStatusExtractor(results, function (err, issuesWithStatuses) {
                        console.error(JSON.stringify(issuesWithStatuses.toArray(), undefined, 4));
                        allTransitionsForThisSet = linq.from(issuesWithStatuses)
                            .select(function (issue) {
                                var allTransitionsForThisIssue;
                                var startDate,
                                    endDate;
                                allTransitionsForThisIssue = linq.from(issue.statuses)
                                    .aggregate("Triage, " + issue.createdDate,
                                               function (csv, status) {
                                                    return csv + ", " + status.to + ", " + moment(status.date).format("YYYY-MM-DD HH:MM:SS");
                                                });
                                return issue.key + ", " + allTransitionsForThisIssue;
                            });
                    });
                    return allTransitionsForThisSet;
                });
            console.log(JSON.stringify(allTransitions.toArray(), undefined, 4));
        });
    }
);

