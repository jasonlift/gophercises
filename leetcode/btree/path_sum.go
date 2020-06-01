package btree

/**
112
https://leetcode.com/problems/path-sum/
 */

// preorder traverse
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && sum == root.Val {
		return true
	}

	sum -= root.Val
	l := hasPathSum(root.Left, sum)
	r := hasPathSum(root.Right, sum)

	return l || r
}