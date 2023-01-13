package main

func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] < target {
		return letters[0]
	}

	l := 0
	r := len(letters)
	res := 0

	for l < r {
		mid := l + (r-l)>>1
		ch := letters[mid]

		if ch <= target {
			l = mid + 1
		} else {
			r = mid
			res = mid
		}
	}

	return letters[res]
}

func main() {
	nextGreatestLetter([]byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'}, 'e')
}
