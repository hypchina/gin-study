package example

import (
	"fmt"
	"testing"
)

func TestNewCache(t *testing.T) {

	cache := NewCache(3)

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	cache.Get("b")
	cache.Set("e", 4)

	cache.cacheList.Each(func(index int, doubleNode *DoubleNode) bool {
		fmt.Println(index, doubleNode.data, doubleNode.prev.data, doubleNode.next.data)
		return true
	})

	/**
		fmt.Println("---------")

	//cache.Get("c")
	cache.Set("b", 5)
	fmt.Println(cache.cacheList.Head().data, cache.cacheList.Tail().data)

	fmt.Println("---------")
	cache.cacheList.Each(func(index int, doubleNode *DoubleNode) bool {
		fmt.Println(index, doubleNode.data, doubleNode.prev.data, doubleNode.next.data)
		return true
	})
	*/
}
