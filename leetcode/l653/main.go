package l653

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, k, make(map[int]bool))
}

func dfs(node *TreeNode, k int, m map[int]bool) bool {
	if node == nil {
		return false
	}

	if m[k-node.Val] {
		return true
	}

	m[node.Val] = true
	return dfs(node.Left, k, m) || dfs(node.Right, k, m)
}
