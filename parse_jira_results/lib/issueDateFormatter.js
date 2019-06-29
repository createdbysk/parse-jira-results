/**
 * Given a date, returns a date formatted with the following format
 * "YYYY-MM-DD hh:mm:ss"
 */
define(['moment'], function (moment) {
    'use strict';
    var formatString;

    formatString = "YYYY-MM-DD hh:mm:ss";

    /**
     * Given a date, returns a date formatted with the following format
     * @param  {String}     date The UTC date.
     * @param  {Function}   callback The callback function of the form function (err, formattedDate)
     */
    return function (date, callback) {
        var formattedDate;
        if (!date) {
            formattedDate = "";
        }
        else {
            formattedDate = moment(date).format(formatString);
            if (formattedDate === 'Invalid date') {
                formattedDate = "";
            }
        }
        callback(null, formattedDate);
    };
});
