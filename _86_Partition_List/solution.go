package _86_Partition_List

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// two pointer
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	// 두 개의 더미 노드와 두 개의 포인터를 생성합니다.
	beforeDummy := &ListNode{}
	afterDummy := &ListNode{}
	before := beforeDummy
	after := afterDummy

	// 리스트를 순회하면서 각 노드를 적절한 리스트에 추가합니다.
	curr := head
	for curr != nil {
		if curr.Val < x {
			before.Next = curr
			before = before.Next
		} else {
			after.Next = curr
			after = after.Next
		}
		curr = curr.Next
	}

	// 'before' 리스트의 끝과 'after' 리스트의 시작을 연결합니다.
	before.Next = afterDummy.Next
	// 'after' 리스트의 끝을 null로 설정합니다.
	after.Next = nil
	return beforeDummy.Next
}
