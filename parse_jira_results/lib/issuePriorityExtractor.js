/**
 * Given an issue, returns its priority.
 */
define(function () {
    return function (issueAndModuleConfiguration, callback) {
        callback(null, issueAndModuleConfiguration.issue.fields.priority.name);
    }
});
