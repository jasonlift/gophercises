package btree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func NewTree(arr []int) *TreeNode {
	return levelBuilder(arr, 0)
}

func levelBuilder(arr []int, i int) *TreeNode {
	if i >= len(arr) {
		return nil
	}
	if arr[i] == -1 {
		return nil
	}
	return &TreeNode{
		Val: arr[i],
		Left: levelBuilder(arr, i*2+1),
		Right: levelBuilder(arr, i*2+2),
	}
}
