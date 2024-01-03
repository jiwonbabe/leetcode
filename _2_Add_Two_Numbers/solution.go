package _2_Add_Two_Numbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	result := &ListNode{
		Val:  -1,
		Next: nil,
	}
	curr := result
	sum := 0
	for p1 != nil || p2 != nil {
		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}
		curr.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		sum /= 10
		curr = curr.Next
	}

	if sum != 0 {
		curr.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
	}
	return result.Next
}
