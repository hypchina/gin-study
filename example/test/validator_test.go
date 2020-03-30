package test

import (
	"testing"
)

type User struct{
	Username    string  `form:text,valid:required`
	Email       string  `form:text,valid:required|valid_email`
}

func TestValidator(t *testing.T) {

}
