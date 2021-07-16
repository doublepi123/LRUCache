package LRUCache

import "container/list"

type item struct {
	key   string
	value string
	pos   *list.Element
}

type LRU struct {
	ul      list.List
	data    map[string]*item
	Maxsize int
}

func (lru LRU) Get(key string) string {
	i := lru.data[key]
	if i != nil {
		lru.ul.Remove(i.pos)
		lru.ul.PushBack(i)
		return i.value
	} else {
		return ""
	}
}

func (lru LRU) Set(key string, value string) {
	i := &item{key: key, value: value}
	lru.data[key] = i
	lru.ul.Remove(i.pos)
	lru.ul.PushBack(i)
	if lru.ul.Len() > lru.Maxsize {
		lru.ul.Remove(lru.ul.Front())
	}
	i.pos = lru.ul.Back()
}

func (lru LRU) init() {
	lru.data = make(map[string]*item)
	lru.ul = list.List{}
}
