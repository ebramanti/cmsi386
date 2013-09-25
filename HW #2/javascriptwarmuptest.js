$(function () {

    var anagram = function (s, t) {
        if (s.length !== t.length) {
            return false;
        }
        var a = s.split(""), b = t.split("");
        return $(a).not(b).length === 0 && $(b).not(a).length === 0
    };

    test("Change Tests", function () {
        deepEqual(change(0), [0, 0, 0, 0]);
        deepEqual(change(1), [0, 0, 0, 1]);
        deepEqual(change(5), [0, 0, 1, 0]);
        deepEqual(change(13), [0, 1, 0, 3]);
/*5*/   deepEqual(change(10), [0, 1, 0, 0]);
        deepEqual(change(25), [1, 0, 0, 0]);
        deepEqual(change(97), [3, 2, 0, 2]);
        deepEqual(change(100), [4, 0, 0, 0]);
        deepEqual(change(1000), [40, 0, 0, 0]);
/*A*/   deepEqual(change(41), [1, 1, 1, 1]);
        throws(function() {change(-10)}, IllegalArgumentException);
        throws(function() {change(-47)}, IllegalArgumentException);
        throws(function() {change(11.3)}, IllegalArgumentException);
        throws(function() {change(13/7)}, IllegalArgumentException);
/*F*/   deepEqual(change(9007199254740992), [360287970189639, 1, 1, 2]);
        // More needed
    });

    test("Strip Vowels Tests", function () {
        deepEqual(stripVowels("Hello, world"), "Hll, wrld");
        deepEqual(stripVowels("AEIOU"), "");
        deepEqual(stripVowels("aeiou"), "");
        deepEqual(stripVowels("Zane Quincy Kansil"), "Zn Qncy Knsl");
        deepEqual(stripVowels("Tom Marvolo Riddle"), "Tm Mrvl Rddl");
        deepEqual(stripVowels("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"), "BbCcDdFfGgHhJjKkLlMmNnPpQqRrSsTtVvWwXxYyZz");
        deepEqual(stripVowels("Ballin like I got Oprah Winfrey money~"), "Blln lk  gt prh Wnfry mny~");
        deepEqual(stripVowels("Sweet home Alabama"), "Swt hm lbm");
        deepEqual(stripVowels("Once upon a time there was a big huge giant Moon up there."), "nc pn  tm thr ws  bg hg gnt Mn p thr.");
        deepEqual(stripVowels("Toy Story 2"), "Ty Stry 2");
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
        expect(0);          //Means tests are needed but will not crash tester.
    });

    test("Powers of n Tests", function () {
        expect(0);         
    });

    test("Interleave Tests", function () {
        expect(0);
    });

    test("Stutter Tests", function () {
        expect(0);
    });

    test("Word Count Tests", function () {
        deepEqual(wordCount("I'm here fireman"), {"i'm":1, "here":1, "fireman":1});
        deepEqual(wordCount("na na na na-na na na NA BATMAN!"), {"na": 8, "batman":1});
        deepEqual(wordCount("On to the next on on to the next on... "), {"on":4, "to":2, "the":2, "next":2});
        deepEqual(wordCount(""), {});
        deepEqual(wordCount("I can'''t do that, I just can''t can't do that."), {"i":2, "can'''t":1, "do":2, "that":2, "just":1, "can''t":1, "can't":1});
        deepEqual(wordCount("Good\nMorning!"), {"good":1, "morning":1});
        deepEqual(wordCount("If you dog a dog, you'll\nbe DOG-TIRED."), {"if":1, "you":1, "dog":3, "a":1, "you'll":1, "be":1, "tired":1});
        //  Add more
    });
});







