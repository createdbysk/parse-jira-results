define(['fs', 
        'linq'
       ], 
    function (fs, linq) {
        return function (filename, processFn, callback) {
            'use strict';
            fs.readFile(filename, 'utf-8', function (err, resultsJSON) {
                if (err) {
                    callback(err);
                }
                else {
                    var allResults,
                        allTransitions;
                    allResults = JSON.parse(resultsJSON);
                    if (!allResults.length) {
                        allResults = [allResults];
                    }
                    console.error("Number of JIRA queries", allResults.length);
                    allTransitions = linq.from(allResults)
                        .selectMany(processFn);
                    callback(null, allTransitions);
                }
            });   
        };  
    }
);

