$(function () {

    var anagram = function (s, t) {
        if (s.length !== t.length) {
            return false;
        }
        var a = s.split(""), b = t.split("");
        return $(a).not(b).length === 0 && $(b).not(a).length === 0
    };

    test("Change Tests", function () {
        deepEqual(change(97), [3, 2, 0, 2]);
        deepEqual(change(100), [4, 0, 0, 0]);
        deepEqual(change(13), [0, 1, 0, 3]);
        deepEqual(change(0), [0, 0, 0, 0]);
        deepEqual(change(1000), [250, 0, 0, 0]);
        deepEqual(change(97), [3, 2, 0, 1]);
        deepEqual(change(41), [1, 1, 1, 1]);
        deepEqual(change(1), [0, 0, 0, 1]);
        deepEqual(change(5), [0, 0, 1, 0]);
        deepEqual(change(10), [0, 1, 0, 0]);
        deepEqual(change(25), [1, 0, 0, 0]);
        throws(change(-10), TypeError, "Negative integer argument found");
        throws(change(-47), TypeError, "Negative integer argument found");
        throws(change(11.3), TypeError, "Negative integer argument found");
        throws(change((13/7)), TypeError, "Negative integer argument found");
        deepEqual(change(9007199254740992), [360287970189639, 1, 1, 2]);
        // More needed
    });

    test("Strip Vowels Tests", function () {
        deepEqual(stripVowels("Hello, world"), "Hll, wrld");
        deepEqual(strip_vowels("AEIOU"), "");
        deepEqual(strip_vowels("aeiou"), "");
        deepEqual(strip_vowels("Zane Quincy Kansil"), "Zn Qncy Knsl");
        deepEqual(strip_vowels("Tom Marvolo Riddle"), "Tm Mrvl Rddl");
        deepEqual(strip_vowels("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"), "BbCcDdFfGgHhJjKkLlMmNnPpQqRrSsTtVvWwXxYyZz");
        deepEqual(strip_vowels("Ballin like I got Oprah Winfrey money~"), "Blln lk  gt prh Wnfry mny~");
        deepEqual(strip_vowels("Sweet home Alabama"), "Swt hm lbm");
        deepEqual(strip_vowels("Once upon a time there was a big huge giant Moon up there."), "nc pn  tm thr ws  bg hg gnt Mn p thr.");
        deepEqual(strip_vowels("Toy Story 2"), "Ty Stry 2");
        deepEqual(stripVowels("3.14159"), "3.14159");  
        // More needed
    });

    test("Scramble Tests", function () {
        var data = ["", "a", "rat", "zzz", "^*&^*&^▱ÄÈËɡɳɷ"]
        data.forEach(function (s) {
            ok(anagram(s, scramble(s)));
        });
    });

    test("Powers of Two Tests", function () {
        // Implement these
    });

    // Test powers here

    // Test interleave here

    // Test stutter here

    test("Word Count Tests", function () {
        deepEqual(wordCount)

        //  Add more
    });
});