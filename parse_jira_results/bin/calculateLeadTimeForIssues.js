// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('../configuredRequirejs.js');

requirejs(['fs', 'linq', 'lib/issueStatusExtractor', 'lib/statusFilter'], 
    function (fs, linq, issueStatusExtractor, statusFilter) {
        fs.readFile('test_input/SW-14155.txt', 'utf-8', function (err, resultsJSON) {
            var results;
            results = JSON.parse(resultsJSON);
            issueStatusExtractor(results, function (err, issuesStatus) {
                console.log(JSON.stringify(issuesStatus.toArray(), undefined, 4));
                linq.from(issuesStatus)
                    .select(function (issue) {
                    })
            });
        })
    }
);

