package linked

import "testing"

func TestSortList(t *testing.T) {
	head := array2LinkedList([]int{4,2,1,3})
	res := sortList(head)
	t.Logf("get %v", linkedList2Array(res))
}
