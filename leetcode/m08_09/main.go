package m08_09

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	tmp := NewCharBuilder(n << 1)
	perm(&res, tmp, 0, n)
	return res
}

type CharBuilder struct {
	buffer []byte
	size   int
}

func (t *CharBuilder) Append(ch byte) {
	t.buffer[t.size] = ch
	t.size++
}

func (t *CharBuilder) Len() int {
	return t.size
}

func (t *CharBuilder) String() string {
	return string(t.buffer)
}

func (t *CharBuilder) RemoveLast() {
	t.size--
}

func NewCharBuilder(n int) *CharBuilder {
	cb := new(CharBuilder)
	cb.buffer = make([]byte, n)
	return cb
}

func perm(res *[]string, tmp *CharBuilder, cur int, n int) {
	if tmp.Len() == n<<1 {
		*res = append(*res, tmp.String())
		return
	}

	if cur < n {
		tmp.Append('(')
		perm(res, tmp, cur+1, n)
		tmp.RemoveLast()
	}

	if tmp.Len() < cur*2 {
		tmp.Append(')')
		perm(res, tmp, cur, n)
		tmp.RemoveLast()
	}
}
