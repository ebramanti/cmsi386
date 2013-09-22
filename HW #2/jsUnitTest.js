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
        deepEqual(change(8), [0, 0, 1, 3]);
        // More needed
    });

    test("Strip Vowels Tests", function () {
        deepEqual(stripVowels("Hello, world"), "Hll, wrld");
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

    // Test word count here
});