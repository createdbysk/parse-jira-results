/**
 * Given an issue, returns its created date.
 */
define(function () {
    return function (issue, callback) {
        callback(null, issue.fields.created);
    }
});