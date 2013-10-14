package gowarmup

import "testing"

func TestAverage(t *testing.T) {
	//cut from toals site..
    average, error := Average([]float64{10.0, 1.0, 4.0})
    if average != 5.0 && error != nil {
        t.Errorf("Average of [10, 1, 4] should be 5")
    }
    _, error = Average([]float64{})
    if error == nil {
        t.Errorf("Average of empty slice should return an error")
    }
}