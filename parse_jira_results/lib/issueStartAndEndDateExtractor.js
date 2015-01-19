/**
 * Given the statuses for an issue, extracts the start and end date for the issue.
 */
define(['linq', 'lib/statusFilter'], function (linq, statusFilter) {
    /**
     * Given the statuses for an issue, extracts the start and end date for the issue
     * @param  {Array}      statuses The array of statuses
     * @param  {Function}   callback The callback function of the form function (err, startDate, endDate)
     */
    return function (statuses, callback) {
        var startDate,
            endDate;
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
        callback(null, startDate, endDate);
    };
});
