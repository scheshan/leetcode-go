package main

import "container/list"

func wallsAndGates(rooms [][]int) {
	queue := list.New()
	for i, _ := range rooms {
		for j, _ := range rooms[i] {
			if rooms[i][j] == 0 {
				queue.PushBack([]int{i, j})
			}
		}
	}

	step := 0
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			p := queue.Remove(queue.Front()).([]int)
			if rooms[p[0]][p[1]] > step {
				rooms[p[0]][p[1]] = step
			}

			addQueue(queue, rooms, p[0]-1, p[1])
			addQueue(queue, rooms, p[0]+1, p[1])
			addQueue(queue, rooms, p[0], p[1]-1)
			addQueue(queue, rooms, p[0], p[1]+1)
		}
		step++
	}
}

func addQueue(queue *list.List, rooms [][]int, i int, j int) {
	if i < 0 || j < 0 || i >= len(rooms) || j >= len(rooms[i]) || rooms[i][j] != 2147483647 {
		return
	}

	queue.PushBack([]int{i, j})
}

func main() {
	rooms := [][]int{
		[]int{2147483647, 0, 2147483647, 2147483647, 0, 2147483647, -1, 2147483647},
	}
	wallsAndGates(rooms)
}
