package _117_Populating_Next_Right_Pointers_in_Each_Node_II

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	leftMost := root
	for leftMost != nil {
		curr := leftMost
		var prev *Node
		leftMost = nil

		for curr != nil {
			if curr.Left != nil {
				if prev == nil {
					leftMost = curr.Left
				} else {
					prev.Next = curr.Left
				}
				prev = curr.Left
			}

			if curr.Right != nil {
				if prev == nil {
					leftMost = curr.Right
				} else {
					prev.Next = curr.Right
					prev = prev.Next
				}
				prev = curr.Right
			}
			curr = curr.Next
		}
	}

	return root
}
