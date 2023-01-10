package main

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func bstFromPreorder(preorder []int) *TreeNode {
	return buildTree(preorder, 0, len(preorder)-1)
}

func buildTree(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	node := &TreeNode{
		Val: nums[left],
	}

	mid := left + 1
	for mid < len(nums) && nums[mid] < node.Val {
		mid++
	}

	node.Left = buildTree(nums, left+1, mid-1)
	node.Right = buildTree(nums, mid, right)
	return node
}

func main() {
	bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
}
