package linked

import "testing"

func TestReverseListMN(t *testing.T) {
	root := Array2LinkedList([]int{1,2,3,4,5,6,7})
	outList := reverseBetween(root, 1, 5)
	if outList != nil {
		t.Logf("Got %v", LinkedList2Array(outList))
	}
}
