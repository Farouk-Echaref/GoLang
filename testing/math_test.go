package math

import (
	"testing"
)

// testing the Abs function individually
func TestAbs(t *testing.T){
	//there's different log functions in t
	// test abs for 3 different values: -1, 0, 1
	if Abs(-1) != 1 {
		t.Errorf("abs(-1) = %d; want 1", Abs(-1))
	}
	if Abs(0) != 0 {
		t.Errorf("abs(-1) = %d; want 1", Abs(-1))
	}
	if Abs(1) != 1 {
		t.Errorf("abs(-1) = %d; want 1", Abs(-1))
	}
}

