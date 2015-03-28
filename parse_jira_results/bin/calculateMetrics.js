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

requirejs(['linq',
          'moment',
          'lib/timeInStatusCalculator',
          'lib/statusMetricsCalculator',
          'lib/issueReportGenerator'
          ],
    function (linq, moment, timeInStatusCalculator, statusMetricsCalculator, issueReportGenerator) {
        'use strict';
        streamEnumerableCreator(process.stdin, function (err, lines) {
            console.log('Name, Type, Story Points, Created Date, Start Date, End Date, Priority,' + program.order.join(', '));
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
                        startDate,
                        endDate,
                        formatString,
                        initialString;
                    formatString = 'YYYY-MM-DD HH:mm:ss';
                    createdDate = moment(issue.createdDate).format(formatString);
                    startDate = issue.startDate ? moment(issue.startDate).format(formatString) : null;
                    endDate = issue.endDate ? moment(issue.endDate).format(formatString) : null;

                    initialString = [issue.name, issue.type, issue.storyPoints, createdDate, startDate, endDate, issue.priority].join(',');
                    statusData =
                        linq.from(program.order)
                            .aggregate(initialString,
                                       function (line, status) {
                                            status = issue.report[status];
                                            if (status === null) {
                                                status = "-1"
                                            }
                                            else if (status === undefined) {
                                                status = '-2';
                                            }
                                            return line += ', ' + status;
                                        }
                            );
                    console.log(statusData);
                }
            );
        }
    );
});
