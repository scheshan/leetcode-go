package m04_02

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) >> 1
	node := &TreeNode{Val: nums[mid]}
	node.Left = buildTree(nums, left, mid-1)
	node.Right = buildTree(nums, mid+1, right)
	return node
}
