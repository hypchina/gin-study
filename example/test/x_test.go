package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type X interface {
	Id() string
}

type UserX struct {
}

func (User UserX) Id() string {
	return "Hello"
}

func NewX(fn func()) X {
	fn()
	return UserX{}
}

func TestHttp(t *testing.T) {

	resp, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s,err:=ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))

}

func TestX(t *testing.T) {
	x := 1
	NewX(func() {
		x += 2
	})
	{
		NewX(func() {
			x++
			NewX(func() {
				x++
			})
		})
		NewX(func() {
			x++
		})
	}
	fmt.Println(x)
}
