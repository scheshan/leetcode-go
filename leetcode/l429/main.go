package l429

import "github.com/scheshan/leetcode/common"

type Node = common.Node

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		row := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[0]
			row[i] = node.Val

			for _, child := range node.Children {
				queue = append(queue, child)
			}
			queue = queue[1:]
		}

		res = append(res, row)
	}
	return res
}
