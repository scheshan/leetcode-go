package l559

import "github.com/scheshan/leetcode/common"

type Node = common.Node

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	res := 0
	for _, child := range root.Children {
		n := maxDepth(child)
		if n > res {
			res = n
		}
	}
	return res + 1
}
