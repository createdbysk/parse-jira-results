// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('../configuredRequirejs.js');

requirejs(['fs', 
          'linq', 
          'lib/issueStatusExtractor', 
          'lib/statusFilter',
          'lib/leadTimeCalculator'], 
    function (fs, linq, issueStatusExtractor, statusFilter, leadTimeCalculator) {
        fs.readFile('test_input/SW-14155.txt', 'utf-8', function (err, resultsJSON) {
            var results;
            results = JSON.parse(resultsJSON);
            issueStatusExtractor(results, function (err, issuesWithStatuses) {
                var leadTimes;
                leadTimes = linq.from(issuesWithStatuses)
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
                                startDate = linq.from(possibleCommitmentPoints)
                                                .select(function (status) {
                                                    return status.date;
                                                })
                                                .firstOrDefault();
                            }
                        );
                        statusFilter(issue.statuses, 
                            function (status) {
                                return status.to === "Closed";
                            },
                            function (error, possibleExitPoints) {
                                endDate = linq.from(possibleExitPoints)
                                                .select(function (status) {
                                                    return status.date;
                                                })
                                                .lastOrDefault();
                            }
                        );
                        return {key: issue.key, leadTime: leadTimeCalculator(startDate, endDate)};
                    });
                console.log(JSON.stringify(leadTimes.toArray(), undefined, 4));
            });
        })
    }
);

