package _103_Binary_Tree_Zigzag_Level_Order_Traversal_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{
		root,
	}
	level := 1
	result := make([][]int, 0)
	for len(queue) > 0 {
		currLevel := queue
		levelNums := make([]int, 0)
		for i := 0; i < len(currLevel); i++ {
			q := currLevel[i]
			queue = queue[1:]
			levelNums = append(levelNums, q.Val)
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		if level%2 == 0 {
			for i := 0; i < len(levelNums)/2; i++ {
				j := len(levelNums) - i - 1
				levelNums[i], levelNums[j] = levelNums[j], levelNums[i]
			}
		}
		result = append(result, levelNums)
		level++
	}

	return result
}
