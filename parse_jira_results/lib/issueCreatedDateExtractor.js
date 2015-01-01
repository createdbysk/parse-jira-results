/**
 * Given an issue, returns its created date.
 */
define(function () {
    return function (issue) {
        return issue.fields.created;
    }
});