package btree

/**
113
https://leetcode.com/problems/path-sum-ii/
 */
func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	pathHelper(root, &res, []int{}, sum)
	return res
}

func pathHelper(root *TreeNode, res *[][]int, p []int, sum int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		p = append(p, root.Val)
		tmp := make([]int, len(p))
		copy(tmp, p)
		*res = append(*res, tmp)
		return
	}

	sum -= root.Val
	p = append(p, root.Val)
	pathHelper(root.Left, res, p, sum)
	pathHelper(root.Right, res, p, sum)
	p = p[:len(p)-1]
}