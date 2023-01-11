package j35

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	return copy(m, head)
}

func copy(m map[*Node]*Node, node *Node) *Node {
	if node == nil {
		return nil
	}

	if m[node] != nil {
		return m[node]
	}

	res := new(Node)
	m[node] = res
	res.Val = node.Val

	res.Random = copy(m, node.Random)
	res.Next = copy(m, node.Next)
	return res
}
