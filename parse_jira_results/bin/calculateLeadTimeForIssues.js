// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('../configuredRequirejs.js');

requirejs(['bin/readFileAndIterate',
          'linq', 
          'moment',
          'lib/issueStatusExtractor', 
          'lib/statusFilter',
          'lib/leadTimeCalculator'], 
    function (readFileAndIterate, linq, moment, issueStatusExtractor, statusFilter, leadTimeCalculator) {
        var processResults,
            displayResults;
        processResults = function (results) {
            var leadTimesForThisSet;
            leadTimesForThisSet = 
                linq.from(results.issues)
                    .select(function (issue) {
                        issueStatusExtractor(issue, function (err, statuses) {
                            console.error(JSON.stringify(statuses.toArray(), undefined, 4));
                            statusFilter(statuses, 
                                function (status) {
                                    return status.from === "Open" &&
                                        status.to !== "Triage" &&
                                        status.to !== "Closed";
                                },
                                function (error, possibleCommitmentPoints) {
                                    console.error("%s COMMITMENT %s", issue.key, JSON.stringify(possibleCommitmentPoints.toArray(), undefined, 4));
                                    startDate = linq.from(possibleCommitmentPoints)
                                                    .select(function (status) {
                                                        return status.date;
                                                    })
                                                    .firstOrDefault();
                                    console.error("COMMITMENT", JSON.stringify(possibleCommitmentPoints.toArray(), undefined, 4));
                                    console.error("START DATE", startDate);
                                }
                            );
                            statusFilter(statuses, 
                                function (status) {
                                    return status.to === "Closed";
                                },
                                function (error, possibleExitPoints) {
                                    console.error("EXIT", JSON.stringify(possibleExitPoints.toArray(), undefined, 4));
                                    endDate = linq.from(possibleExitPoints)
                                                    .select(function (status) {
                                                        return status.date;
                                                    })
                                                    .lastOrDefault();
                                    console.error("END DATE", endDate);
                                }
                            );
                        });
                        return {key: issue.key, startDate: startDate, endDate: endDate, leadTime: leadTimeCalculator(startDate, endDate)};
                    });
            return leadTimesForThisSet;
        };
        displayResults = function (err, leadTimes) {
            console.error("Total number of records, including no leadtime values", leadTimes.count());
            console.log('ISSUE, COMMIT, CLOSE, LEAD TIME');
            linq.from(leadTimes)
                .forEach(function (record) {
                    if (record.leadTime) {
                        console.log('%s, %s, %s, %d', record.key, 
                                    moment(record.startDate).format('YYYY-MM-DD'), 
                                    moment(record.endDate).format('YYYY-MM-DD'), 
                                    record.leadTime);                            
                    }                        
                });
        };
        readFileAndIterate(processResults, displayResults);
    }
);

