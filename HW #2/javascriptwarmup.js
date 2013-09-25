//  Problem #1
function change(cents) {
    if (cents < 0) {
        //throws "Negative integer argument found.";
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

//  Problem #2
function stripVowels(s) {
    return s.replace(/[aeiou]/ig, '');
}

//  Problem #3
function scramble(s) {
    return s.split("").sort(function(){return Math.random() - 0.5;}).join("");
}

//  Problem #4
function powersOfTwo(max, f) {
    for (var i = 1; i <= max; i *= 2) {
        f(i);
    }
}

//  Problem #5
function powers(base, max, f) {
    for (var i = 1; i <= max; i *= base) {
        f(i);
    }
}

//  Problem #6
function interleave(a, b) {
    var result = [];
    var i = 0;
    var j = 0;
    while (result.length < a.length + b.length) {
        if (i < a.length) {
            result.push(a[i]);
            i++;
        }
        if (j < b.length) {
            result.push(b[j]);
            j++;
        }
    }
    return result;
}

//  Problem #7
function stutter(array) {
    return interleave(array, array);
} 

//  Problem #8
function wordCount(s) {
    var result = {};
    var wordArray = s.toLowerCase().replace("-", " ").split(/\s/);
    var check = /^[?'.!,:;a-zA-Z0-9._\-]*$/;
    for(var i = 0; i < wordArray.length; i++) {
        if (check.test(wordArray[i]) && wordArray[i] !== "") {
            var word = wordArray[i];
            word = word.replace(/[?.!,:;]/ig, '');
            if (result.hasOwnProperty(word)) {
                result[word] += 1;
            }
            else {
                result[word] = 1;
            }
        }
    }
    return result;
}













// + 1 for Zane & Ed for semicolons