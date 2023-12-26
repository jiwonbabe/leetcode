package _199_Binary_Tree_Right_Side_View

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	result := make([]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			left := queue[i]
			if left.Left != nil {
				queue = append(queue, left.Left)
			}
			if left.Right != nil {
				queue = append(queue, left.Right)
			}
		}
		result = append(result, queue[l-1].Val)
		queue = queue[l:]
	}

	return result
}
