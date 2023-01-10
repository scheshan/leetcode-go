package l107

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		row := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			row[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		res = append(res, row)
	}

	l := 0
	r := len(res) - 1

	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}

	return res
}
