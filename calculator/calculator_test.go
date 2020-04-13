package calculator

import (
	"testing"
	"time"

	"github.com/molin0000/secretMaster/qlog"
)

func TestRun(t *testing.T) {
	c := NewCalcGame()
	qlog.Println(c.Start())
	qlog.Println(c.GiveResult(986))
	qlog.Println(c.GiveResult(1459))
	qlog.Println(c.GiveResult(1296))

	c = NewCalcGame()
	qlog.Println(c.Start())
	qlog.Println(c.GiveResult(986))
	qlog.Println(c.GiveResult(1458))
	qlog.Println(c.GiveResult(1296))

	c = NewCalcGame()
	qlog.Println(c.Start())
	time.Sleep(11 * time.Second)

	qlog.Println(c.GiveResult(986))
	qlog.Println(c.GiveResult(1459))
	qlog.Println(c.GiveResult(1296))

}
