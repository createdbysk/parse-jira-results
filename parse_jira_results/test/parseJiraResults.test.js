var fs,
    linq;
fs = require('fs');
linq = require('linq');
describe('parse jira results', function () {
    'use strict';
    var results;

    beforeEach(function (done) {
        results = JSON.parse(fs.readFileSync('test_input/SW-14155.txt', 'utf-8'));
        done();
    });
    it('reads extracts the first and last states of an issue status', function (done) {
        var issues;
        issues = linq.from(results.issues)
            .select(function (issue) {
                var statuses;
                statuses = linq.from(issue.changelog.histories)
                    .select(function (history) {
                        var status;
                        status = linq.from(history.items)
                            .where(function (item) {
                                return item.field === 'status';
                            })
                            .select(function (item) {
                                return {from: item.fromString, to: item.toString};
                            });
                        return {date: history.created, status: status.toArray()};
                    });
                return {key: issue.key, statuses: statuses.toArray()};
            });
        // TODO: Define expected output here and then update the linq queries to generate that output.
        console.log(JSON.stringify(issues.toArray(), undefined, 4));
    });
});