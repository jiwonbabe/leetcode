package _83_Remove_Duplicates_from_Sorted_List_

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	prev := &ListNode{
		Val:  -101,
		Next: head,
	}
	dp := head
	for dp != nil {
		if prev.Val == dp.Val {
			prev.Next = dp.Next
			dp = dp.Next
		} else {
			prev = dp
			dp = dp.Next
		}
	}

	return head
}
