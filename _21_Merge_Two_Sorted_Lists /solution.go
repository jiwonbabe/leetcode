package _21_Merge_Two_Sorted_Lists_

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	result := &ListNode{
		Val: -1,
	}
	curr := result
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}
		curr = curr.Next
	}

	if p1 != nil && p2 == nil {
		curr.Next = p1
	} else if p2 != nil && p1 == nil {
		curr.Next = p2
	}
	return result.Next
}
