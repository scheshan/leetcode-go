package j_32_03

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	level := 0

	for len(queue) > 0 {
		level++

		size := len(queue)
		r := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[0]

			if level&1 == 1 {
				r[i] = node.Val
			} else {
				r[size-i-1] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}

		res = append(res, r)
	}
	return res
}
