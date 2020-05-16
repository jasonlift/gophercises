package structures

import (
	"container/list"
)

/*
Description:
https://leetcode.com/problems/lru-cache/
Reference:
https://www.jianshu.com/p/970f1a8dd9cf
 */

type LRUCache struct {
	capacity int
	dlist *list.List
	cache map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		dlist: list.New(),
		cache: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if this.cache == nil {
		return -1
	}

	elem, hit := this.cache[key]
	if hit {
		this.dlist.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	if this.cache == nil {
		this.cache = make(map[int]*list.Element)
		this.dlist = list.New()
	}
	elem, ok := this.cache[key]
	if ok { // exists
		this.dlist.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	// 将新添加的缓存项放入双向链表的最前端
	e := this.dlist.PushFront(&entry{key, value})
	this.cache[key] = e

	if this.capacity != 0 && len(this.cache) > this.capacity {
		this.RemoveOldest()
	}
}

// RemoveOldest 移除双向链表中访问时间最远的那一项
func (this *LRUCache) RemoveOldest() {
	if this.cache == nil {
		return
	}
	ele := this.dlist.Back()
	if ele != nil {
		this.dlist.Remove(ele)
		kv := ele.Value.(*entry)
		delete(this.cache, kv.key)
	}
}
