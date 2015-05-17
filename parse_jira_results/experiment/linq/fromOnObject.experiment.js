var obj,
    linq;
linq = require('linq');
obj = {
    two: "two",
    one: "one"
};

linq.from(obj)
    .forEach(function (keyValuePair, index) {
        console.log(keyValuePair, index);
    });

linq.from(obj)
    .aggregate({}, function (combined, keyValuePair) {
        combined[keyValuePair.key] = keyValuePair.value;
        console.log(combined, keyValuePair);
        return combined;
    });
