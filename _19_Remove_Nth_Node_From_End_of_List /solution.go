package _19_Remove_Nth_Node_From_End_of_List_

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	if length == 1 {
		return nil
	}

	if length == n {
		return head.Next
	}
	curr = head
	for i := 0; i < length-n-1; i++ {
		curr = curr.Next
	}
	next := curr.Next
	curr.Next = next.Next

	return head
}
