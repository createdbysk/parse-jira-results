// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var fs,
    linq,
    results;
fs = require('fs');
linq = require('linq');
results = JSON.parse(fs.readFileSync('test_input/SW-14155.txt', 'utf-8'));

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
console.log(JSON.stringify(issues.toArray(), undefined, 4));
