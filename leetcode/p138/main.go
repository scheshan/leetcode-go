package p138

import "github.com/scheshan/leetcode/common"

type Node = common.Node

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	return copy0(head, m)
}

func copy0(node *Node, m map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if exist, ok := m[node]; ok {
		return exist
	} else {
		exist = &Node{
			Val: node.Val,
		}
		m[node] = exist
		exist.Next = copy0(node.Next, m)
		exist.Random = copy0(node.Random, m)
		return exist
	}
}
