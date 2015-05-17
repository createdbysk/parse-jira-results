define(['linq'], function (linq) {
    'use strict';
    var transformer;

    /**
     *  input       - The input that each transform operates on
     *  transforms  - The object that contains key value pairs as follows
     *                {
     *                      transformName: transformFunction(input, callback(transformedResult));
     *                }
     */
    transformer = function (input, transforms) {
        var applyTransforms,
            transformsArray,
            transformName;
        // linq.from does not work if the value in the object is a function.
        // Therefore, transform the object into an array.
        transformsArray = [];
        // Iterate over all the properties of the object
        for(transformName in transforms) {
            if (transforms.hasOwnProperty(transformName)) {
                // add the pair to the array.
                transformsArray.push({
                    name: transformName,
                    fn: transforms[transformName]
                });
            }
        }
        // Define the function to transform the input to generate the result.
        applyTransforms =
            function (input) {
                var results;
                // Iterate over the transforms array, apply each transform to the input,
                // and store the result paired with the transform name in combination
                // object that represents the result.
                results =
                    linq.from(transformsArray)
                        .aggregate({},
                            function (combination, transform) {
                                transform.fn(input, function (result) {
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
