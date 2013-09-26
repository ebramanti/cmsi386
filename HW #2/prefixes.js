<<<<<<< HEAD
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

=======
function prefixes(word, i) {
    if (i <= word.length) {
        var substring = i === 0 ? "&nbsp;" : word.substring(0, i);
        $("body").append('<div>' + substring + '</div>');
        setTimeout(function () {
            prefixes(word, i + 1);
        }, 1000);
    }
}
>>>>>>> a4f9e24fa140602487537c43bf3d317b7cbea66e
