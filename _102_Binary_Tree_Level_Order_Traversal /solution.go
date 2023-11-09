package _102_Binary_Tree_Level_Order_Traversal_

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	q := make([]TreeNode, 0)
	q = append(q, *root)
	for len(q) > 0 {
		qCopy := q
		currLevel := make([]int, 0)
		for i := 0; i < len(qCopy); i++ {
			left := q[0]
			currLevel = append(currLevel, left.Val)
			q = q[1:]
			if left.Left != nil {
				q = append(q, *left.Left)
			}
			if left.Right != nil {
				q = append(q, *left.Right)
			}
		}
		result = append(result, currLevel)
	}

	return result
}
