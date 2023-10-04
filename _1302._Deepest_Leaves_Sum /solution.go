package _1302__Deepest_Leaves_Sum_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{
		root,
	}

	var currLevelSum int
	for len(queue) > 0 {
		currLevelLen := len(queue)
		currLevelSum = 0
		for i := 0; i < currLevelLen; i++ {
			q := queue[0]
			currLevelSum += q.Val
			queue = queue[1:]
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
	}

	return currLevelSum

}
