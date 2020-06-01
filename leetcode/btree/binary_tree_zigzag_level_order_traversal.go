package btree

/**
103
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	stack1 := make([]*TreeNode, 0)
	stack2 := make([]*TreeNode, 0)
	isReversed := false
	stack1 = append(stack1, root)

	for len(stack1) > 0 || len(stack2) > 0 {
		if !isReversed { // forward order
			isReversed = true
			tmp := []int{}
			for len(stack1) > 0 {
				node := stack1[len(stack1)-1]
				tmp = append(tmp, node.Val)
				stack1 = stack1[:len(stack1)-1]
				if node.Left != nil {
					stack2 = append(stack2, node.Left)
				}
				if node.Right != nil {
					stack2 = append(stack2, node.Right)
				}
			}
			if len(tmp) > 0 {
				res = append(res, tmp)
			}
		} else { // backward order
			isReversed = false
			tmp := []int{}
			for len(stack2) > 0 {
				node := stack2[len(stack2)-1]
				tmp = append(tmp, node.Val)
				stack2 = stack2[:len(stack2)-1]
				if node.Right != nil {
					stack1 = append(stack1, node.Right)
				}
				if node.Left != nil {
					stack1 = append(stack1, node.Left)
				}
			}
			if len(tmp) > 0 {
				res = append(res, tmp)
			}
		}
	}
	return res
}