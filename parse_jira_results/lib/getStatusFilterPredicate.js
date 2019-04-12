define(function () {
    'use strict';
    return function (moduleConfiguration, property, issueType, callback) {
      var filter;
      filter = undefined;
      if (moduleConfiguration.hasOwnProperty(property)) {
        var dateConfiguration;
        dateConfiguration = moduleConfiguration[property];
        if (dateConfiguration.hasOwnProperty(issueType)) {
          filter = dateConfiguration[issueType];
        }
        else {
          if (dateConfiguration.hasOwnProperty("Default")) {
            filter = dateConfiguration["Default"];
          }
        }
      }
      if (!filter) {
        callback("ERROR: moduleConfiguration does not have a filter defined for " + issueType + " for " + property + ".");
      }
      else {
        callback(null, filter);
      }
    }
});
