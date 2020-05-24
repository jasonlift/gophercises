package linked

import "testing"


func TestArr2LinkedList(t *testing.T) {
	arr1 := []int{1,2,3}
	arr2 := []int{}
	arr3 := []int{1,3,5,8}

	if !verifyCorrect(array2LinkedList(arr1), arr1) {
		t.Errorf("%v wrong", arr1)
	}
	if !verifyCorrect(array2LinkedList(arr2), arr2) {
		t.Errorf("%v wrong", arr2)
	}
	if !verifyCorrect(array2LinkedList(arr3), arr3) {
		t.Errorf("%v wrong", arr3)
	}
}
