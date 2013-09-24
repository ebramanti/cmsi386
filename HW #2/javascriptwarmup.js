function change(cents) {
    if (cents < 0) {
        throw "Not an integer greater than 0.";
    }
    //if cents isBoolean: raise ValueError("Input unrounded number of cents (float found)")
    var usCoinDenominations = [25, 10, 5, 1];
    var results = [];
    var remaining = cents;
    for (i = 0; i < usCoinDenominations.length; i++) {
        s = usCoinDenominations[i];
        partialSolution = [Math.floor(remaining / s), remaining % s];  //ask toal
        results.push(partialSolution[0]);
        remaining = partialSolution[1];
    }
    return results;
}

function stripVowels(s) {
    return s.replace(/[aeiou]/ig, '');
}

function scramble(s) {
    return s.split("").sort(function(){return Math.random() - 0.5;}).join("");
}

function powersOfTwo(max) {
    return -1;
}

function powers(base, max) {
    return -1;
}

function interleave(array1, array2) {
    var result = [];
    var i = 0;
    var j = 0;
    while (result.length < array1.length + array2.length) {
        if (i < array1.length) {
            result.push(array1[i]);
        }
        if (j < array2.length) {
            result.push(array2[j]);
        }
        i++, j++;
    }
    return '[' + result + ']'; //check to see if brackets are necessary
}

function stutter(arrayA) {
    return interleave(arrayA, arrayA);
} 

function wordCount(s) {
    return -1;
}