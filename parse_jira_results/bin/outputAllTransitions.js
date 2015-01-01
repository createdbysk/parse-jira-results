// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('../configuredRequirejs.js');

requirejs(['bin/readFileAndIterate',
          'linq', 
          'moment',
          'lib/issueStatusExtractor', 
          'lib/issueCreatedDateExtractor', 
          'lib/statusFilter',
          'lib/leadTimeCalculator'], 
    function (readFileAndIterate, linq, moment, issueStatusExtractor, issueCreatedDateExtractor, statusFilter, leadTimeCalculator) {
        'use strict';
        var processResults,
            displayResults,
            formatDate;
        formatDate = function (rawDate) {
            if (rawDate) {
                return moment(rawDate).format("YYYY-MM-DD hh:mm:ss");
            }
        };
        processResults = function (results) {
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
        };
        displayResults = function (err, allTransitions) {
            linq.from(allTransitions)
                .forEach(function (transition) {
                    console.log(transition);
                });
        };
        readFileAndIterate(processResults, displayResults);
    }
);

