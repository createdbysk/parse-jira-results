define(['commander',
          'fs', 
          'linq'
          ], 
    function (program, fs, linq) {
        return function (processFn, callback) {
            'use strict';
            program
                .version('0.0.1')
                .parse(process.argv);
            fs.readFile(program.args[0], 'utf-8', function (err, resultsJSON) {
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

