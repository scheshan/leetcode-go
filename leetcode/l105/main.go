package l105

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}

	return build(preorder, inorder, 0, len(preorder), 0, len(inorder), m)
}

func build(preorder []int, inorder []int, l1 int, r1 int, l2 int, r2 int, m map[int]int) *TreeNode {
	if l1 >= r1 {
		return nil
	}

	root := &TreeNode{Val: preorder[l1]}

	left := m[root.Val] - l2

	root.Left = build(preorder, inorder, l1+1, l1+1+left, l2, m[root.Val], m)
	root.Right = build(preorder, inorder, l1+1+left, r1, m[root.Val]+1, r2, m)
	return root
}
