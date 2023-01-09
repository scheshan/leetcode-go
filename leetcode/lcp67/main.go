package lcp67

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func expandBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		root.Left = &TreeNode{Val: -1, Left: expandBinaryTree(root.Left)}
	}
	if root.Right != nil {
		root.Right = &TreeNode{Val: -1, Right: expandBinaryTree(root.Right)}
	}
	return root
}
