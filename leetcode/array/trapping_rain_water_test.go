package array

import "testing"

func TestGeneral(t *testing.T) {
	pool := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	r := trap(pool)
	if r != 6 {
		t.Errorf("got %d, want %d", r, 6)
	} else {
		t.Logf("TestGeneral pass")
	}
}
