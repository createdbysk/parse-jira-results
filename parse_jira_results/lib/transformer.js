define(['linq'], function (linq) {
    'use strict';
    var transformer;

    /**
     *  input       - The input that each transform operates on
     *  transforms  - The object that contains key value pairs as follows
     *                {
     *                      transformName: transformFunction(input, callback(err, transformedResult));
     *                }
     */
    transformer = function (input, transforms) {
        var applyTransforms,
            transformsArray,
            transformName;
        // linq.from does not work if the value in the object is a function.
        // Therefore, transform the object into an array.
        transformsArray = [];
        for(transformName in transforms) {
            if (transforms.hasOwnProperty(transformName)) {
                transformsArray.push({
                    name: transformName,
                    fn: transforms[transformName]
                });
            }
        }
        applyTransforms =
            function (input) {
                var results;
                results =
                    linq.from(transformsArray)
                        .aggregate({},
                            function (combination, transform) {
                                transform.fn(input, function (err, result) {
                                    combination[transform.name] =
                                        (result instanceof linq) ? result.toArray() : result;
                                });
                                return combination;
                            }
                        );
                return results;
            };
        return applyTransforms(input);
    };

    return transformer;
});
