package _270_Closest_Binary_Search_Tree_Value

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	curr := root
	result := root.Val
	for curr != nil {
		diff := math.Abs(target - float64(result))
		candDiff := math.Abs(target - float64(curr.Val))

		if diff > candDiff {
			result = curr.Val
		} else if diff == candDiff && (result > curr.Val) {
			result = curr.Val
		}
		if float64(curr.Val) > target {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return result
}
