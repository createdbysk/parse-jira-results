var requirejs,
    streamEnumerableCreator;

requirejs = require('library/configuredRequirejs');
streamEnumerableCreator = require('library/streamEnumerableCreator');

requirejs(['lib/timeInStatusCalculator', 
          'lib/statusMetricsCalculator'], 
    function (timeInStatusCalculator, statusMetricsCalculator) {
        'use strict';
        streamEnumerableCreator(process.stdin, function (err, lines) {
            lines.select(function (line) {
                return JSON.parse(line);
            })
            .select(function (issue) {
                return {
                    name: issue.name,
                    timeInStatuses:  timeInStatusCalculator(issue.statuses, 'Closed')
                }
            })
            .select(function (issue) {
                return {
                    name: issue.name,
                    statusMetrics:  statusMetricsCalculator(issue.timeInStatuses).toArray()
                }
            })
            .forEach('console.log("value", $)');
        }
    );
});