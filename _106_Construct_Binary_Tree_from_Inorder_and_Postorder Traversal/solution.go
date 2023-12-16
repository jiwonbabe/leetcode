package _106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	return build(&inorder, &postorder, 0, len(inorder)-1)
}

func build(inorder *[]int, postorder *[]int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	nodeVal := (*postorder)[len(*postorder)-1]
	node := &TreeNode{
		Val: nodeVal,
	}
	*postorder = (*postorder)[0 : len(*postorder)-1]
	var index int
	for i, e := range *inorder {
		if e == node.Val {
			index = i
			break
		}
	}

	node.Right = build(inorder, postorder, index+1, end)
	node.Left = build(inorder, postorder, start, index-1)

	return node
}
