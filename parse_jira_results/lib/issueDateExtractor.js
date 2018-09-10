/**
 * Given an issue, extracts the start date for the issue.
 */
define(['linq',
        'lib/statusFilter',
        'lib/getStatusFilterPredicate',
        'lib/issueTypeExtractor',
        'lib/issueStatusExtractor',
        'lib/issueDateFormatter'],
    function (linq, statusFilter, getStatusFilterPredicate, issueTypeExtractor, issueStatusExtractor, issueDateFormatter) {
        'use strict';
        /**
         * Given an issue, extracts the start date for the issue
         * @param  {Objct}      issue The issue
         * @param  {String}     whichDate startDate or endDate
         * @param  {Function}   callback The callback function of the form function (err, startDate)
         */
        return function (issueAndModuleConfiguration, whichDate, callback) {
            issueTypeExtractor(issueAndModuleConfiguration, function (err, issueType) {
              issueStatusExtractor(issueAndModuleConfiguration, function (err, statuses) {
                  if (err) {
                      callback(err);
                  }
                  else {
                      getStatusFilterPredicate(issueAndModuleConfiguration.moduleConfiguration,
                         whichDate, issueType,
                         function(err, filter) {
                           if (err) {
                             callback(err);
                           }
                           statusFilter(statuses,
                               filter,
                               function (error, possibleCommitmentPoints) {
                                   var startDate;
                                   startDate = linq.from(possibleCommitmentPoints)
                                                   .select(function (status) {
                                                       return status.date;
                                                   })
                                                   .firstOrDefault();
                                   issueDateFormatter(startDate,
                                       function (err, formattedStartDate) {
                                           callback(null, formattedStartDate);
                                       }
                                   );
                               }
                           );
                        }
                      )
                  }
              });
            })
        };
    }
);
