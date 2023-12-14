package _101_Symmetric_Tree_

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return traverse(root.Left, root.Right)
}

func traverse(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return traverse(left.Left, right.Right) && traverse(left.Right, right.Left)
}
