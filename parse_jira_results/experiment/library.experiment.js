(function () {
    'use strict';
    var createStreamEnumerable,
        requireInjector,
        injector,
        expect;
    createStreamEnumerable = require('library/streamEnumerableCreator');

    createStreamEnumerable(process.stdin, function (err, streamEnumerable) {
        if (err) {
            console.error('ERROR: ', err)
        }
        else {
            streamEnumerable.forEach(function (line) {
                console.error("LINE :", line);
            });            
        }
    });
    expect = require('expect.js');
    injector = require('library/test_utilities/requireInjector').createInjector();
    injector.require(['experiment/helper/library.experiment.require-helper'], function (success) {
        console.log(success);
        expect(success).to.be('success');
    });
})();
