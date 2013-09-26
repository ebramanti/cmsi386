function prefixes(word) {
        
    var next = word.substring(0, j);
        $("body").append("<div>" + next + "</div>");
        setTimeout(function () {
            j++;
            prefixes();
        }, 1000);
}

prefixes("Hello, world");