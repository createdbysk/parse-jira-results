/**
 * Given an issue and a given fieldToExtract transforms the statusMetrics to 
 * {statusName: fieldToExtract} pairs.
 */
define(['linq'], function (linq) {
    return function (statusMetrics, fieldToExtract) {
        return linq.from(statusMetrics)
                    .aggregate({},
                        function(combination, statusMetric) {
                            combination[statusMetric.status] = statusMetric[fieldToExtract];
                            return combination;
                        }
                    );
    };
});