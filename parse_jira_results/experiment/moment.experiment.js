// Documentation here - https://code.google.com/p/datejs/wiki/APIDocumentation
var moment,
    dateToConvert,
    anotherDateToConvert,
    date1,
    date2,
    diff;
// This call adds 'static' methods to the JS Date class.
moment = require('moment');
// Display today.
console.log("NOW", moment().format("YYYY-MM-DD HH:MM:SS"));
// Convert a date
dateToConvert = '2014-12-01T15:58:25.000+0000';
date1 = moment(dateToConvert);
console.log(dateToConvert, date1.format("YYYY-MM-DD HH:MM:SS"), "DATE", date1.date());
// Lead time
anotherDateToConvert = '2014-12-18T12:12:21.000+0000';
date2 = moment(anotherDateToConvert);
console.log(anotherDateToConvert, date2.format("YYYY-MM-DD HH:MM:SS"), "DAY", date2.date());
console.log("LEAD TIME", date2.startOf('day').diff(date1.startOf('day'), 'days'));

date1 = "2014-12-00T15:58:25.000+0000";
date2 = "2014-12-17T12:12:21.000+0000";

diff = moment(date2).diff(moment(date1), 'days', true);
console.log('DIFF', diff);