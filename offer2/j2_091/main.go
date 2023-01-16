package j2_091

func minCost(costs [][]int) int {
	for i := 1; i < len(costs); i++ {
		costs[i][0] += min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += min(costs[i-1][0], costs[i-1][1])
	}

	last := len(costs) - 1
	return min(costs[last][0], min(costs[last][1], costs[last][2]))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
