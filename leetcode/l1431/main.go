package l1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	max := 0

	for _, n := range candies {
		if n > max {
			max = n
		}
	}

	for i, n := range candies {
		if n+extraCandies >= max {
			res[i] = true
		}
	}

	return res
}
