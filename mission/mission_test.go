package mission

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/molin0000/secretMaster/qlog"
)

func TestRandomMission(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		ms := NewRandomMission("D:\\序列战争版本更新\\酷Q\\coolq\\data\\app\\me.cqp.molin.secretMaster\\mission")
		fmt.Printf("%+v\n\n", ms)
	}
}

func TestRunMission(t *testing.T) {
	for m := 0; m < 10; m++ {
		ms := NewRandomMission("/Users/molin/coolq/data/app/me.cqp.molin.secretMaster/mission")
		qlog.Println(ms.ShowEvent(0))
		for i := 0; i < 100; i++ {
			msg, ret := ms.SelectOption(int(ms.Event), rand.Intn(len(ms.Ms.Events[ms.Event].Options)+1))
			qlog.Println(msg)
			if ret {
				break
			}
		}
		qlog.Println(ms.Finish())
	}
}
