/**
 * Given an issue, extracts the start date for the issue.
 */
define(['lib/issueDateExtractor'],
    function (issueDateExtractor) {
        'use strict';
        /**
         * Given an issue, extracts the start date for the issue
         * @param  {Objct}      issue The issue
         * @param  {Function}   callback The callback function of the form function (err, startDate)
         */
        return function (issueAndModuleConfiguration, callback) {
          issueDateExtractor(issueAndModuleConfiguration, "startDate", callback);
        };
    }
);
