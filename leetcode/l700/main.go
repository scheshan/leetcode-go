package l700

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return nil
}
