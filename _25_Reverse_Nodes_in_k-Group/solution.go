package _25_Reverse_Nodes_in_k_Group

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}

	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	prev := dummy
	curr = prev.Next
	for i := 0; i < length/k; i++ {
		for j := 0; j < k-1; j++ {
			next := curr.Next
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		prev = curr
		curr = curr.Next
	}

	return dummy.Next
}
