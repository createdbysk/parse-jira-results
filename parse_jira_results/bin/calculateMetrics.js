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
                var timeInStatuses = timeInStatusCalculator(issue.statuses, 'Closed');
                issue.statusMetrics = statusMetricsCalculator(timeInStatuses).toArray();
                delete issue.statuses;
                return issue;
            })
            .forEach('console.log("value", $)');
        }
    );
});