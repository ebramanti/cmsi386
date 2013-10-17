package gowarmup

import "testing"

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
