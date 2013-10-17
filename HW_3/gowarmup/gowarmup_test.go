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

func echo(i int) {
	return i
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
	}

	for i, sl := range a {
		PowersOfTwo(sl.x, echo(i))
	}
}

func TestPowers(t *testing.T) {
	t.Errorf("TestChange not implemented")
}

func TestInterleave(t *testing.T) {
	t.Errorf("TestChange not implemented")
}

func TestStutter(t *testing.T) {
	t.Errorf("TestChange not implemented")
}

//space
