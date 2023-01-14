package l104

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	n1 := maxDepth(root.Left)
	n2 := maxDepth(root.Right)

	res := n1
	if n2 > res {
		res = n2
	}

	return res + 1
}
