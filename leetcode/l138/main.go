package l138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	visit := make(map[*Node]*Node)

	dummy := new(Node)
	tail := dummy

	for head != nil {
		n := clone(head, visit)
		n.Random = clone(head.Random, visit)
		tail.Next = n
		tail = tail.Next

		head = head.Next
	}

	return dummy.Next
}

func clone(node *Node, visit map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if visit[node] != nil {
		return visit[node]
	}

	res := new(Node)
	res.Val = node.Val
	visit[node] = res
	return res
}
