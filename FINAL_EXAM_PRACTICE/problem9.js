function withoutEvens(a) {
    function helper(i, result) {
        return (i < a.length) ? helper(i + 2, result.concat(a[i])) : result;
    }
    return helper(0, []);
}

console.log(withoutEvens([0,1,2,3,4,5,6]));
console.log(withoutEvens([4,6,true,null,2.3,"xyz"]))