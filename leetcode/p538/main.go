package p538

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	total := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		node.Val += total
		total = node.Val
		dfs(node.Left)
	}
	dfs(root)
	return root
}
