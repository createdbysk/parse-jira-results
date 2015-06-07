var requireInjector,
    expect,
    sinon,
    moment;
requireInjector = require('library/test_utilities/requireInjector');
expect = require('expect.js');
sinon = require('sinon');
moment = require('moment');

describe('issue date formatter', function () {
    'use strict';
    var injector,
        issueDateFormatter;

    beforeEach(function (done) {
        injector = requireInjector.createInjector();
        injector
            .mock('moment', moment)
            .require(['lib/issueDateFormatter'], function (theIssueDateFormatter) {
                issueDateFormatter = theIssueDateFormatter;
                done();
            });
    });
    it('should format the given date',
        function (done) {
            var date,
                expectedFormattedDate;
            // Date in UTC.
            date = "2014-12-18T12:12:21.000+0000";
            expectedFormattedDate = "2014-12-18 06:12:21";
            issueDateFormatter(date, function (err, formattedDate) {
                expect(formattedDate).to.eql(expectedFormattedDate);
                done();
            });
        }
    );
    it('should return the original value if the Invalid date',
        function (done) {
            var date,
                expectedFormattedDate;
            // Invalid date.
            date = null;
            expectedFormattedDate = null;
            issueDateFormatter(date, function (err, formattedDate) {
                expect(formattedDate).to.eql(expectedFormattedDate);
                done();
            });
        }
    );
});
