package l763

func partitionLabels(S string) []int {
	res := make([]int, 0)

	left := 0

	count := make([]int, 26)
	for i := range S {
		count[S[i]-'a']++
	}

	for left < len(S) {
		cnt := count[S[left]-'a']
		count[S[left]-'a'] = 0
		right := left

		for cnt != right-left+1 && right < len(S) {
			right += 1
			cnt += count[S[right]-'a']
			count[S[right]-'a'] = 0
		}

		res = append(res, right-left+1)
		left = right + 1
	}

	return res
}
