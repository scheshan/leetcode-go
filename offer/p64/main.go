package p64

type State struct {
	total int
}

func sumNums(n int) int {
	s := &State{}
	recursive(s, n)
	return s.total
}

func recursive(s *State, n int) bool {
	s.total += n
	return n > 0 && recursive(s, n-1)
}
