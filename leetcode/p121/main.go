package p121

func maxProfit(prices []int) int {
	res := 0
	lowest := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i]-lowest > res {
			res = prices[i] - lowest
		}
		if prices[i] < lowest {
			lowest = prices[i]
		}
	}

	return res
}
