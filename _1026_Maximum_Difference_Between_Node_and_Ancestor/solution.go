package _1026_Maximum_Difference_Between_Node_and_Ancestor

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	result := -1
	helper(root, &result, root.Val, root.Val)
	return result
}

func helper(root *TreeNode, result *int, currMax int, currMin int) {
	if root == nil {
		return
	}

	t1 := math.Abs(float64(currMax - root.Val))
	t2 := math.Abs(float64(currMin - root.Val))

	maxCand := int(math.Max(t1, t2))

	if *result < maxCand {
		*result = maxCand
	}
	if currMax < root.Val {
		currMax = root.Val
	}
	if currMin > root.Val {
		currMin = root.Val
	}

	helper(root.Left, result, currMax, currMin)
	helper(root.Right, result, currMax, currMin)
}
