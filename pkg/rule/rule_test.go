package rule

import (
	"testing"
)

func containsExact(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func RunParse(t *testing.T, in string, born, survive []int) {
	res, err := ParseRule(in)
	if err != nil {
		t.Error(err)
	}
	if !containsExact(res.Born, born) {
		t.Errorf("Born is wrong")
	}
	if !containsExact(res.Survive, survive) {
		t.Errorf("Survive is wrong got %v", res.Survive)
	}
}
func TestParse(t *testing.T) {
	for _, test := range []struct {
		s string
		a []int
		b []int
	}{{
		"B3/S23",
		[]int{3},
		[]int{2, 3},
	}, {
		"B2/S",
		[]int{2},
		[]int{},
	}, {
		"B36/S23",
		[]int{3, 6},
		[]int{2, 3},
	}} {
		RunParse(t, test.s, test.a, test.b)
	}
}
