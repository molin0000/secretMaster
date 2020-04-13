package qlog

import (
	"testing"
)

func TestPrint(t *testing.T) {
	msg := "hello world"
	n1 := 123
	Println(msg, "nihao", "wo hao", n1)
	Println(t)
}
