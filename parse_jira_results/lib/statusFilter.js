define(['linq'], function (linq) {
    'use strict';
    var statusFilter;

    statusFilter = function (statuses, predicate, callback) {
        var filteredStatuses;
        filteredStatuses = 
            linq.from(statuses)
                .where(function (status) {
                    return predicate(status);
                });
        callback(null, filteredStatuses);
    };

    return statusFilter;
});
