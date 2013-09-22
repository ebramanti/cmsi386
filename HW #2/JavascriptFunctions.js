function change(cents) {
    //if cents < 0: raise ValueError("Not an int > 0.")
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

function strip_vowels(s) {
    return s.replace(/[aeiou]/ig, '');
}