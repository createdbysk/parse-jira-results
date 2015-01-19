/**
 * Given an issue, extracts the start date for the issue.
 */
define(['linq', 'lib/statusFilter', 'lib/issueStatusExtractor'], function (linq, statusFilter, issueStatusExtractor) {
    'use strict';
    /**
     * Given an issue, extracts the start date for the issue
     * @param  {Objct}      issue The issue
     * @param  {Function}   callback The callback function of the form function (err, startDate)
     */
    return function (issue, callback) {
        var startDate;
        issueStatusExtractor(issue, function (err, statuses) {
            if (err) {
                callback(err);
            }
            else {
                statusFilter(statuses, 
                    function (status) {
                        return status.from === "Open" &&
                            status.to !== "Triage" &&
                            status.to !== "Closed";
                    },
                    function (error, possibleCommitmentPoints) {
                        console.error("COMMITMENT %s", JSON.stringify(possibleCommitmentPoints.toArray(), undefined, 4));
                        startDate = linq.from(possibleCommitmentPoints)
                                        .select(function (status) {
                                            return status.date;
                                        })
                                        .firstOrDefault();
                        console.error("COMMITMENT", JSON.stringify(possibleCommitmentPoints.toArray(), undefined, 4));
                        console.error("START DATE", startDate);
                    }
                );
                callback(null, startDate);
            }
        });
    };
});
