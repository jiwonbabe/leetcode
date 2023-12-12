package _104_Maximum_Depth_of_Binary_Tree_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return traverse(root, 0)
}

func traverse(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth++
	return max(traverse(root.Left, depth), traverse(root.Right, depth))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
