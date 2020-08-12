package l133

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visit := make(map[*Node]*Node)

	return clone(node, visit)
}

func clone(node *Node, visit map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if visit[node] != nil {
		return visit[node]
	}

	res := new(Node)
	visit[node] = res

	res.Val = node.Val
	res.Neighbors = make([]*Node, len(node.Neighbors))
	for i, n := range node.Neighbors {
		res.Neighbors[i] = clone(n, visit)
	}
	return res
}
