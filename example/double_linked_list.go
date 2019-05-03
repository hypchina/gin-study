package example

import "fmt"

type DoubleNode struct {
	data interface{}
	prev *DoubleNode
	next *DoubleNode
}

type doubleLinkedList struct {
	size int
	head *DoubleNode
	tail *DoubleNode
}

func NewDoubleLinkedList() *doubleLinkedList {
	doubleLinkedList := &doubleLinkedList{}
	doubleLinkedList.init()
	return doubleLinkedList
}

func (doubleLinkedList *doubleLinkedList) init() {
	doubleLinkedList.size = 0
	doubleLinkedList.head = nil
	doubleLinkedList.tail = nil
}

func (doubleLinkedList *doubleLinkedList) Append(data interface{}) (doubleNode *DoubleNode) {
	doubleNode = &DoubleNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if doubleLinkedList.size == 0 {
		doubleLinkedList.head = doubleNode
	} else {
		doubleNode.prev = doubleLinkedList.tail
		doubleNode.next = doubleLinkedList.head
		doubleLinkedList.head.prev = doubleNode
		doubleLinkedList.tail.next = doubleNode
	}
	doubleLinkedList.tail = doubleNode
	doubleLinkedList.size++
	return
}

func (doubleLinkedList *doubleLinkedList) Insert(index int, data interface{}) (doubleNode *DoubleNode, err error) {

	if doubleLinkedList.size < index {
		return nil, fmt.Errorf("index %d is out of list", index)
	}

	doubleNode = &DoubleNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if index == 0 {
		doubleNode.prev = doubleLinkedList.tail
		doubleNode.next = doubleLinkedList.head
		doubleLinkedList.head.prev = doubleNode
		doubleLinkedList.tail.next = doubleNode
		doubleLinkedList.head = doubleNode
	} else {
		var i int
		itemNode := doubleLinkedList.head
		for i = 0; i < index; i++ {
			if i == index-1 {
				doubleNode.prev = itemNode
				doubleNode.next = itemNode.next
				itemNode.next.prev = doubleNode
				itemNode.next = doubleNode
				break
			}
			itemNode = itemNode.next
		}
		if index == doubleLinkedList.size {
			doubleLinkedList.tail.prev = doubleNode
			doubleLinkedList.tail = doubleNode
		}
	}
	doubleLinkedList.size++
	return doubleNode, nil
}

func (doubleLinkedList *doubleLinkedList) MovToHead(doubleNode *DoubleNode) {
	if doubleLinkedList.size == 0 {
		doubleLinkedList.Append(doubleNode.data)
		return
	}
	if doubleLinkedList.size == 1 {
		return
	}
	doubleNode.prev.next = doubleNode.next
	doubleNode.next.prev = doubleNode.prev
	doubleNode.prev = doubleLinkedList.tail
	doubleNode.next = doubleLinkedList.head
	doubleLinkedList.tail.next = doubleNode
	doubleLinkedList.head.prev = doubleNode
	doubleLinkedList.head = doubleNode
	return
}

func (doubleLinkedList *doubleLinkedList) RemoveByIndex(index int) {

	return
}

func (doubleLinkedList *doubleLinkedList) Remove(doubleNode *DoubleNode) {
	if doubleLinkedList.size == 1 {
		doubleLinkedList.init()
		return
	}
	doubleNode.prev.next = doubleNode.next
	doubleNode.next.prev = doubleNode.prev
	doubleLinkedList.size--
	return
}

func (doubleLinkedList *doubleLinkedList) InsertToHead(data interface{}) (doubleNode *DoubleNode, err error) {
	if doubleLinkedList.size == 0 {
		return doubleLinkedList.Append(data), nil
	} else {
		return doubleLinkedList.Insert(0, data)
	}
}

func (doubleLinkedList *doubleLinkedList) InsertToTail(data interface{}) (doubleNode *DoubleNode, err error) {
	if doubleLinkedList.size == 0 {
		return doubleLinkedList.Append(data), nil
	} else {
		return doubleLinkedList.Insert(doubleLinkedList.size, data)
	}
}

func (doubleLinkedList *doubleLinkedList) Get(index int) (doubleNode *DoubleNode, err error) {
	if doubleLinkedList.size == 0 && doubleLinkedList.head == nil {
		return nil, fmt.Errorf("index %d is out of list", index)
	}
	var i int
	itemNode := doubleLinkedList.head
	for i = 0; i < index; i++ {
		itemNode = itemNode.next
	}
	return itemNode, nil
}

func (doubleLinkedList *doubleLinkedList) Head() *DoubleNode {
	return doubleLinkedList.head
}

func (doubleLinkedList *doubleLinkedList) Tail() *DoubleNode {
	return doubleLinkedList.tail
}

func (doubleLinkedList *doubleLinkedList) Size() int {
	return doubleLinkedList.size
}

func (doubleLinkedList *doubleLinkedList) Each(fn func(index int, doubleNode *DoubleNode) bool) {
	if doubleLinkedList.size > 0 && doubleLinkedList.head != nil {
		itemNode := doubleLinkedList.head
		var i int
		for i = 0; i < doubleLinkedList.size; i++ {
			if !fn(i, itemNode) {
				break
			}
			itemNode = itemNode.next
		}
	}
}
