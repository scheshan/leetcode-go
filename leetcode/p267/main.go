package main

func generatePalindromes(s string) []string {
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	res := make([]string, 0)
	if !valid(cnt) {
		return res
	}
	path := make([]byte, 0, (len(s)+1)>>1)

	var backtrack func(int)
	backtrack = func(ind int) {
		if ind == cap(path) {
			res = append(res, generate(path, len(s)))
			return
		}

		if len(path) == cap(path)-1 && len(s)&1 == 1 {
			for i, c := range cnt {
				if c == 1 {
					cnt[i] -= 1
					path = append(path, byte('a'+i))
					backtrack(ind + 1)
					cnt[i] += 1
					path = path[:len(path)-1]
				}
			}
		} else {
			for i, c := range cnt {
				if c >= 2 {
					cnt[i] -= 2
					path = append(path, byte('a'+i))
					backtrack(ind + 1)
					cnt[i] += 2
					path = path[:len(path)-1]
				}
			}
		}
	}
	backtrack(0)
	return res
}

func valid(cnt []int) bool {
	flag := false
	for _, c := range cnt {
		if c&1 == 1 {
			if flag {
				return false
			}
			flag = true
		}
	}
	return true
}

func generate(arr []byte, size int) string {
	res := make([]byte, size)
	for i, c := range arr {
		res[i] = c
		res[size-i-1] = c
	}
	return string(res)
}

func main() {
	generatePalindromes("aabb")
}
