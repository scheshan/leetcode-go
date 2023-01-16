package main

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l <= r {
		if !isLetterNumber(s[l]) {
			l++
		} else if !isLetterNumber(s[r]) {
			r--
		} else if toLower(s[l]) == toLower(s[r]) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isLetterNumber(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func main() {
	isPalindrome("0P")
}
