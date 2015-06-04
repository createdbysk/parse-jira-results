var requireInjector,
    expect,
    sinon;
requireInjector = require('library/test_utilities/requireInjector');
expect = require('expect.js');
sinon = require('sinon');

describe('google spreadsheet client', function () {
    'use strict';
    var injector,
        googleSpreadsheetClient;

    beforeEach(function (done) {
        injector = requireInjector.createInjector();
        injector
            .require(['lib/googleSpreadsheetClient'], function (theGoogleSpreadsheetClient) {
            googleSpreadsheetClient = theGoogleSpreadsheetClient;
            done();
        });
    });
    it('should exist',
        function (done) {
            expect(googleSpreadsheetClient).to.be.ok();
            done();
        }
    );
});
