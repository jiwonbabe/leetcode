package _141_Linked_List_Cycle

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	s, f := head, head.Next
	for s != f {
		if f == nil || f.Next == nil {
			return false
		}
		s = s.Next
		f = f.Next.Next
	}

	return true
}
