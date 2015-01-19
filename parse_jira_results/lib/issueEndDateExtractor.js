/**
 * Given an issue, extracts the end date for the issue.
 */
define(['linq', 'lib/statusFilter', 'lib/issueStatusExtractor'], function (linq, statusFilter, issueStatusExtractor) {
    'use strict';
    /**
     * Given an issue, extracts the end date for the issue
     * @param  {Object}     issue The issue
     * @param  {Function}   callback The callback function of the form function (err, endDate)
     */
    return function (issue, callback) {
        var endDate;
        issueStatusExtractor(issue, function (err, statuses) {
            if (err) {
                callback(err);
            }
            else {
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
                callback(null, endDate);
            }
        });
    };
});
