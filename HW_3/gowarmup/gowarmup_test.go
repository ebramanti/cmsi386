package gowarmup

import (
	"strings"
	"testing"
)

func TestDivmod(t *testing.T) {
	a := []struct {
		q, w, e, r int
	}{
		{25, 5, 5, 0},
		{30, 5, 6, 0},
		{13, 10, 1, 3},
		{31, 15, 2, 1},
		{7, 2, 3, 1},
	}

	for _, sl := range a {
		bat, cat := Divmod(sl.q, sl.w)
		if bat != sl.e || cat != sl.r {
			t.Errorf("Expected (%d, %d), for Divmod(%d,%d)", sl.q, sl.w, sl.e, sl.r)
		}
	}
}

func TestChange(t *testing.T) {
	a := []struct {
		input, r0, r1, r2, r3 int
	}{
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{5, 0, 0, 1, 0},
		{8, 0, 0, 1, 3},
		{10, 0, 1, 0, 0},
		{25, 1, 0, 0, 0},
		{97, 3, 2, 0, 2},
		{100, 4, 0, 0, 0},
		{1000, 40, 0, 0, 0},
		{41, 1, 1, 1, 1},
	}

	for i, sl := range a {
		aa, bb, cc, dd := Change(sl.input)
		if aa != sl.r0 || bb != sl.r1 || cc != sl.r2 || dd != sl.r3 {
			t.Errorf("Bad change given for test %d: %v", i, sl)
		}
	}
}

func TestRemoveVowels(t *testing.T) {
	a := []struct {
		input, output string
	}{
		{"Hello, world", "Hll, wrld"},
		{"AEIOU", ""},
		{"aeiou", ""},
		{"Zane Quincy Kansil", "Zn Qncy Knsl"},
		{"Tom Marvolo Riddle", "Tm Mrvl Rddl"},
		{"AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz", "BbCcDdFfGgHhJjKkLlMmNnPpQqRrSsTtVvWwXxYyZz"},
		{"Ballin like I got Oprah Winfrey money~", "Blln lk  gt prh Wnfry mny~"},
		{"Sweet home Alabama", "Swt hm lbm"},
		{"Once upon a time there was a big huge giant Moon up there.", "nc pn  tm thr ws  bg hg gnt Mn p thr."},
		{"Toy Story 2", "Ty Stry 2"},
		{"3.14159", "3.14159"},
		//{"Hello", "Hello"}, //ensuring failure possible
	}

	for i, sl := range a {
		if RemoveVowels(sl.input) != sl.output {
			t.Errorf("Bad vowel detection for test %d", i)
		}
	}
}

func sameCharSet(s, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		for i := range s {
			if !strings.Contains(t, string(s[i])) {
				return false
			}
		}
	}
	return true
}

/*
 *  Limited to testing Scramble for unicode-8 characters only.
 */
func TestScramble(t *testing.T) {
	a := []struct {
		s string
	}{
		{""},
		{"a"},
		{"rat"},
		{"zzz"},
		//{"^*&^*&^▱ÄÈËɡɳɷ"},
		{"a"},
		{"rat"},
		{"BSOD"},
		{"BDFL"},
		{"Python testing"},
		{" a b c "},
		{"Hannah Montana"},
		{"142452729"},
	}

	for i, sl := range a {
		if !sameCharSet(sl.s, Scramble(sl.s)) || len(sl.s) != len(Scramble(sl.s)) {
			t.Errorf("Bad scramble, impossible scramble executed at test: %d", i)
		}
	}

}

func TestPowersOfTwo(t *testing.T) {
	a := []struct {
		x int
		y []int
	}{
		{0, []int{}},
		{1, []int{1}},
		{7, []int{1, 2, 4}},
		{8, []int{1, 2, 4, 8}},
		{15, []int{1, 2, 4, 8}},
		{16, []int{1, 2, 4, 8, 16}},
		{17, []int{1, 2, 4, 8, 16}},
		{31, []int{1, 2, 4, 8, 16}},
		{32, []int{1, 2, 4, 8, 16, 32}},
		{33, []int{1, 2, 4, 8, 16, 32}},
		{70, []int{1, 2, 4, 8, 16, 32, 64}},
		{130, []int{1, 2, 4, 8, 16, 32, 64, 128}},
		{255, []int{1, 2, 4, 8, 16, 32, 64, 128}},
		{256, []int{1, 2, 4, 8, 16, 32, 64, 128, 256}},
		{65536, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536}},
		{-22, []int{}},
		//{8, []int{1, 2, 4}},            // Really, really, really blows up for these.
		//{8, []int{1, 2, 4, 7}},         // used Fatalf
		//{8, []int{1, 2, 4, 8, 10}},
	}

	for i, sl := range a {
		loud := func(upTo int) []int {
			bat := []int{}
			toSlice := func(i int) {
				bat = append(bat, i)
			}
			PowersOfTwo(upTo, toSlice)
			return bat
		}
		tiger := loud(sl.x)
		if len(tiger) != len(sl.y) {
			t.Fatalf("Bad test result, compared strings of different length at test: %d", i)
		}
		for o := 0; o < len(tiger); o++ {
			if tiger[o] != sl.y[o] {
				t.Errorf("Bad power outputted during test %d... %d != %d", i, tiger[o], sl.y[o])
			}
		}
	}
}

func TestPowers(t *testing.T) {
	a := []struct {
		x, y int
		z    []int
	}{
		{0, 0, []int{}},
		{1, 6, []int{1}},
		{1, 20, []int{1}},
		{27, 0, []int{}},
		{14, -1, []int{}},
		{-3, 30, []int{1, -3, 9, -27}},
		{-3, 81, []int{1, -3, 9, -27, 81, -243}},
		{3, 30, []int{1, 3, 9, 27}},
		{3, 27, []int{1, 3, 9, 27}},
		{3, 26, []int{1, 3, 9}},
		{4, 64, []int{1, 4, 16, 64}},
		{10, 999, []int{1, 10, 100}},
		{10, 1000, []int{1, 10, 100, 1000}},
		{5, 125, []int{1, 5, 25, 125}},
		{5, 126, []int{1, 5, 25, 125}},
		{2, 130, []int{1, 2, 4, 8, 16, 32, 64, 128}},
		{3, 1, []int{1}},
		{3, 3, []int{1, 3}},
		{3, 10, []int{1, 3, 9}},
		{2, 65536, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536}},
		//{3, 10, []int{1, 3, 8}},            //failure tests
		//{3, 10, []int{1, 3, 9, 15}},
	}
	for i, sl := range a {
		loud := func(base, upTo int) []int {
			bat := []int{}
			toSlice := func(i int) {
				bat = append(bat, i)
			}
			Powers(base, upTo, toSlice)
			return bat
		}
		tiger := loud(sl.x, sl.y)
		if len(tiger) != len(sl.z) {
			t.Fatalf("Bad test result, compared strings of different length at test: %d", i)
		}
		for o := 0; o < len(tiger); o++ {
			if tiger[o] != sl.z[o] {
				t.Errorf("Bad power outputted during test %d... %d != %d", i, tiger[o], sl.z[o])
			}
		}
	}
}

func TestInterleave(t *testing.T) {
	a := []struct {
		x, y, z []string
	}{
		{[]string{"1", "2"}, []string{"3", "4"}, []string{"1", "3", "2", "4"}},
		{[]string{"1", "2"}, []string{"3", "4", "5"}, []string{"1", "3", "2", "4", "5"}},
		{[]string{"1", "2"}, []string{"3", "4", "5", "6"}, []string{"1", "3", "2", "4", "5", "6"}},
		{[]string{"a", "b", "c"}, []string{"3", "4"}, []string{"a", "3", "b", "4", "c"}},
		{[]string{"1"}, []string{"3"}, []string{"1", "3"}},
		{[]string{"1", "2"}, []string{}, []string{"1", "2"}},
		{[]string{"Aqua", "Start up"}, []string{"1", "2"}, []string{"Aqua", "1", "Start up", "2"}},
		{[]string{"1", "2"}, []string{"Aqua", "Start up"}, []string{"1", "Aqua", "2", "Start up"}},
		{[]string{"1"}, []string{"2", "3", "4", "5", "6", "7", "8", "9"}, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}},
		{[]string{}, []string{}, []string{}},
		//{[]string{"1"}, []string{"3"}, []string{"1", "2"}}, //failure test
	}
	for i, sl := range a {
		gira := Interleave(sl.x, sl.y)
		rhin := sl.z
		for byteIndex, codePoint := range gira {
			if rhin[byteIndex] != codePoint {
				t.Errorf("Bad interleave, out-of-sequence value in test: %d", i)
			}
		}

	}
}

func TestStutter(t *testing.T) {
	a := []struct {
		a, b []string
	}{
		{[]string{}, []string{}},
		{[]string{"a", "b", "c"}, []string{"a", "a", "b", "b", "c", "c"}},
		{[]string{"a", "a", "a"}, []string{"a", "a", "a", "a", "a", "a"}},
		{[]string{"1", "2", "3"}, []string{"1", "1", "2", "2", "3", "3"}},
		{[]string{""}, []string{"", ""}},
		{[]string{"Repeat that"}, []string{"Repeat that", "Repeat that"}},
		{[]string{"777", "qwert"}, []string{"777", "777", "qwert", "qwert"}},
		//{[]string{"777", "qwert"}, []string{"777", "777", "qwert", "+++++"}}, //failure test
	}
	for i, sl := range a {
		gira := Stutter(sl.a)
		rhin := sl.b
		for byteIndex, codePoint := range gira {
			if rhin[byteIndex] != codePoint {
				t.Errorf("Bad interleave, out-of-sequence value in test: %d", i)
			}
		}
	}
}
