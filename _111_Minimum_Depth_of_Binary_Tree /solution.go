package _111_Minimum_Depth_of_Binary_Tree_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftTree := minDepth(root.Left)
	rightTree := minDepth(root.Right)

	if leftTree == 0 {
		return rightTree + 1
	} else if rightTree == 0 {
		return leftTree + 1
	} else if leftTree < rightTree {
		return leftTree + 1
	} else {
		return rightTree + 1
	}
}
