package example

import (
	"fmt"
	"os"
	"testing"
)

func TestNewLinkedList(t *testing.T) {

	linkedList := NewLinkedList()
	linkedList.Append("0:index")
	linkedList.Append("1:index")
	linkedList.Append("3:index")
	fmt.Println(linkedList.Tail().data)
	err2 := linkedList.Insert(3, "2:index")
	if err2 != nil {
		fmt.Println(err2)
	}

	linkedList.Each(func(index uint, node *node) bool {
		fmt.Println(index, node.data)
		return true
	})

	_ = linkedList.Remove(3)
	_ = linkedList.Remove(2)
	//_ = linkedList.Remove(1)
	//_ = linkedList.Remove(0)

	fmt.Println(linkedList.Tail().data, linkedList.Head().data)
	os.Exit(1)
}

func TestNewDoubleLinkedList(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	doubleLinkedList.Append("0:index")
	doubleLinkedList.Append("1:index")
	doubleLinkedList.Each(func(index int, doubleNode *DoubleNode) bool {
		fmt.Println(index, doubleNode.data, doubleNode.prev.data, doubleNode.next.data)
		return true
	})
	fmt.Println("----------")
	_, _ = doubleLinkedList.InsertToHead("2:index")
	//_ = doubleLinkedList.InsertToTail("3:index")
	doubleLinkedList.Each(func(index int, doubleNode *DoubleNode) bool {
		fmt.Println(index, doubleNode.data, doubleNode.prev.data, doubleNode.next.data)
		return true
	})
	node, err := doubleLinkedList.Get(2)
	fmt.Println(node, err)
	//fmt.Println("----", doubleLinkedList.Head(), doubleLinkedList.Tail())
}

func TestNewArray(t *testing.T) {

	array := NewArray(3)

	array.Put("a")
	array.Put("d")
	array.Put("e")

	xx := map[string]string{
		"xx": "yy",
	}

	array.Each(func(index int, data interface{}) bool {
		fmt.Println(index, data)
		fmt.Println(xx)
		return true
	})

	_, err := array.Get(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("----add-----")

	err = array.Add(1, "b")
	err = array.Add(2, "c")
	err = array.Add(5, "f")

	err = array.Update(1, "x")
	if err != nil {
		fmt.Println(err)
		return
	}

	array.Each(func(index int, data interface{}) bool {
		fmt.Println(index, data)
		return true
	})

	err = array.Update(1, "x")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("----remove-----")
	err = array.Remove(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	array.Each(func(index int, data interface{}) bool {
		fmt.Println(index, data)
		return true
	})

	fmt.Println("----put-----")
	array.Put("j")
	array.Each(func(index int, data interface{}) bool {
		fmt.Println(index, data)
		return true
	})
}
