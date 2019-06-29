/**
 * Given an issue, returns its created date.
 */
define(function () {
    return function (issueAndModuleConfiguration, callback) {
        callback(null, issueAndModuleConfiguration.issue.fields.created);
    }
});
