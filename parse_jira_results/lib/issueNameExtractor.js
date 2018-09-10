/**
 * Given an issue, returns its name.
 */
define(function () {
    return function (issueAndModuleConfiguration, callback) {
        callback(null, issueAndModuleConfiguration.issue.key);
    }
});
