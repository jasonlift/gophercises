package array

import "testing"

func TestGeneralExample(t *testing.T) {
	arr := []int{1,8,6,2,5,4,8,3,7}
	maxArea := maxArea(arr)
	if maxArea != 49 {
		t.Errorf("got %d, want %d", maxArea, 49)
	} else {
		t.Logf("TestGeneralExample pass")
	}
}

func TestUnilateral(t *testing.T) {
	arr := []int{1,2,7}
	maxArea := maxArea(arr)
	if maxArea != 2 {
		t.Errorf("got %d, want %d", maxArea, 2)
	} else {
		t.Logf("TestUnilateral pass")
	}
}