var Some = function (value) { return value; };
var None = null;
var myval = Some(3);
var match = function (value) {
    return function (value) {
        return value;
    };
};
var x = match(myval)((function () {
    switch (myval) {
        case None:
            return 1;
        default:
            return myval;
    }
})());
// console.log(myval + 3);
//
console.log(x);
