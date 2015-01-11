var requireInjector,
    expect,
    sinon,
    linq;
requireInjector = require('library/test_utilities/requireInjector');
expect = require('expect.js');
sinon = require('sinon');
linq = require('linq');

describe('status metrics calculator', function () {
    'use strict';
    var injector,
        leadTimeCalculator,
        date1,
        date2,
        date3,
        status1,
        status2,
        duration12,
        duration23,
        terminalStatus,
        statusMetricsCalculator,
        verify;

    verify = function (timesInStatuses, expectedStatusMetricses, done) {
        var statusMetrics;
        statusMetrics = statusMetricsCalculator(timesInStatuses).toArray();
        expect(statusMetrics).to.eql(expectedStatusMetricses);
        done();
    }
        
    beforeEach(function (done) {
        injector = requireInjector.createInjector();
        date1 = 'date1';
        date2 = 'date2';
        status1 = 'status1';
        status2 = 'status2';
        duration12 = 1;
        duration23 = 2;
        terminalStatus = 'terminalStatus';
        leadTimeCalculator = sinon.stub();
        leadTimeCalculator.withArgs(sinon.match.string, null).returns(null);
        leadTimeCalculator.withArgs(date1, date2).returns(duration12);
        leadTimeCalculator.withArgs(date2, date3).returns(duration23);
        injector
            .mock('linq', linq)
            .mock('lib/leadTimeCalculator', leadTimeCalculator)
            .require(['lib/statusMetricsCalculator'], function (theStatusMetricsCalculator) {
            statusMetricsCalculator = theStatusMetricsCalculator;
            done();
        });
    });
    it('should set leadtime to null and number of entries to 1 if the "to date" is null', function (done) {
        var timesInStatuses,
            expectedStatusMetrics;
        timesInStatuses = [{
            status: status1,
            start: date1,
            end: null
        }];
        expectedStatusMetrics = [
            {
                status: status1,
                numberOfEntries: 1,
                duration: null
            }
        ];
        verify(timesInStatuses, expectedStatusMetrics, done);
    });
    it('should set leadtime to the difference of the from and to time for a status', function (done) {
        var timesInStatuses,
            statusMetrics,
            expectedStatusMetrics;
        timesInStatuses = [
            {
                status: status1,
                start: date1,
                end: date2
            }
        ];
        expectedStatusMetrics = [
            {
                status: status1,
                numberOfEntries: 1,
                duration: duration12
            }
        ];
        verify(timesInStatuses, expectedStatusMetrics, done);
    });
    it('should set leadtime to the sum of the durations and the number of entries for each status', function (done) {
        var timesInStatuses,
            statusMetrics,
            expectedStatusMetrics;
        timesInStatuses = [
            {
                status: status1,
                start: date1,
                end: date2
            },
            {
                status: status1,
                start: date2,
                end: date3
            },
        ];
        expectedStatusMetrics = [
            {
                status: status1,
                numberOfEntries: 2,
                duration: duration12+duration23
            }
        ];
        verify(timesInStatuses, expectedStatusMetrics, done);
    });
    it('should set leadtime to the sum of the durations for each status even when the statuses are non-contiguous', function (done) {
        var timesInStatuses,
            expectedStatusMetrics;
        timesInStatuses = [
            {
                status: status1,
                start: date1,
                end: date2
            },
            {
                status: status2,
                start: date1,
                end: date2
            },
            {
                status: status1,
                start: date2,
                end: date3
            },
        ];
        expectedStatusMetrics = [
            {
                status: status1,
                numberOfEntries: 2,
                duration: duration12+duration23
            },
            {
                status: status2,
                numberOfEntries: 1,
                duration: duration12
            }
        ];
        verify(timesInStatuses, expectedStatusMetrics, done);
    });
    it('should set leadtime to null if the "to date" of the last status is null', function (done) {
        var timesInStatuses,
            expectedStatusMetrics;
        timesInStatuses = [
            {
                status: status1,
                start: date1,
                end: date2
            },
            {
                status: status2,
                start: date1,
                end: date2
            },
            {
                status: status1,
                start: date2,
                end: null
            },
        ];
        expectedStatusMetrics = [
            {
                status: status1,
                numberOfEntries: 2,
                duration: null
            },
            {
                status: status2,
                numberOfEntries: 1,
                duration: duration12
            }
        ];
        verify(timesInStatuses, expectedStatusMetrics, done);
    });
});