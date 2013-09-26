window.prefixes = function (word) {
    var j = 0;
    function next() {
        if (j <= word.length) {
            $("body").append("<div>" + word.substring(0, j) + "</div>");
            setTimeout(function () {next();}, 1000);
            j++;
        }
    }  
    next();
}
/*	Had to use window.prefixes to get prefixes noticed by the html
	part of jsfiddle.
*/

