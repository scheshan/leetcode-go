package lcp44

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func numColor(root *TreeNode) int {
	m := make(map[int]int)
	travel(root, m)

	return len(m)
}

func travel(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}

	m[root.Val] = m[root.Val] + 1
	travel(root.Left, m)
	travel(root.Right, m)
}
