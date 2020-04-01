package mission

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomMission(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		ms := NewRandomMission("/Users/molin/Downloads/mission")
		fmt.Printf("%+v\n\n", ms)
	}
}

func TestRunMission(t *testing.T) {
	for m := 0; m < 10; m++ {
		ms := NewRandomMission("/Users/molin/coolq/data/app/me.cqp.molin.secretMaster/mission")
		fmt.Println(ms.ShowEvent(0))
		for i := 0; i < 100; i++ {
			msg, ret := ms.SelectOption(int(ms.Event), rand.Intn(len(ms.Ms.Events[ms.Event].Options)+1))
			fmt.Println(msg)
			if ret {
				break
			}
		}
		fmt.Println(ms.Finish())
	}
}
