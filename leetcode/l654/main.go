package l654

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructTree(nums, 0, len(nums)-1)
}

func constructTree(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	max := nums[left]
	mi := left

	for i := left + 1; i <= right; i++ {
		if nums[i] > max {
			max = nums[i]
			mi = i
		}
	}

	node := &TreeNode{
		Val: nums[mi],
	}
	node.Left = constructTree(nums, left, mi-1)
	node.Right = constructTree(nums, mi+1, right)
	return node
}
