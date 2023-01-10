package l1325

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Right = removeLeafNodes(root.Right, target)
	root.Left = removeLeafNodes(root.Left, target)

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
