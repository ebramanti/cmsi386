package gowarmup

import (
	"errors"
	"strconv"
	"testing"
)

/*
func TestAverage(t *testing.T) {
    average, error := Average([]float64{10.0, 1.0, 4.0})
    if average != 5.0 && error != nil {
        t.Errorf("Average of [10, 1, 4] should be 5")
    }
    _, error = Average([]float64{})
    if error == nil {
        t.Errorf("Average of empty slice should return an error")
    }
}
*/

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
		input          int
		supposedResult []int
	}{
		{0, []int{0, 0, 0, 0}},
		{1, []int{0, 0, 0, 1}},
		{5, []int{0, 0, 1, 0}},
		{8, []int{0, 1, 0, 3}},
		{10, []int{0, 1, 0, 0}},
		{25, []int{1, 0, 0, 0}},
		{97, []int{3, 2, 0, 2}},
		{100, []int{4, 0, 0, 0}},
		{1000, []int{40, 0, 0, 0}},
		{41, []int{1, 1, 1, 1}},
	}

	for _, sl := range a {
		if !bytes.Equal(Change(sl.input), sl.supposedResult) {
			t.Errorf("Expected %v for Change(%d)", sl.supposedResult, sl.input)
		}
	}
}

func TestRemoveVowels(t *testing.T) {
	t.Errorf("TestRemove not implemented")
}

func TestScramble(t *testing.T) {
	t.Errorf("TestChange not implemented")
}

func TestPowersOfTwo(t *testing.T) {
	t.Errorf("TestChange not implemented")
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

func sliceCompare(x, y []int) (bool, error) {
	if x.length != y.length {
		c := [x.length]int{x}
		d := [y.length]int{y}
	} else {
		return false, errors.New("Cannot compare slices of variant length.")
	}
}

//space
