package _876__Middle_of_the_Linked_List_

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next == nil {
		return slow
	}

	if fast.Next.Next == nil {
		return slow.Next
	}

	return nil
}
