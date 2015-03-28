/**
 * Given an issue, returns its priority.
 */
define(function () {
    return function (issue, callback) {
        callback(null, issue.fields.issuetype.name);
    }
});
