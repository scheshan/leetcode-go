package l404

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	return dfs(root, false)
}

func dfs(node *TreeNode, left bool) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		if left {
			return node.Val
		}
		return 0
	}

	return dfs(node.Left, true) + dfs(node.Right, false)
}
