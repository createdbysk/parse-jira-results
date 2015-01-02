// ALWAYS RUN THIS WITH CURRENT WORKING DIRECTORY AS THE ROOT OF THIS PROJECT.
var requirejs;
requirejs = require('library/configuredRequirejs');

requirejs(['commander',
          'bin/readFileAndIterate',
          'linq',
          'path',
          'lib/issueStatusExtractor', 
          'lib/issuePriorityExtractor'
          ], 
    function (program, readFileAndIterate, linq, path, issueStatusExtractor, issuePriorityExtractor) {
        'use strict';
        var formatDate,
            extractors,
            processIssue,
            processResults,
            displayResults;
        extractors = function (value, collection) {
            var withoutExtension,
                module,
                name;
            value = value.split(',');
            module = value[0];
            name = value[1];
            withoutExtension = path.join(path.dirname(module), path.basename(module, '.js'));
            collection.modules.push(withoutExtension);
            collection.names.push(name);
            return collection;
        };

        program
            .version('0.0.1')
            .option('-e --extractor [module,name]', 'requirejs modules that extract values', extractors, {modules:[], names:[]})
            .parse(process.argv);

        requirejs(program.extractor.modules, function () {
            var args;
            args = arguments;
            processIssue = 
                function (issue) {
                    var resultsWithExtractedFields;
                    resultsWithExtractedFields = 
                        linq.from(program.extractor.names)
                            .zip(linq.from(args), 
                                function (name, extractor) {
                                    var retval;
                                    extractor(issue, function (err, result) {
                                        retval = {name: name, 
                                                  result: (result instanceof linq) ? result.toArray() : result
                                                };
                                    });
                                    console.error('result', resultsWithExtractedFields);
                                    return retval;
                                }
                            )
                            .aggregate({}, function (combination, value) {
                                combination[value.name] = value.result;
                                return combination;
                            });
                    console.error('result', resultsWithExtractedFields);
                    return resultsWithExtractedFields;
                };
            processResults = function (results) {
                var issuesWithExtractFields;
                issuesWithExtractFields = 
                    linq.from(results.issues)
                        .select(processIssue);
                return issuesWithExtractFields;
            };
            displayResults = function (err, allIssuesWithExtractFields) {
                linq.from(allIssuesWithExtractFields)
                    .forEach(function (issueWithExtractFields) {
                        console.log(JSON.stringify(issueWithExtractFields));
                    });
            };
            readFileAndIterate(program.args[0], processResults, displayResults);
        });
    }
);

