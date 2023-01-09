package m16_02

type WordsFrequency struct {
	m map[string]int
}

func Constructor(book []string) WordsFrequency {
	m := make(map[string]int)
	for _, word := range book {
		m[word] = m[word] + 1
	}

	wf := WordsFrequency{
		m: m,
	}

	return wf
}

func (this *WordsFrequency) Get(word string) int {
	return this.m[word]
}
