var a = 1;
var b = 2;
var c = 3;
var d = 4;

function f(n) {
    c = 10;
    return c;
}

alert(a - f(b) - c * d);