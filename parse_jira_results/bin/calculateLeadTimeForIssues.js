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
        fs.readFile('test_input/SVS-500.txt', 'utf-8', function (err, resultsJSON) {
            var allResults,
                leadTimes;
            allResults = JSON.parse(resultsJSON);
            console.error("Number of JIRA queries", allResults.length);
            leadTimes = linq.from(allResults)
                .selectMany(function (results) {
                    var leadTimesForThisSet;
                    issueStatusExtractor(results, function (err, issuesWithStatuses) {
                        console.error(JSON.stringify(issuesWithStatuses.toArray(), undefined, 4));
                        leadTimesForThisSet = linq.from(issuesWithStatuses)
                            .select(function (issue) {
                                var startDate,
                                    endDate;
                                statusFilter(issue.statuses, 
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
                                statusFilter(issue.statuses, 
                                    function (status) {
                                        return status.to === "Closed" || status.to === "In Review";
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
                                return {key: issue.key, startDate: startDate, endDate: endDate, leadTime: leadTimeCalculator(startDate, endDate)};
                            });
                    });
                    return leadTimesForThisSet;
                });
            console.error("Total number of records, including no leadtime values", leadTimes.count());
            console.log('ISSUE, COMMIT, CLOSE, LEAD TIME')
            linq.from(leadTimes)
                .forEach(function (record) {
                    if (record.leadTime) {
                        console.log('%s, %s, %s, %d', record.key, 
                                    moment(record.startDate).format('YYYY-MM-DD'), 
                                    moment(record.startDate).format('YYYY-MM-DD'), 
                                    record.leadTime);                            
                    }                        
                });
        });
    }
);

