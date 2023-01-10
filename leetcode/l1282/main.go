package main

func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int)

	res := make([][]int, 0)

	for i, c := range groupSizes {
		if c == 1 {
			res = append(res, []int{i})
			continue
		}

		group := m[c]
		if group == nil {
			group = make([]int, 0, c)
		}
		group = append(group, i)
		if len(group) == c {
			res = append(res, group)
			delete(m, c)
		} else {
			m[c] = group
		}
	}
	return res
}
