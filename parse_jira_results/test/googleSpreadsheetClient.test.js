var requireInjector,
    expect,
    sinon;
requireInjector = require('library/test_utilities/requireInjector');
expect = require('expect.js');
sinon = require('sinon');

describe('google spreadsheet client factory', function () {
    'use strict';
    var injector,
        googleapis,
        GoogleSpreadsheet,
        googleSpreadsheetClientFactory;

    beforeEach(function (done) {
        googleapis = {
            auth : {
                JWT: function () {}
            }
        };

        GoogleSpreadsheet = sinon.stub();

        sinon.stub(googleapis.auth, 'JWT');
        injector = requireInjector.createInjector();
        injector
            .mock('googleapis', googleapis)
            .mock('google-spreadsheet', GoogleSpreadsheet)
            .require(['lib/googleSpreadsheetClient'], function (theGoogleSpreadsheetClientFactory) {
            googleSpreadsheetClientFactory = theGoogleSpreadsheetClientFactory;
            done();
        });
    });
    it('should exist',
        function (done) {
            expect(googleSpreadsheetClientFactory).to.be.ok();
            done();
        }
    );

    describe('createJwtClient', function () {
        var jwtConfiguration,
            clientEmail,
            clientPemFilePath;
        beforeEach(function (done) {
            clientEmail = 'abc@email.com';
            clientPemFilePath = 'a.pem';
            jwtConfiguration = {
                clientEmail: clientEmail,
                clientPemFilePath: clientPemFilePath
            }
            done();
        });
        it('should create a Jwt Client given configuration parameters', function (done) {
            var jwtClient;
            jwtClient = googleSpreadsheetClientFactory.createJwtClient(jwtConfiguration,
                function (err, jwtClient) {
                    sinon.assert.calledWithNew(googleapis.auth.JWT);
                    sinon.assert.calledWith(
                        googleapis.auth.JWT,
                        clientEmail,
                        clientPemFilePath,
                        null,
                        // scope from https://developers.google.com/google-apps/spreadsheets/#authorizing_requests_with_oauth_20
                        'https://spreadsheets.google.com/feeds'
                    );
                    expect(err).to.not.be.ok();
                    expect(jwtClient).to.be.ok();
                    done();
                }
            );
        });
    });
});
