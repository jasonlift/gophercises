package array

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	arr := []int {5,9,1,4,7,2}
	answer := []int {1,2,4,5,7,9}
	quicksort(arr, 0, len(arr)-1)
	if !reflect.DeepEqual(arr, answer) {
		t.Errorf("wrong answer %v", arr)
	}
}
