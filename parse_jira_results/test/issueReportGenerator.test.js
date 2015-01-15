var requireInjector,
    expect,
    sinon,
    linq;
requireInjector = require('library/test_utilities/requireInjector');
expect = require('expect.js');
sinon = require('sinon');
linq = require('linq');

describe('issue report generator', function () {
    'use strict';
    var injector,
        status1,
        status2,
        nonExistentStatus,
        value1,
        value2,
        issueReportGenerator;
        
    beforeEach(function (done) {
        injector = requireInjector.createInjector();
        status1 = 'status1';
        status2 = 'status2';
        nonExistentStatus = 'nonExistentStatus';
        value1 = 1;
        value2 = 2;
        injector
            .mock('linq', linq)
            // .mock('lib/timeService', timeService)
            .require(['lib/issueReportGenerator'], function (theIssueReportGenerator) {
            issueReportGenerator = theIssueReportGenerator;
            done();
        });
    });
    it('should extract times for the given statuses into the report in the given order', 
        function (done) {
            var statusMetrics,
                nameOfFieldToExtract,
                report,
                expectedReport;
            nameOfFieldToExtract = 'fieldToExtract';
            statusMetrics = [
                {
                    status: status1,
                    fieldToExtract: value1
                },
                {
                    status: status2,
                    fieldToExtract: value2
                }
            ];
            expectedReport = {
                status1: value1,
                status2: value2
            }
            report = issueReportGenerator(statusMetrics, nameOfFieldToExtract);
            expect(report).to.eql(expectedReport);
            done();
        }
    );
});