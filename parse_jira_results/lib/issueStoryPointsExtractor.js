define(['linq'], function (linq) {
    return function (issue, callback) {
        var storyPoints;
        storyPoints = linq.from(issue.changelog.histories)
            .selectMany(function (history) {
                var storyPoints;
                storyPoints = linq.from(history.items)
                    .where(function (item) {
                        return item.field === 'Story Points';
                    })
                    .select(function (item) {
                        console.error("Story points", item.toString);
                        return item.toString;
                    });
                return storyPoints;
            })
            .lastOrDefault(null);
        callback(null, storyPoints);
    };
});
