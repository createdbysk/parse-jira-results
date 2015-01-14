define(['linq'], function (linq) {
    return function (issue, callback) {
        var storyPoints;
        storyPoints = linq.from(issue.changelog.histories)
            .select(function (history) {
                var latestStoryPoints;
                latestStoryPoints = linq.from(history.items)
                    .where(function (item) {
                        return item.field === 'Story Points';
                    })
                    .select(function (item) {
                        return item.toString;
                    })
                    .lastOrDefault(null);
                return latestStoryPoints;
            })
            .lastOrDefault(null);
        callback(null, storyPoints);
    };
});
