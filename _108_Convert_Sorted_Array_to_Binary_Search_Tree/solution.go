package _108_Convert_Sorted_Array_to_Binary_Search_Tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return makeTree(nums)
}

func makeTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	node := &TreeNode{
		Val: nums[mid],
	}
	node.Left = makeTree(nums[0:mid])
	node.Right = makeTree(nums[mid+1:])
	return node
}
