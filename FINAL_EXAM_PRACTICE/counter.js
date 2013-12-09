var counter = function() {
    var n = 0;
    return {
        increment: function() {
            n = n + 1
            console.log(n);
        },
        decrement: function() {
            n = n - 1;
            console.log(n);
        }
    };
}();