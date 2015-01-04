var requireInjector,
    expect;
requireInjector = require('library/test_utilities/requireInjector');
expect = require('expect.js');

describe('time in status calculator', function () {
    var injector,
        timeInStatusCalculator;
    beforeEach(function (done) {
        injector = requireInjector.createInjector();
        injector.require(['lib/timeInStatusCalculator'], function (theTimeInStatusCalculator) {
            timeInStatusCalculator = theTimeInStatusCalculator;
            done();
        });
    });
    it('should set time to null if there is only no end for the status', function (done) {
        var statuses,
            expectedTimeInStatus;
        statuses = [
            {
                date: "date1",
                from: "from",
                to: "to"
            }
        ];
        expectedTimeInStatus = [
            {
                status: "to",
                timeInStatus: null
            }
        ];
        timeInStatusCalculator(statuses, function (err, timesInStatus) {
            expect(timesInStatus).to.eql(expectedTimeInStatus);
            done();
        });
    });
});