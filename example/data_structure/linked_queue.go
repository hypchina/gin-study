package data_structure

import "fmt"

type linkedQueueNode struct {
	data interface{}
	next *linkedQueueNode
}

type linkedQueue struct {
	maxSize   int
	head      *linkedQueueNode
	tail      *linkedQueueNode
	queueSize int
}

func NewLinkQueue(maxSize int) *linkedQueue {
	linkedQueue := &linkedQueue{}
	linkedQueue.init(maxSize)
	return linkedQueue
}

func (linkedQueue *linkedQueue) init(maxSize int) *linkedQueue {
	linkedQueue.maxSize = maxSize
	linkedQueue.queueSize = 0
	linkedQueue.head = nil
	linkedQueue.tail = nil
	return linkedQueue
}

func (linkedQueue *linkedQueue) Push(data interface{}) error {

	if linkedQueue.maxSize <= linkedQueue.queueSize {
		return fmt.Errorf("queueSize is out of maxSize")
	}

	linkedQueueNode := &linkedQueueNode{
		data: data,
	}

	if linkedQueue.queueSize == 0 {
		linkedQueue.head = linkedQueueNode
		linkedQueue.queueSize++
		return nil
	}

	if linkedQueue.queueSize == 1 {
		linkedQueue.tail = linkedQueueNode
		linkedQueue.head.next = linkedQueueNode
		linkedQueue.queueSize++
		return nil
	}

	linkedQueue.tail.next = linkedQueueNode
	linkedQueue.tail = linkedQueueNode

	linkedQueue.queueSize++
	return nil
}

func (linkedQueue *linkedQueue) Pop() (interface{}, error) {

	if linkedQueue.queueSize < 1 {
		return nil, fmt.Errorf("queueSize is nil")
	}

	headNode := linkedQueue.head
	if linkedQueue.queueSize == 1 {
		linkedQueue.init(linkedQueue.maxSize)
		return headNode.data, nil
	}

	linkedQueue.head = headNode.next
	linkedQueue.queueSize--
	return headNode.data, nil
}

func (linkedQueue *linkedQueue) Size() int {
	return linkedQueue.queueSize
}

func (linkedQueue *linkedQueue) Each(fn func(index int, data interface{})) {
	if linkedQueue.queueSize > 0 {
		headNode := linkedQueue.head
		for i := 0; i < linkedQueue.queueSize; i++ {
			fn(i, headNode.data)
			headNode = headNode.next
			if headNode == nil {
				break
			}
		}
	}
}
