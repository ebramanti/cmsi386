//  Problem 9: javascriptwarmuptest.js
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
    });

    test("Scramble Tests", function () {
        var data = [    "", "a", "rat", "zzz", "^*&^*&^▱ÄÈËɡɳɷ","a", "rat", 
                        "BSOD", "BDFL", "Python testing", " a b c ", 
                        "Hannah Montana", "142452729"
                    ]
        data.forEach(function (s) {
            ok(anagram(s, scramble(s)));
        });
    });

    test("Powers of Two Tests", function () {
        deepEqual(collectPowersOfTwo(0), []);
        deepEqual(collectPowersOfTwo(1), [1,]);
        deepEqual(collectPowersOfTwo(7), [1,2,4]);
        deepEqual(collectPowersOfTwo(8), [1,2,4,8]);
        deepEqual(collectPowersOfTwo(15), [1,2,4,8]);
        deepEqual(collectPowersOfTwo(16), [1,2,4,8,16]);
        deepEqual(collectPowersOfTwo(17), [1,2,4,8,16]);
        deepEqual(collectPowersOfTwo(31), [1,2,4,8,16]);
        deepEqual(collectPowersOfTwo(32), [1,2,4,8,16,32]);
        deepEqual(collectPowersOfTwo(33), [1,2,4,8,16,32]);
        deepEqual(collectPowersOfTwo(70), [1,2,4,8,16,32,64]);
        deepEqual(collectPowersOfTwo(130), [1,2,4,8,16,32,64,128]);
        deepEqual(collectPowersOfTwo(255), [1,2,4,8,16,32,64,128]);
        deepEqual(collectPowersOfTwo(256), [1,2,4,8,16,32,64,128,256]);
        deepEqual(collectPowersOfTwo(65536), [1,2,4,8,16,32,64,128,256,512,1024,2048,4096,8192,16384,32768,65536]);
        deepEqual(collectPowersOfTwo(-22), []);
    });

    test("Powers of n Tests", function () {
        deepEqual(collectPowers(0,0), []);
        deepEqual(collectPowers(1,6), [1]);
        deepEqual(collectPowers(1,20), [1]);
        deepEqual(collectPowers(27,0), []);
        deepEqual(collectPowers(14,-1), []);
        deepEqual(collectPowers(-3,30), [1,-3,9,-27], "Negative base");
        deepEqual(collectPowers(-3,81), [1,-3,9,-27,81,-243], "Negative magnitude larger than max magnitude");
        deepEqual(collectPowers(3,30), [1,3,9,27]);
        deepEqual(collectPowers(3,27), [1,3,9,27]);
        deepEqual(collectPowers(3,26), [1,3,9]);
        deepEqual(collectPowers(4,64), [1,4,16,64]);
        deepEqual(collectPowers(10,999), [1,10,100]);
        deepEqual(collectPowers(10,1000), [1,10,100,1000]);
        deepEqual(collectPowers(5,125), [1,5,25,125]);
        deepEqual(collectPowers(5,126), [1,5,25,125]);
        deepEqual(collectPowers(2,130), [1,2,4,8,16,32,64,128]);
        deepEqual(collectPowers(3,1), [1,]);
        deepEqual(collectPowers(3,3), [1,3]);
        deepEqual(collectPowers(3,10), [1,3,9]);
        deepEqual(collectPowers(2,65536), [1,2,4,8,16,32,64,128,256,512,1024,2048,4096,8192,16384,32768,65536]);         
    });

    test("Interleave Tests", function () {
        deepEqual(interleave([1,2],[3,4]), [1,3,2,4]);
        deepEqual(interleave([1,2],[3,4,5]), [1,3,2,4,5]);
        deepEqual(interleave([1,2],[3,4,5,6]), [1,3,2,4,5,6]);
        deepEqual(interleave(['a','b','c'],[3,4]), ['a',3,'b',4,'c']);
        deepEqual(interleave([1],[3]), [1,3]);
        deepEqual(interleave([1,2],[]), [1,2]);
        deepEqual(interleave(['Aqua','Start up'],[1,2]), ['Aqua',1,'Start up',2]);
        deepEqual(interleave([1,2],['Aqua','Start up']), [1,'Aqua',2,'Start up']);
        deepEqual(interleave([1],[2,3,4,5,6,7,8,9]), [1,2,3,4,5,6,7,8,9]);
        deepEqual(interleave([],[]), []);
    });

    test("Stutter Tests", function () {
        deepEqual(stutter([]), []);
        deepEqual(stutter(['a','b','c']), ['a','a','b','b','c','c']);
        deepEqual(stutter(['a','a','a']), ['a','a','a','a','a','a']);
        deepEqual(stutter([1,2,3]), [1,1,2,2,3,3]);
        deepEqual(stutter([[1],[2,3]]), [[1],[1],[2,3],[2,3]], "Array[]'s passed");
        deepEqual(stutter([[1,2,3]]), [[1,2,3],[1,2,3]], "Array[]'s passed");
        deepEqual(stutter([[1,2,[3]]]), [[1,2,[3]],[1,2,[3]]], "Array[]'s passed");
        deepEqual(stutter(['Repeat that']), ['Repeat that', 'Repeat that']);
        var baconObject = {isTasty: true};
        deepEqual(stutter([baconObject]), [baconObject, baconObject], "Objects passed");
    });

    test("Word Count Tests", function () {
        deepEqual(wordCount("I'm here fireman"), {"i'm":1, "here":1, "fireman":1});
        deepEqual(wordCount("na na na na-na na na NA BATMAN!"), {"na": 8, "batman":1});
        deepEqual(wordCount("On to the next on on to the next on... "), {"on":4, "to":2, "the":2, "next":2});
        deepEqual(wordCount(""), {});
        deepEqual(wordCount("Boss."), {"boss": 1});
        deepEqual(wordCount("I can'''t do that, I just can''t can't do that."), {"i":2, "can'''t":1, "do":2, "that":2, "just":1, "can''t":1, "can't":1});
        deepEqual(wordCount("Good\nMorning!"), {"good":1, "morning":1});
        deepEqual(wordCount("If you dog a dog, you'll\nbe DOG-TIRED."), {"if":1, "you":1, "dog":3, "a":1, "you'll":1, "be":1, "tired":1});
    });
});







