package l2379

func minimumRecolors(blocks string, k int) int {
	res := 0
	r := 0

	for r < k {
		if blocks[r] != 'B' {
			res++
		}
		r++
	}

	cnt := res
	for r < len(blocks) {
		if blocks[r] != 'B' {
			cnt++
		}
		if blocks[r-k] != 'B' {
			cnt--
		}

		if cnt < res {
			res = cnt
		}
		r++
	}
	return res
}
