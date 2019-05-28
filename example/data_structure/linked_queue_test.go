package data_structure

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewLinkQueue(t *testing.T) {
	maxSize := 20
	queue := NewLinkQueue(maxSize)

	for i := 0; i < maxSize; i++ {
		err := queue.Push(strconv.Itoa(i) + ":node")
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	queue.Each(func(index int, data interface{}) {
		fmt.Println(index, data)
	})

	fmt.Println("======each queue========")

	for i := 0; i < 10; i++ {
		data, err := queue.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(data)
	}

	fmt.Println("======pop queue========")

	queue.Each(func(index int, data interface{}) {
		fmt.Println(index, data)
	})

	if queue.Size() != 10 {
		t.Fatalf("queue size is error")
	}

	dataX, errX := queue.Pop()
	fmt.Println(dataX, errX)

	fmt.Println("======each queue========")
}
