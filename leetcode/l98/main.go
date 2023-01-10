package l98

import (
	"github.com/scheshan/leetcode/common"
	"math"
)

type TreeNode = common.TreeNode

type State struct {
	max int
}

func isValidBST(root *TreeNode) bool {
	state := &State{max: math.MinInt}
	return dfs(root, state)
}

func dfs(root *TreeNode, state *State) bool {
	if root == nil {
		return true
	}

	if !dfs(root.Left, state) {
		return false
	}

	if root.Val <= state.max {
		return false
	}

	state.max = root.Val

	if !dfs(root.Right, state) {
		return false
	}

	return true
}
