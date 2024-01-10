package _61_Rotate_List

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 1
	curr := head
	for curr.Next != nil {
		length++
		curr = curr.Next
	}
	// curr is now at the last node, so point it to the head to make a circle
	// cycle 을 만들면 prev node 들을 search 가능
	curr.Next = head

	k = k % length
	if k == 0 {
		curr.Next = nil // Break the circle, as no rotation is needed
		return head
	}

	// Move (length - k) steps to get to the new head
	for i := 0; i < length-k; i++ {
		curr = curr.Next
	}

	newHead := curr.Next
	curr.Next = nil

	return newHead
}
