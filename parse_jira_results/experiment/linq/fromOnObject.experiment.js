var obj,
    linq,
    i,
    f;
linq = require('linq');
f = function () {};
obj = {
    two: "two",
    one: f
};

i = 0;

linq.from(obj)
    .forEach(function (keyValuePair, index) {
        console.log(++i, arguments, keyValuePair, index);
    });