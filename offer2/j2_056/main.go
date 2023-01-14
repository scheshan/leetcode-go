package j2_056

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, root, k)
}

func dfs(root *TreeNode, node *TreeNode, k int) bool {
	if node == nil {
		return false
	}

	if n := find(root, k-node.Val); n != nil && n != node {
		return true
	}
	return dfs(root, node.Left, k) || dfs(root, node.Right, k)
}

func find(root *TreeNode, n int) *TreeNode {
	for root != nil {
		if root.Val == n {
			return root
		} else if root.Val > n {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
