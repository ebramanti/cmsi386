//	Problem #10
//	http://jsfiddle.net/jadengore/qfhLh/15/

/*	Used window.prefixes to get prefixes() noticed by HTML
 *  in jsfiddle.
 */
window.prefixes = function(word, i) {
    if (i <= word.length) {
        var substring = i === 0 ? "&nbsp;" : word.substring(0, i);
        $("body").append('<div>' + substring + '</div>');
        setTimeout(function () {
            prefixes(word, i + 1);
        }, 1000);
    }
}