define(['linq', 'lib/leadTimeCalculator'], function (linq, leadTimeCalculator) {
    'use strict';
    var durationInStatusCalculator;
    durationInStatusCalculator = function (timesInStatuses) {
        var durationInStatuses;

        durationInStatuses = 
            linq.from(timesInStatuses)
                .groupBy(
                    function (timeInStatus) {
                        return timeInStatus.status;
                    },
                    function (timeInStatus) {
                        return leadTimeCalculator(timeInStatus.from, timeInStatus.to);
                    },
                    function (status, durations) {
                        return {
                            status: status,
                            // If the last element is null, then set the whole duration to null
                            // because time in this status is not complete.
                            duration: durations.last() === null ? null : durations.sum()
                        };
                    }
                );

        return durationInStatuses;
    };
    return durationInStatusCalculator;
});