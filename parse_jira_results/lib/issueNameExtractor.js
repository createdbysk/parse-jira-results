/**
 * Given an issue, returns its name.
 */
define(function () {
    return function (issue, callback) {
        callback(null, issue.key);
    }
});