package example

import "fmt"

type Array struct {
	data []interface{}
	size int
}

func NewArray(initCapacity int) *Array {
	array := &Array{}
	array.init(initCapacity)
	array.size = 0
	return array
}

func (array *Array) Put(data interface{}) {
	if array.size+1 > len(array.data) {
		array.resize()
	}
	array.data[array.size] = data
	array.size++
}

func (array *Array) Get(index int) (data interface{}, err error) {
	if index >= array.size {
		return nil, fmt.Errorf("index %d is out of array", index)
	}
	return array.data[index], nil
}

func (array *Array) Update(index int, data interface{}) error {
	if index > array.size-1 {
		return fmt.Errorf("index %d is out of array", index)
	}
	array.data[index] = data
	return nil
}

func (array *Array) Add(index int, data interface{}) error {

	if index == 0 {
		array.Put(data)
		return nil
	}

	if index > array.size {
		return fmt.Errorf("index %d is out of array", index)
	}

	if array.size+1 > len(array.data) {
		array.resize()
	}

	if index == array.size {
		array.data[array.size] = data
		array.size++
		return nil
	}

	for i := array.size - 1; i > array.size-index-2; i-- {
		array.data[i+1] = array.data[i]
	}

	array.data[index] = data
	array.size++
	return nil
}

func (array *Array) Remove(index int) error {
	if index > array.size-1 {
		return fmt.Errorf("index %d is out of array", index)
	}
	for i := index; i < array.size-1; i++ {
		array.data[i] = array.data[i+1]
	}
	array.size--
	return nil
}

func (array *Array) Each(fn func(index int, data interface{}) bool) {
	if array.size > 0 {
		for i := 0; i < array.size; i++ {
			if !fn(i, array.data[i]) {
				break
			}
		}
	}
}

func (array *Array) init(capacity int) {
	array.data = make([]interface{}, capacity)
	array.size = 0
}

func (array *Array) resize() {
	size := array.size
	data := array.data
	array.init(len(array.data) * 2)
	for i := 0; i < size; i++ {
		array.data[i] = data[i]
		array.size++
	}
}
