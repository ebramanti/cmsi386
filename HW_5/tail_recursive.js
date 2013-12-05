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