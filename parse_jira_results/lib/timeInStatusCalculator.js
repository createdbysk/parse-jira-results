define(function () {
    var timeInStatusCalculator;
    timeInStatusCalculator = function (statuses, callback) {
        callback(null, [
            {
                status: statuses[0].to,
                timeInStatus: null
            }
        ]);
    };
    return timeInStatusCalculator;
});