var requirejs,
    program,
    convertToArrayOfStatuses,
    streamEnumerableCreator;

requirejs = require('library/configuredRequirejs');
streamEnumerableCreator = require('library/streamEnumerableCreator');
program = require('commander');

convertToArrayOfStatuses = function (commaSeparatedStatuses) {
    return commaSeparatedStatuses.split(',');
};

program
    .usage('--order <comma separated statuses> <name of field to report>')
    .option('-o --order <comma separated statuses>', 'Comma separated list of statuses', convertToArrayOfStatuses)
    .parse(process.argv);

requirejs(['lib/timeInStatusCalculator', 
          'lib/statusMetricsCalculator',
          'lib/issueReportGenerator',
          'linq',
          'moment'], 
    function (timeInStatusCalculator, statusMetricsCalculator, issueReportGenerator, linq, moment) {
        'use strict';
        streamEnumerableCreator(process.stdin, function (err, lines) {
            console.log('Name, Created Date, ' + program.order.join(', '));
            lines.select(function (line) {
                return JSON.parse(line);
            })
            .select(function (issue) {
                var timeInStatuses,
                    statusMetrics,
                    report;
                timeInStatuses = timeInStatusCalculator(issue.statuses, 'Closed');
                statusMetrics = statusMetricsCalculator(timeInStatuses)
                issue.report = issueReportGenerator(statusMetrics, program.args[0]);
                delete issue.statuses;
                return issue;
            })
            .forEach(function (issue) 
                {
                    var statusData,
                        createdDate;
                    createdDate = moment(issue.createdDate).format('YYYY-MM-DD');
                    statusData = 
                        linq.from(program.order)
                            .aggregate(issue.name + ', ' + createdDate,
                                       function (line, status) {
                                            return line += ', ' + issue.report[status];
                                        }
                            );
                    console.log(statusData);
                }
            );
        }
    );
});