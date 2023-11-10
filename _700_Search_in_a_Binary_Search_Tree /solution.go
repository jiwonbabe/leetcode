package _700_Search_in_a_Binary_Search_Tree_

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := searchBST(root.Left, val)
	right := searchBST(root.Right, val)
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
