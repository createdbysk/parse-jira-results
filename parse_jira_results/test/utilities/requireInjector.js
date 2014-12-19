/**
 * Provides a requirejs injector based on Squire.
 */
'use strict';
var path,
    createInjector;

path = require('path');

createInjector = function () {
    'use strict';
    var requirejs,
        Squire;
    requirejs = require('requirejs');
    requirejs.config({
        baseUrl: path.resolve(__dirname, "../.."),
        nodeRequire: require,
        packages: [
            {
                name: "squirejs",
                location: "node_modules/squirejs",
                main: "src/Squire"
            }
        ]
    });
    Squire = requirejs('squirejs');
    return new Squire();
};

module.exports = {
    createInjector : createInjector
};
