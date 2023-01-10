package l637

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	res := make([]float64, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		var total float64 = 0
		var n float64 = 0

		for i := 0; i < size; i++ {
			node := queue[0]
			total += float64(node.Val)
			n++

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
		}

		res = append(res, total/n)
	}

	return res
}
