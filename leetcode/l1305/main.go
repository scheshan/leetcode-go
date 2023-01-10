package main

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	flat(root1, &arr1)
	flat(root2, &arr2)

	return combine(arr1, arr2)
}

func flat(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	flat(node.Left, arr)

	*arr = append(*arr, node.Val)

	flat(node.Right, arr)
}

func combine(arr1, arr2 []int) []int {
	res := make([]int, len(arr1)+len(arr2))
	l := 0
	r := 0

	ind := 0
	for l < len(arr1) && r < len(arr2) {
		if arr1[l] < arr2[r] {
			res[ind] = arr1[l]
			l++
		} else {
			res[ind] = arr2[r]
			r++
		}
		ind++
	}

	for l < len(arr1) {
		res[ind] = arr1[l]
		ind++
		l++
	}
	for r < len(arr2) {
		res[ind] = arr2[r]
		ind++
		r++
	}
	return res
}

func main() {
	combine([]int{1, 2, 4}, []int{0, 1, 3})
}
