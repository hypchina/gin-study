package utils

import "reflect"

func Try(f func()) *tryStruct {
	return &tryStruct{
		catches: make(map[reflect.Type]ExceptionHandler),
		hold:    f,
	}
}

type ExceptionHandler func(interface{})

type tryStruct struct {
	catches map[reflect.Type]ExceptionHandler
	hold    func()
}

func (t *tryStruct) Catch(e interface{}, f ExceptionHandler) *tryStruct {
	t.catches[reflect.TypeOf(e)] = f
	return t
}

func (t *tryStruct) Finally(f func()) {
	defer func() {
		if e := recover(); nil != e {
			if h, ok := t.catches[reflect.TypeOf(e)]; ok {
				h(e)
			}
		}
		f()
	}()
	t.hold()
}
