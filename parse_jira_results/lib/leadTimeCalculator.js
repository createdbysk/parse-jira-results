define(['linq', 'moment'], function (linq, moment) {
    'use strict';
    var leadTimeCalculator;

    /**
     * Calculates the number of days, including fractional days,
     * between the startDate and endDate. Returns null if either
     * startDate or endDate is null.
     * @param  {date} startDate     The start date
     * @param  {date} endDate       The end date
     * @return {days}               The difference, including fraction, in number of days between
     *                              startDate and endDate
     */
    leadTimeCalculator = function (startDate, endDate) {
        var startMoment,
            endMoment,
            leadTime;
        startMoment = moment(startDate);
        endMoment = moment(endDate);
        // See http://momentjs.com/docs/
        // Second parameter specifies get the number of days.
        // Third parameter specifies return fractional days.
        leadTime = endMoment.diff(startMoment, 'days', true);
        return isNaN(leadTime) ? null : leadTime;
    }
    return leadTimeCalculator;
});
