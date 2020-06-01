package btree

import "testing"

func TestBuildTree(t *testing.T) {
	input := []int{3,9,20,-1,-1,15,7}
	tree := NewTree(input)
	print := levelOrder(tree)
	if tree != nil {
		t.Logf("level order print: %v", print)
	}
}
