/**
 * Given an issue, returns its priority.
 */
define(function () {
    return function (issue) {
        return issue.fields.priority.name;
    }
});