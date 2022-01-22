package sum

import "testing"

func TestSum(t *testing.T) {

	r := Sum([]int{1, 2, 3, 4}...)
	if r != 10 {
		t.Error("Expected 2 but got ", r)
	}

	r = Sum(0, 0)
	if r != 0 {
		t.Error("Expected 0 but got ", r)
	}

	r = Sum(-1, -1)
	if r != -2 {
		t.Error("Expected -2 but got ", r)
	}

	r = Sum(0)
	if r != 0 {
		t.Error("Expected 0 but got ", r)
	}

	r = Sum()
	if r != 0 {
		t.Error("Expected 0 but got ", r)
	}
}
