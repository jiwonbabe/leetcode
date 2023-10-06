package _701_Insert_into_a_Binary_Search_Tree_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	curr := root
	for curr != nil {
		if curr.Val < val {
			if curr.Right != nil {
				curr = curr.Right
			} else {
				curr.Right = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
				return root
			}
		} else {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr.Left = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
				return root
			}
		}
	}

	return root
}
