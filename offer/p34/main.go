package p34

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	total := 0

	var backtrack func(*TreeNode)
	backtrack = func(node *TreeNode) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		total += node.Val
		if node.Left == nil && node.Right == nil {
			if total == target {
				arr := make([]int, len(path))
				copy(arr, path)
				res = append(res, arr)
			}
		} else {
			backtrack(node.Left)
			backtrack(node.Right)
		}
		total -= node.Val
		path = path[:len(path)-1]
	}
	backtrack(root)
	return res
}
