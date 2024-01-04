package _138__Copy_List_with_Random_Pointer

//Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cur := head
	nodeMap := make(map[*Node]*Node)
	nodeMap[nil] = nil
	for cur != nil {
		copied := &Node{
			Val: cur.Val,
		}
		nodeMap[cur] = copied
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		newNode := nodeMap[cur]
		newNode.Next = nodeMap[cur.Next]
		newNode.Random = nodeMap[cur.Random]
		cur = cur.Next
	}

	return nodeMap[head]
}
