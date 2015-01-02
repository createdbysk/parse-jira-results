var program,
    extractors,
    path,
    requirejs,
    util,
    linq;
program = require('commander');
path = require('path');
requirejs = require('library/configuredRequirejs');
util = require('util');
linq = require('linq');
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
    .option('-e --extractor [value]', 'requirejs modules that extract values', extractors, {modules:[], names:[]})
    .parse(process.argv);

console.log(program.extractor);
requirejs(program.extractor.modules, function () {
    var tasks;
    tasks = linq.from(program.extractor.names)
        .zip(linq.from(arguments), 
            function (value1, value2) {
                return [value1, value2];
            }
        )
        .aggregate({}, function (combination, value) {
            combination[value[0]] = value[1];
            return combination;
        });
    console.log(tasks);
});