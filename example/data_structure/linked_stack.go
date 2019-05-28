package data_structure

import "fmt"

type linkedStackNode struct {
	data interface{}
	prev *linkedStackNode
}

type linkedStack struct {
	maxSize   int
	tail      *linkedStackNode
	stackSize int
}

func NewLinkedStack(maxSize int) *linkedStack {
	linkedStack := &linkedStack{}
	linkedStack.init(maxSize)
	return linkedStack
}

func (linkedStack *linkedStack) init(maxSize int) *linkedStack {
	linkedStack.stackSize = 0
	linkedStack.tail = nil
	linkedStack.maxSize = maxSize
	return linkedStack
}

func (linkedStack *linkedStack) Push(data interface{}) error {

	if linkedStack.maxSize <= linkedStack.stackSize {
		return fmt.Errorf("stackSize is out of maxSize")
	}

	node := &linkedStackNode{
		data: data,
	}

	if linkedStack.stackSize > 0 {
		node.prev = linkedStack.tail
	}

	linkedStack.tail = node
	linkedStack.stackSize++

	return nil
}

func (linkedStack *linkedStack) Pop() (interface{}, error) {

	if linkedStack.stackSize < 1 {
		return nil, fmt.Errorf("stackSize is nil")
	}

	linkedStackNode := linkedStack.tail
	if linkedStack.stackSize == 1 {
		linkedStack.init(linkedStack.maxSize)
		return linkedStackNode.data, nil
	}

	linkedStack.tail = linkedStackNode.prev
	linkedStack.stackSize--

	return linkedStackNode.data, nil
}

func (linkedStack *linkedStack) Size() int {
	return linkedStack.stackSize
}
