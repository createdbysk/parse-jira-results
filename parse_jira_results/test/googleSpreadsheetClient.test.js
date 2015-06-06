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

    describe('createClient', function () {
        var configuration,
            clientEmail,
            clientPemFilePath,
            jwtClient,
            tokens;
        beforeEach(function (done) {
            clientEmail = 'abc@email.com';
            clientPemFilePath = 'a.pem';
            configuration = {
                clientEmail: clientEmail,
                clientPemFilePath: clientPemFilePath
            }
            jwtClient = {
                authorize: function () {}
            };

            sinon.stub(jwtClient, "authorize");
            tokens = {
                token_type: "type",
                access_token: "value"
            }
            jwtClient.authorize.withArgs(sinon.match.typeOf('function'))
                                .callsArgWith(0, null, tokens);
            googleapis.auth.JWT.returns(jwtClient);

            done();
        });
        it('should create a Client given configuration parameters', function (done) {
            googleSpreadsheetClientFactory.createClient(configuration,
                function (err, client) {
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
                    expect(client).to.be.ok();
                    done();
                }
            );
        });
        describe('Client', function () {
            var client,
                spreadsheetName;
            beforeEach(function (done) {
                googleSpreadsheetClientFactory.createClient(configuration,
                    function (err, theClient) {
                        client = theClient;
                        done();
                    }
                );
                spreadsheetName = "spreadsheetName";
            });
            it('should create a spreadsheet with the given key', function (done) {
                client.getSpreadsheet(spreadsheetName,
                    function (err, spreadsheet) {
                        var auth;
                        auth = {
                            type: tokens.token_type,
                            value: tokens.access_token
                        };
                        sinon.assert.calledWithNew(GoogleSpreadsheet);
                        sinon.assert.calledWith(
                            GoogleSpreadsheet,
                            spreadsheetName,
                            auth
                        );
                        expect(err).to.not.be.ok();
                        expect(spreadsheet).to.be.ok();
                        done();
                    }
                );
            });
        });
    });
});
