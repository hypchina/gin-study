package utils

import (
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v8"
)

func GetValidator() *validator.Validate {
	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)
	return validate
}

func FormatValidatorError(err error) error {
	errorMap, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	for _, x := range errorMap {
		return errors.New(fmt.Sprintf("参数%s,数据类型%s,验证规则%s", x.Field, x.Kind, x.Tag+":"+x.Param))
	}
	return err
}
