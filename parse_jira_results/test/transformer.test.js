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
        transformer;

    beforeEach(function (done) {
        injector = requireInjector.createInjector();

        injector
            .mock('linq', linq)
            // .mock('lib/timeService', timeService)
            .require(['lib/transformer'], function (theTransformer) {
            transformer = theTransformer;
            done();
        });
    });
    it('should apply the given transforms and return the value in the corresponding named field.',
        function (done) {
            var input,
                transform1,
                transform2,
                output1,
                output2,
                transforms,
                expectedResults,
                results;
            input = 'input';
            transform1 = sinon.stub();
            output1 = 'output1';
            transform1.withArgs(input, sinon.match.typeOf('function'))
                      .callsArgWith(1, output1);
            transform2 = sinon.stub();
            output2 = 'output2';
            transform2.withArgs(input, sinon.match.typeOf('function'))
                      .callsArgWith(1, output2);
            transforms = {
                transform1: transform1,
                transform2: transform2
            };
            expectedResults = {
                transform1: output1,
                transform2: output2
            };
            results = transformer(input, transforms);
            expect(results).to.eql(expectedResults);
            done();
        }
    );
});
