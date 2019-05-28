package algorithm

import "gin-study/example/data_structure"

type entry struct {
	key string
	val interface{}
}

type cache struct {
	cacheSize int
	cacheList *data_structure.DoubleLinkedList
	cacheMap  map[string]*data_structure.DoubleNode
}

func NewCache(limit int) *cache {
	return &cache{
		cacheSize: limit,
		cacheList: data_structure.NewDoubleLinkedList(),
		cacheMap:  make(map[string]*data_structure.DoubleNode),
	}
}

func (cache *cache) Set(key string, val interface{}) {

	if cache.cacheMap == nil {
		cache.cacheList = data_structure.NewDoubleLinkedList()
		cache.cacheMap = make(map[string]*data_structure.DoubleNode)
	}

	if entryEle, exists := cache.cacheMap[key]; exists {
		cache.cacheList.MovToHead(entryEle)
		entryEle.Data.(*entry).val = val
		return
	}

	ele, _ := cache.cacheList.InsertToHead(&entry{
		key: key,
		val: val,
	})
	cache.cacheMap[key] = ele
	if cache.cacheSize > 0 && cache.cacheList.Size() > cache.cacheSize {
		lastEle := cache.cacheList.Tail()
		cache.removeElement(lastEle)
	}
}

func (cache *cache) Get(key string) (val interface{}, ok bool) {
	if cache.cacheMap == nil {
		return
	}
	if ele, hit := cache.cacheMap[key]; hit {
		cache.cacheList.MovToHead(ele)
		return ele.Data.(*entry).val, true
	}
	return
}

func (cache *cache) Remove(key string) {
	if cache.cacheMap == nil {
		return
	}
	if ele, hit := cache.cacheMap[key]; hit {
		cache.removeElement(ele)
	}
}

func (cache *cache) Clear() {
	cache.cacheMap = nil
	cache.cacheList = nil
}

func (cache *cache) All() *data_structure.DoubleLinkedList {
	return cache.cacheList
}

func (cache *cache) removeElement(ele *data_structure.DoubleNode) {
	cache.cacheList.Remove(ele)
	delete(cache.cacheMap, ele.Data.(*entry).key)
}
