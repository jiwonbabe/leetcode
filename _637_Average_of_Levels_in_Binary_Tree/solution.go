package _637_Average_of_Levels_in_Binary_Tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	result := make([]float64, 0)
	for len(queue) > 0 {
		length := len(queue)
		sum := 0
		for i := 0; i < length; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, float64(sum)/float64(length))
		queue = queue[length:]
	}

	return result
}
