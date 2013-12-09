//  OUR WAY
var a = [];
for (var i = 0; i < 10; i++) {
    var functionCreator = (function() {
        var index = i;
        a[index] = function(x) {return (x / index * index);}
    }());
}


//  TOAL'S WAY
var a = [];
for (var i = 0; i < 10; i++) {
    a[i] = (function(i) {
        return function(x) {return x / (i*i);}
    }(i));
}