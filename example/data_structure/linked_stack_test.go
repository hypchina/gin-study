package data_structure

import (
	"fmt"
	"testing"
)

func TestNewLinkedStack(t *testing.T) {

	maxSize := 20

	linkedStack := NewLinkedStack(maxSize)
	for i := 0; i < maxSize+1; i++ {
		err := linkedStack.Push(i)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	for i := 0; i < maxSize+1; i++ {
		data, err := linkedStack.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(data)
	}

	fmt.Println("size:", linkedStack.Size())
}
