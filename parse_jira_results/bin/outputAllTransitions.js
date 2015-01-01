// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('../configuredRequirejs.js');

requirejs(['commander',
          'fs', 
          'linq', 
          'moment',
          'lib/issueStatusExtractor', 
          'lib/issueCreatedDateExtractor', 
          'lib/statusFilter',
          'lib/leadTimeCalculator'], 
    function (program, fs, linq, moment, issueStatusExtractor, issueCreatedDateExtractor, statusFilter, leadTimeCalculator) {
        'use strict';
        program
            .version('0.0.1')
            .parse(process.argv);
        fs.readFile(program.args[0], 'utf-8', function (err, resultsJSON) {
            var allResults,
                leadTimes,
                allTransitions,
                formatDate;
            formatDate = function (rawDate) {
                if (!rawDate) {
                    console.error("Date to format is undefined");
                }
                return moment(rawDate).format("YYYY-MM-DD hh:mm:ss");
            }
            allResults = JSON.parse(resultsJSON);
            if (!allResults.length) {
                allResults = [allResults];
            }
            console.error("Number of JIRA queries", allResults.length);
            allTransitions = linq.from(allResults)
                .selectMany(function (results) {
                    var allTransitionsForThisSet;
                    allTransitionsForThisSet = 
                        linq.from(results.issues)
                            .select(function (issue) {
                                var allTransitionsForThisIssue;
                                issueStatusExtractor(issue, function (err, statuses) {
                                    console.error("status count : ", statuses.count());
                                    allTransitionsForThisIssue = linq.from(statuses)
                                        .aggregate(issue.key,
                                                   function (csv, status) {
                                                        return csv + ", " + status.to + ", " + formatDate(status.date);
                                                    });
                                    console.error("transitions : ", allTransitionsForThisIssue);
                                });
                                return allTransitionsForThisIssue;
                            });
                    return allTransitionsForThisSet;
                });
            linq.from(allTransitions)
                .forEach(function (transition) {
                    console.log(transition);
                });
        });
    }
);

