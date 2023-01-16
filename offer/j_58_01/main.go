package main

func reverseWords(s string) string {
	if s == "" {
		return s
	}

	l := 0
	for l < len(s) && s[l] == ' ' {
		l++
	}

	r := len(s) - 1
	for r >= 0 && s[r] == ' ' {
		r--
	}

	if l > r {
		return ""
	}

	res := make([]byte, r-l+1)
	slow := 0

	for l <= r {
		fast := r
		for fast >= l && s[fast] != ' ' {
			fast--
		}

		for i := fast + 1; i <= r; i++ {
			res[slow] = s[i]
			slow++
		}

		if fast >= l && s[fast] == ' ' {
			for fast >= l && s[fast] == ' ' {
				fast--
			}
			res[slow] = ' '
			slow++
		}
		r = fast
	}

	return string(res[:slow])
}

func main() {
	reverseWords(" ")
}
