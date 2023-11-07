package _2095_Delete_the_Middle_Node_of_a_Linked_List_

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	d := &ListNode{
		Val:  -1,
		Next: head,
	}
	s, f := head, head
	prev := d
	for f.Next != nil {
		prev = s
		s = s.Next
		f = f.Next
		if f.Next != nil {
			f = f.Next
		}
	}

	prev.Next = s.Next
	return d.Next
}
