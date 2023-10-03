package _543__Diameter_of_Binary_Tree_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *int) int {
	if root == nil {
		return -1
	}

	left := dfs(root.Left, result)
	right := dfs(root.Right, result)

	cand := left + right + 2
	if *result < cand {
		*result = cand
	}

	if left < right {
		return right + 1
	} else {
		return left + 1
	}
}
