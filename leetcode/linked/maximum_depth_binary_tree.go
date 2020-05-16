package linked

/*
Description:
https://leetcode.com/problems/maximum-depth-of-binary-tree/
Example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	hl := maxDepth(root.Left)+1
	hr := maxDepth(root.Right)+1
	if hl > hr {
		return hl
	} else {
		return hr
	}
}

func maxDepthSelf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return middleTraverse(root, 0)
}

func middleTraverse(p *TreeNode, level int) int {
	if p == nil {
		return level
	}

	lMax := middleTraverse(p.Left, level+1)
	rMax := middleTraverse(p.Right, level+1)
	if lMax > rMax {
		return lMax
	} else {
		return rMax
	}
}
