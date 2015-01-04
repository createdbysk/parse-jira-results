define(['linq'], function (linq) {
    'use strict';
    var timeInStatusCalculator;
    /**
     * Given a list of status transitions, calculates the completed time spent in each status.
     *
     * Preconditions:
     *     * For each pair of statuses, the to of the first element matches the from of the second.
     * 
     * If the final status is not the finalStatus, then the timeInStatus for the final "to" status will be null,
     * which indicates that it is ongoing and the completed time in that status is not known.
     * 
     * @param  {array of strings}   statuses        Required: array of status.
     * @param  {string}             terminalStatus  Required: expected 'to' status of the final element. 
     * @return {array of objects}   An array of objects of type {status: <string>, timeInStatus: <string>}
     */
    timeInStatusCalculator = function (statuses, terminalStatus) {
            var timesInStatus,
                lastStatus,
                enumerable;
        if (!statuses) {
            throw new Error('Expected statuses to be passed');
        }
        if (!terminalStatus) {
            throw new Error('Expected terminalStatus to be passed');
        }
        enumerable = linq.from(statuses);
        timesInStatus = 
            enumerable.takeExceptLast().
                zip(enumerable.skip(1), 
                    function (first, second) {
                        return {
                            status: first.to,
                            start: first.date,
                            end: second.date
                        };
                    }
                );
        lastStatus = enumerable.lastOrDefault(null);
        if (null !== lastStatus) {
            if (lastStatus.to !== terminalStatus) {
                timesInStatus = timesInStatus.concat(linq.make({
                    status: lastStatus.to,
                    start: lastStatus.date,
                    end: null
                }));
            }
        }
        return timesInStatus.toArray();
    };
    return timeInStatusCalculator;
});