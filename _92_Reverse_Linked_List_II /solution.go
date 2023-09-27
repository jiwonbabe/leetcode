package _92_Reverse_Linked_List_II_

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}

	before := dummy
	index := 1
	for index < left {
		before = before.Next
		index++
	}
	prev := before
	curr := before.Next
	for index <= right {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
		index++
	}
	before.Next.Next = curr
	before.Next = prev

	return dummy.Next
}
