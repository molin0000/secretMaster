package calculator

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	c := NewCalcGame()
	fmt.Println(c.Start())
	fmt.Println(c.GiveResult(986))
	fmt.Println(c.GiveResult(1459))
	fmt.Println(c.GiveResult(1296))

	c = NewCalcGame()
	fmt.Println(c.Start())
	fmt.Println(c.GiveResult(986))
	fmt.Println(c.GiveResult(1458))
	fmt.Println(c.GiveResult(1296))

	c = NewCalcGame()
	fmt.Println(c.Start())
	time.Sleep(11 * time.Second)

	fmt.Println(c.GiveResult(986))
	fmt.Println(c.GiveResult(1459))
	fmt.Println(c.GiveResult(1296))

}
