package main

func translateNum(num int) int {
	path := make([]rune, 0)
	res := make(map[string]bool)

	dfs(num, &path, res)

	return len(res)
}

func dfs(num int, path *[]rune, res map[string]bool) {
	if num == 0 {
		arr := make([]rune, len(*path))
		for i := 0; i < len(*path); i++ {
			arr[len(arr)-i-1] = (*path)[i]
		}
		str := string(arr)
		res[str] = true
		return
	}

	n := num % 10
	*path = append(*path, 'a'+rune(n))
	dfs(num/10, path, res)
	*path = (*path)[:len(*path)-1]

	n = num % 100
	if n >= 10 && n <= 25 {
		*path = append(*path, 'a'+rune(n))
		dfs(num/100, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	translateNum(12258)
}
