package example

import "fmt"

type node struct {
	data interface{}
	next *node
}

type linkedList struct {
	size uint
	head *node
	tail *node
}

func NewLinkedList() *linkedList {
	var linkedList linkedList
	linkedList.init()
	return &linkedList
}

func (linkedList *linkedList) init() {
	linkedList.head = nil
	linkedList.tail = nil
	linkedList.size = 0
	return
}

func (linkedList *linkedList) Append(data interface{}) {
	node := &node{
		data: data,
		next: nil,
	}
	if linkedList.size == 0 {
		linkedList.head = node
	} else {
		lastEle := linkedList.tail
		lastEle.next = node
	}
	linkedList.tail = node
	linkedList.size++
}

func (linkedList *linkedList) Insert(index uint, data interface{}) error {

	if linkedList.head == nil || linkedList.size == 0 || (linkedList.size) < index {
		return fmt.Errorf("index %d is out of list", index)
	}
	node := &node{
		data: data,
		next: nil,
	}
	if index == 0 {
		firstEle := linkedList.head
		node.next = firstEle
		linkedList.head = node
	} else {
		stepEle := linkedList.head
		var i uint
		for i = 0; i < index; i++ {
			if i == index-1 {
				tmpNode := stepEle.next
				stepEle.next = node
				node.next = tmpNode
				break
			}
			stepEle = stepEle.next
		}
	}
	linkedList.size++
	return nil
}

func (linkedList *linkedList) Remove(index uint) error {
	if linkedList.size == 0 || linkedList.head == nil || (linkedList.size-1) < index {
		return fmt.Errorf("index %d is out of list", index)
	}
	stepEle := linkedList.head
	var i uint
	if index == 0 {
		if linkedList.size == 1 {
			linkedList.init()
			return nil
		}
		linkedList.head = stepEle.next
	} else {
		for i = 0; i < index; i++ {
			if i == index-1 {
				stepEle.next = stepEle.next.next
				break
			}
			stepEle = stepEle.next
		}
		if index == linkedList.size-1 {
			linkedList.tail = stepEle
		}
	}
	linkedList.size--
	return nil
}

func (linkedList *linkedList) Head() *node {
	return linkedList.head
}

func (linkedList *linkedList) Tail() *node {
	return linkedList.tail
}

func (linkedList *linkedList) Size() uint {
	return linkedList.size
}

func (linkedList *linkedList) Each(fn func(index uint, node *node) bool) {
	if linkedList.size > 0 && linkedList.head != nil {
		node := linkedList.head
		var index uint
		index = 0
		for {
			if node == nil {
				break
			}
			if ok := fn(index, node); !ok {
				break
			}
			node = node.next
			index++
		}
	}
}
