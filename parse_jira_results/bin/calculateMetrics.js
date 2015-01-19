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
            console.log('Name, Created Date, Story Points, ' + program.order.join(', '));
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
                        createdDate,
                        initialString;
                    createdDate = moment(issue.createdDate).format('YYYY-MM-DD');
                    initialString = [issue.name, createdDate, issue.storyPoints].join(',');
                    statusData = 
                        linq.from(program.order)
                            .aggregate(initialString,
                                       function (line, status) {
                                            status = issue.report[status];
                                            if (status === null) {
                                                status = "-1"
                                            }
                                            else if (status === undefined) {
                                                status = '';
                                            }
                                            return line += ', ' + status;
                                        }
                            );
                    console.log(statusData);
                }
            );
            console.log('Lead Time Formula, =IF(COUNTIF(D2:I2,"= active")>0,"",SUM(D2:I2))');
        }
    );
});