package btree

/**
236
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
 */

/**
recursive version
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}

/**
non-recursive version
use slice as a stack
steps:
1. find paths for root->p and root->q
2. traverse both paths and find the first non-equal node
	which is the bifurcation
 */
func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	p1 := getPath(root, p)
	p2 := getPath(root, q)

	var ans *TreeNode
	for i:=0; i<len(p1) && i<len(p2); i++ {
		if p1[i].Val != p2[i].Val {
			break
		} else {
			ans = p1[i]
		}
	}
	return ans
}

func getPath(root, t *TreeNode) []*TreeNode {
	path := make([]*TreeNode, 0)
	var lastVisited *TreeNode
	for root != nil || path != nil {
		if root != nil {
			path = append(path, root)
			root = root.Left
		} else {
			node := path[len(path)-1]
			if node.Right != nil && lastVisited != node.Right {
				root = node.Right
			} else {
				if node == t {
					return path
				}
				lastVisited = path[len(path)-1]
				path = path[:len(path)-1]
				root = nil
			}
		}
	}
	return path
}