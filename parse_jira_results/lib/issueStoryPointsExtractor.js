define(['linq'], function (linq) {
    return function (issueAndModuleConfiguration, callback) {
        var storyPoints,
            issue;
        issue = issueAndModuleConfiguration.issue;
        storyPoints = linq.from(issue.changelog.histories)
            .selectMany(function (history) {
                var storyPoints;
                storyPoints = linq.from(history.items)
                    .where(function (item) {
                        return item.field === 'Story Points';
                    })
                    .select(function (item) {
                        return item.toString;
                    });
                return storyPoints;
            })
            .lastOrDefault(null);
        callback(null, storyPoints);
    };
});
