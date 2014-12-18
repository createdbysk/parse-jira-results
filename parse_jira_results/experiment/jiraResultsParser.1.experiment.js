var requirejs;
requirejs = require('requirejs'); 

requirejs.config({
    baseUrl: __dirname,
    nodeRequire: require
});

requirejs(['fs'], function (fs) {
	
});