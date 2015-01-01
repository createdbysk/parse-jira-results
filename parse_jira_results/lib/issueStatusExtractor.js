define(['linq'], function (linq) {
    return function (issue, callback) {
        var statuses;
        statuses = linq.from(issue.changelog.histories)
            .selectMany(function (history) {
                var status;
                status = linq.from(history.items)
                    .where(function (item) {
                        return item.field === 'status';
                    })
                    .select(function (item) {
                        return {date: history.created, from: item.fromString, to: item.toString};
                    });
                return status;
            });
        callback(null, statuses);
    };
});
