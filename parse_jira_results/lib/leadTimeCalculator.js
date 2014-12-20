define(['linq', 'moment'], function (linq, moment) {
    'use strict';
    var leadTimeCalculator;

    leadTimeCalculator = function (startDate, endDate) {
        var startMoment,
            endMoment,
            leadTime;
        startMoment = moment(startDate).startOf('day');
        endMoment = moment(endDate).startOf('day');
        // See http://momentjs.com/docs/
        // Second parameter says get the number of days.
        leadTime = Math.ceil(endMoment.diff(startMoment, 'days'));
        return isNaN(leadTime) ? null : leadTime;
    }
    return leadTimeCalculator;
});
