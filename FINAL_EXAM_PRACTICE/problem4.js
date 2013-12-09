var muPuzzle = function() {
    var s = "MI";
    return {
        value: function() {
            return s;
        },
        add_U: function() {
            if (s.slice(-1) == "I") s = s.concat("U"); // or s += "U";
            return s;
        },

        appender: function() {
            s = s + s.slice(1);
            return s;
        },

        replace_III: function() {
            s = s.replace(/III/, "U");
            return s;
        },

        remove_UU: function() {
            s = s.replace(/UU/, '');
            return s;
        }
    };
}();

var a = muPuzzle;
console.log(a.value());
console.log(a.appender());
console.log(a.appender());
//console.log(a.appender());
console.log(a.replace_III());
console.log(a.add_U());
console.log(a.add_U());
console.log(a.add_U());
console.log(a.remove_UU());