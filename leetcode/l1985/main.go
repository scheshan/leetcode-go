package l1985

import "container/heap"

type Heap struct {
	arr []string
}

func (t *Heap) Len() int {
	return len(t.arr)
}

func (t *Heap) Less(i, j int) bool {
	return compare(t.arr[i], t.arr[j])
}

func (t *Heap) Swap(i, j int) {
	t.arr[i], t.arr[j] = t.arr[j], t.arr[i]
}

func (t *Heap) Push(x interface{}) {
	t.arr = append(t.arr, x.(string))
}

func (t *Heap) Pop() interface{} {
	s := t.arr[len(t.arr)-1]
	t.arr = t.arr[:len(t.arr)-1]
	return s
}

func compare(s1, s2 string) bool {
	if len(s1) < len(s2) {
		return true
	} else if len(s1) > len(s2) {
		return false
	} else {
		for ind := 0; ind < len(s1); ind++ {
			if s1[ind] < s2[ind] {
				return true
			} else if s1[ind] > s2[ind] {
				return false
			}
		}
	}
	return true
}

func kthLargestNumber(nums []string, k int) string {
	h := &Heap{arr: make([]string, 0, k)}

	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else {
			if !compare(num, h.arr[0]) {
				heap.Pop(h)
				heap.Push(h, num)
			}
		}
	}

	return heap.Pop(h).(string)
}
