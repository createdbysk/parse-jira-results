define(['linq', 'lib/leadTimeCalculator'], function (linq, leadTimeCalculator) {
    'use strict';
    var statusMetricsCalculator;
    statusMetricsCalculator = function (timesInStatuses) {
        var statusMetrics;

        statusMetrics = 
            linq.from(timesInStatuses)
                .groupBy(
                    function (timeInStatus) {
                        return timeInStatus.status;
                    },
                    function (timeInStatus) {
                        return leadTimeCalculator(timeInStatus.start, timeInStatus.end);
                    },
                    function (status, durations) {
                        return {
                            status: status,
                            numberOfEntries: durations.count(),
                            // If the last element is null, then set the whole duration to null
                            // because time in this status is not complete.
                            duration: durations.last() === null ? null : Math.round(durations.sum()*100)/100
                        };
                    }
                );

        return statusMetrics;
    };
    return statusMetricsCalculator;
});