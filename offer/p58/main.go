package p58

func reverseLeftWords(s string, n int) string {
	arr := make([]byte, len(s))
	if n >= len(s) {
		n = 0
	}

	ind := 0
	for i := n; i < len(arr); i++ {
		arr[ind] = s[i]
		ind++
	}
	for i := 0; i < n; i++ {
		arr[ind] = s[i]
		ind++
	}
	return string(arr)
}
