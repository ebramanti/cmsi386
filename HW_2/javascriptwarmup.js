//  CMSI 386 HW #2
//  + 1 for Zane & Ed for semicolons
//  http://jsfiddle.net/jadengore/5J84t/

//  Problem #1
function change(cents) {
    if (cents < 0 || cents % 1 !== 0) {
        throw new IllegalArgumentException(cents, " was passed, non-negative integer expected.");
    }
    var usCoinDenominations = [25, 10, 5, 1];
    var results = [];
    var remaining = cents;
    for (i = 0; i < usCoinDenominations.length; i++) {
        s = usCoinDenominations[i];
        partialSolution = [Math.floor(remaining / s), remaining % s];
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
    if (base === 1) {
        f(1);
    }
    else {
        for (var i = 1; i <= max; i *= base) {
            f(i);
        }
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

//custom IllegalArgumentException() method
function IllegalArgumentException(value, message) {
   this.value = value;
   this.message = message;
   this.toString = function() {
      return this.value + this.message
   };
   throw this;
}

//use this to test powersOfTwo
function collectPowersOfTwo(limit) {
    var result = [];
    powersOfTwo(limit, function(value) {result.push(value)});

    return result;
}

//use this to test powers
function collectPowers(base,limit) {
    var result = [];
    powers(base,limit, function(value) {result.push(value)});

    return result;
}