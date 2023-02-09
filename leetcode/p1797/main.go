package p1797

import "container/list"

type AuthenticationManager struct {
	ttl   int
	m     map[string]*Item
	queue *list.List
}

type Item struct {
	key  string
	time int
	del  bool
}

func Constructor(timeToLive int) AuthenticationManager {
	res := AuthenticationManager{
		ttl:   timeToLive,
		m:     make(map[string]*Item),
		queue: list.New(),
	}
	return res
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.prune(currentTime)

	if exist := this.m[tokenId]; exist != nil {
		exist.del = true
	}
	item := &Item{
		key:  tokenId,
		time: currentTime,
	}
	this.m[tokenId] = item
	this.queue.PushBack(item)
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	this.prune(currentTime)

	if exist, ok := this.m[tokenId]; ok {
		exist.del = true
		this.Generate(tokenId, currentTime)
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	this.prune(currentTime)

	return len(this.m)
}

func (this *AuthenticationManager) prune(currentTime int) {
	for this.queue.Len() > 0 {
		front := this.queue.Front()
		item := front.Value.(*Item)
		if item.del {
			this.queue.Remove(front)
			continue
		}
		if item.time <= currentTime-this.ttl {
			this.queue.Remove(front)
			delete(this.m, item.key)
			continue
		}
		break
	}
}
