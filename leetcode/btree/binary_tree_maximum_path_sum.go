package btree

import "math"

/**
124
https://leetcode.com/problems/binary-tree-maximum-path-sum/
 */

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	path(root, &max)
	return max
}

func path(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}
	left := max(path(root.Left, m), 0)
	right := max(path(root.Right, m), 0)
	*m = max(*m, left+right+root.Val)
	return max(left+root.Val, right+root.Val)
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}