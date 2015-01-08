var requirejs,
    streamEnumerableCreator;

requirejs = require('library/configuredRequirejs');
streamEnumerableCreator = require('library/streamEnumerableCreator');

requirejs(['lib/timeInStatusCalculator'], function (timeInStatusCalculator) {
    streamEnumerableCreator(process.stdin, function (err, lines) {
        lines.select(function (line) {
            return JSON.parse(line);
        })
    });
});