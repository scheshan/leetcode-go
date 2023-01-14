package j_50

func firstUniqChar(s string) byte {
	arr := make([]int, 26)

	for _, ch := range s {
		arr[ch-'a']++
	}

	for _, ch := range s {
		if arr[ch-'a'] == 1 {
			return byte(ch)
		}
	}
	return ' '
}
