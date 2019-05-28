package data_structure

import "fmt"

type DoubleNode struct {
	Data interface{}
	Prev *DoubleNode
	Next *DoubleNode
}

type DoubleLinkedList struct {
	size int
	head *DoubleNode
	tail *DoubleNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	DoubleLinkedList := &DoubleLinkedList{}
	DoubleLinkedList.init()
	return DoubleLinkedList
}

func (DoubleLinkedList *DoubleLinkedList) init() {
	DoubleLinkedList.size = 0
	DoubleLinkedList.head = nil
	DoubleLinkedList.tail = nil
}

func (DoubleLinkedList *DoubleLinkedList) Append(data interface{}) (doubleNode *DoubleNode) {
	doubleNode = &DoubleNode{
		Data: data,
		Prev: nil,
		Next: nil,
	}
	if DoubleLinkedList.size == 0 {
		DoubleLinkedList.head = doubleNode
	} else {
		doubleNode.Prev = DoubleLinkedList.tail
		doubleNode.Next = DoubleLinkedList.head
		DoubleLinkedList.head.Prev = doubleNode
		DoubleLinkedList.tail.Next = doubleNode
	}
	DoubleLinkedList.tail = doubleNode
	DoubleLinkedList.size++
	return
}

func (DoubleLinkedList *DoubleLinkedList) Insert(index int, data interface{}) (doubleNode *DoubleNode, err error) {

	if DoubleLinkedList.size < index {
		return nil, fmt.Errorf("index %d is out of list", index)
	}

	doubleNode = &DoubleNode{
		Data: data,
		Prev: nil,
		Next: nil,
	}
	if index == 0 {
		doubleNode.Prev = DoubleLinkedList.tail
		doubleNode.Next = DoubleLinkedList.head
		DoubleLinkedList.head.Prev = doubleNode
		DoubleLinkedList.tail.Next = doubleNode
		DoubleLinkedList.head = doubleNode
	} else {
		var i int
		itemNode := DoubleLinkedList.head
		for i = 0; i < index; i++ {
			if i == index-1 {
				doubleNode.Prev = itemNode
				doubleNode.Next = itemNode.Next
				itemNode.Next.Prev = doubleNode
				itemNode.Next = doubleNode
				break
			}
			itemNode = itemNode.Next
		}
		if index == DoubleLinkedList.size {
			DoubleLinkedList.tail.Prev = doubleNode
			DoubleLinkedList.tail = doubleNode
		}
	}
	DoubleLinkedList.size++
	return doubleNode, nil
}

func (DoubleLinkedList *DoubleLinkedList) MovToHead(doubleNode *DoubleNode) {
	if DoubleLinkedList.size == 0 {
		DoubleLinkedList.Append(doubleNode.Data)
		return
	}
	if DoubleLinkedList.size == 1 {
		return
	}
	doubleNode.Prev.Next = doubleNode.Next
	doubleNode.Next.Prev = doubleNode.Prev
	doubleNode.Prev = DoubleLinkedList.tail
	doubleNode.Next = DoubleLinkedList.head
	DoubleLinkedList.tail.Next = doubleNode
	DoubleLinkedList.head.Prev = doubleNode
	DoubleLinkedList.head = doubleNode
	return
}

func (DoubleLinkedList *DoubleLinkedList) RemoveByIndex(index int) {

	return
}

func (DoubleLinkedList *DoubleLinkedList) Remove(doubleNode *DoubleNode) {
	if DoubleLinkedList.size == 1 {
		DoubleLinkedList.init()
		return
	}
	doubleNode.Prev.Next = doubleNode.Next
	doubleNode.Next.Prev = doubleNode.Prev
	DoubleLinkedList.size--
	return
}

func (DoubleLinkedList *DoubleLinkedList) InsertToHead(data interface{}) (doubleNode *DoubleNode, err error) {
	if DoubleLinkedList.size == 0 {
		return DoubleLinkedList.Append(data), nil
	} else {
		return DoubleLinkedList.Insert(0, data)
	}
}

func (DoubleLinkedList *DoubleLinkedList) InsertToTail(data interface{}) (doubleNode *DoubleNode, err error) {
	if DoubleLinkedList.size == 0 {
		return DoubleLinkedList.Append(data), nil
	} else {
		return DoubleLinkedList.Insert(DoubleLinkedList.size, data)
	}
}

func (DoubleLinkedList *DoubleLinkedList) Get(index int) (doubleNode *DoubleNode, err error) {
	if DoubleLinkedList.size == 0 && DoubleLinkedList.head == nil {
		return nil, fmt.Errorf("index %d is out of list", index)
	}
	var i int
	itemNode := DoubleLinkedList.head
	for i = 0; i < index; i++ {
		itemNode = itemNode.Next
	}
	return itemNode, nil
}

func (DoubleLinkedList *DoubleLinkedList) Head() *DoubleNode {
	return DoubleLinkedList.head
}

func (DoubleLinkedList *DoubleLinkedList) Tail() *DoubleNode {
	return DoubleLinkedList.tail
}

func (DoubleLinkedList *DoubleLinkedList) Size() int {
	return DoubleLinkedList.size
}

func (DoubleLinkedList *DoubleLinkedList) Each(fn func(index int, doubleNode *DoubleNode) bool) {
	if DoubleLinkedList.size > 0 && DoubleLinkedList.head != nil {
		itemNode := DoubleLinkedList.head
		var i int
		for i = 0; i < DoubleLinkedList.size; i++ {
			if !fn(i, itemNode) {
				break
			}
			itemNode = itemNode.Next
		}
	}
}
