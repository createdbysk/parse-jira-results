define(['linq', 'lib/issueCreatedDateExtractor'], function (linq, issueCreatedDateExtractor) {
    return function (issue, callback) {
        var statuses,
            initialStatus;
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
        // Make the initial transition from created to Triage at the created date 
        // with the assumption that the start state is Triage.
        issueCreatedDateExtractor(issue, function (err, createdDate) {
            initialStatus = {
                date: createdDate,
                from: "Created",
                to: "Triage"
            };
            // Add the start state to the beginning of statuses.
            statuses = linq.make(initialStatus).concat(statuses);
            callback(null, statuses);
        });
    };
});
