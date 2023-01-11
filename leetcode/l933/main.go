package l933

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	queue := make([]int, 0)
	rc := RecentCounter{queue: queue}

	return rc
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.queue) > 0 && this.queue[0] < t-3000 {
		this.queue = this.queue[1:]
	}

	this.queue = append(this.queue, t)
	return len(this.queue)
}
