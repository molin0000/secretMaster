package qlog

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPrint(t *testing.T) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	Println(dir, err)
	msg := "hello world"
	n1 := 123
	Println(msg, "nihao", "wo hao", n1)
	Println(t)
}
