var requireInjector,
    expect,
    sinon,
    linq;
requireInjector = require('library/test_utilities/requireInjector');
expect = require('expect.js');
sinon = require('sinon');
linq = require('linq');

describe('time in status calculator', function () {
    'use strict';
    var injector,
        // timeService,
        date1,
        date2,
        terminalStatus,
        timeInStatusCalculator;
        
    beforeEach(function (done) {
        injector = requireInjector.createInjector();
        // timeService = {
        //     difference : function () {}
        // };
        //sinon.stub(timeService, 'difference');
        date1 = 'date1';
        date2 = 'date2';
        terminalStatus = 'terminalStatus';
        injector
            .mock('linq', linq)
            // .mock('lib/timeService', timeService)
            .require(['lib/timeInStatusCalculator'], function (theTimeInStatusCalculator) {
            timeInStatusCalculator = theTimeInStatusCalculator;
            done();
        });
    });
    it('should set to time to null if there is only one status', function (done) {
        var statuses,
            timesInStatus,
            expectedTimesInStatus;
        statuses = [
            {
                date: date1,
                from: "from",
                to: "to"
            }
        ];
        expectedTimesInStatus = [
            {
                status: "to",
                start: date1,
                end: null
            }
        ];
        timesInStatus = timeInStatusCalculator(statuses, terminalStatus);
        expect(timesInStatus).to.eql(expectedTimesInStatus);
        done();
    });
    it('should set from time from the first status in a pair and to time to second status in a pair', 
        function (done) {
            var statuses,
                timesInStatus,
                expectedTimesInStatus;
            statuses = [
                {
                    date: date1,
                    from: "from",
                    to: "to"
                },
                {
                    date: date2,
                    from: "to",
                    to: terminalStatus
                }
            ];
            expectedTimesInStatus = [
                {
                    status: "to",
                    start: date1,
                    end: date2
                }
            ];
            // timeService.difference.withArgs(date1, date2).returns(difference);
            timesInStatus = timeInStatusCalculator(statuses, terminalStatus);
            expect(timesInStatus).to.eql(expectedTimesInStatus);
            done();
        }
    );
    it('should set to time to null if final status is not a terminal status', 
        function (done) {
            var statuses,
                timesInStatus,
                status1,
                status2,
                status3,
                date3,
                expectedTimesInStatus;
            status1 = '1';
            status2 = '2';
            status3 = '3';
            date3 = 'date3';
            statuses = [
                {
                    date: date1,
                    from: "created",
                    to: status1
                },
                {
                    date: date2,
                    from: status1,
                    to: status2
                },
                {
                    date: date3,
                    from: status2,
                    to: status3
                }
            ];
            expectedTimesInStatus = [
                {
                    status: status1,
                    start: date1,
                    end: date2
                },
                {
                    status: status2,
                    start: date2,
                    end: date3
                },   
                {
                    status: status3,
                    start: date3,
                    end: null
                },   
            ];
            timesInStatus = timeInStatusCalculator(statuses, terminalStatus);
            expect(timesInStatus).to.eql(expectedTimesInStatus);
            done();
        }
    );
    it('should generate exceptions for missing statuses',
        function (done) {
            expect(timeInStatusCalculator).to.throwException(function (e) {
                expect(e.message).to.match(/statuses/);
                done();
            });
        }
    );
    it('should generate exceptions for missing terminal status parameter',
        function (done) {
            expect(timeInStatusCalculator).withArgs([]).to.throwException(function (e) {
                expect(e.message).to.match(/terminalStatus/);
                done();
            });
        }
    );
});