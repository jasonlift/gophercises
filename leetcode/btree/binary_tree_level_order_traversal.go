package btree

/**
102
https://leetcode.com/problems/binary-tree-level-order-traversal/
 */

// iterative BFS solution
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	tmp := []int{}
	stub := len(queue)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		tmp = append(tmp, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		stub--
		if stub == 0 {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			tmp = []int{}
			stub = len(queue)
		}
	}
	return res
}

func levelOrderRecursive(root *TreeNode) [][]int {
	result := [][]int{}
	levelOrderHelper(root, 0, &result)
	return result
}

func levelOrderHelper(node *TreeNode, level int, result *[][]int) {
	if node == nil { return }
	if level == len(*result) {
		// only when result is not ready for level x
		// append empty slice one time
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], node.Val)
	levelOrderHelper(node.Left, level + 1, result)
	levelOrderHelper(node.Right, level + 1, result)
}
