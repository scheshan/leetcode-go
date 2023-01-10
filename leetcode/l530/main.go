package l530

import (
	"github.com/scheshan/leetcode/common"
	"math"
)

type TreeNode = common.TreeNode

type State struct {
	pre *TreeNode
	res int
}

func getMinimumDifference(root *TreeNode) int {
	state := &State{}
	dfs(root, state)
	return state.res
}

func dfs(node *TreeNode, state *State) {
	if node == nil {
		return
	}

	dfs(node.Left, state)

	if state.pre == nil {
		state.res = math.MaxInt
	} else {
		n := node.Val - state.pre.Val
		if n < state.res {
			state.res = n
		}
	}
	state.pre = node

	dfs(node.Right, state)
}
