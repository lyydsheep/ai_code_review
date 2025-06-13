package test

import (
	"errors"
	"fmt"
	"github.com/faiz/llm-code-review/common/errcode"
	"gorm.io/gorm"
	"testing"
)

func TestErr(t *testing.T) {
	err := A()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("hello world")
	}
	fmt.Println("aaa")
}

func A() error {
	return errcode.ErrServer.WithCause(gorm.ErrRecordNotFound)
}
