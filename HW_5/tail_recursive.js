var a = [1,2,3];
var b = [21,7,12,2];
var c = [1, 12, -1, 0];
var d = [1, 20, 500, 6, -500, 40, -500, 67];

function findMinValue(a, i, sofar) {
    if (i === undefined) i = 0;
    if (i === a.length - 1) {
        return (a[i] < sofar) ? a[i] : sofar;
    }
    else if (sofar === undefined) {
        sofar = a[i]
    }
    else if (a[i] < sofar) {
        sofar = a[i]
    }
    
    return findMinValue(a, i + 1, sofar)
}

alert(findMinValue(a));
alert(findMinValue(b));
alert(findMinValue(c));
alert(findMinValue(d));