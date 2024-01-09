package _82_Remove_Duplicates_from_Sorted_List_II

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}

	prev := dummy
	curr := prev.Next
	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}
		curr = curr.Next
	}
	return dummy.Next
}
