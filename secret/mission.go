package secret

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/molin0000/secretMaster/mission"
)

func (b *Bot) checkMission(fromQQ uint64, msg string) string {
	ms := b.getPersonValue("Mission", fromQQ, &MissionState{false, nil}).(*MissionState)
	if !ms.IsPlaying {
		return ""
	}

	if len(msg) < 2 {
		return "请回复：选+数值，如选0，选1"
	}

	msg = strings.TrimLeft(msg, "选")

	selection, err := strconv.Atoi(msg)
	if err != nil {
		strs := strings.Split(msg, "] ")
		if len(strs) <= 1 {
			return ""
		}

		selection, err = strconv.Atoi(strs[1])
		if err != nil {
			return ""
		}
	}

	fmt.Println("event:", ms.Ms.Event)
	msg, finish := ms.Ms.SelectOption(int(ms.Ms.Event), selection)

	if finish {
		msg += ms.Ms.Finish()
		b.setMagic(fromQQ, int(int64(ms.Ms.Magic)))
		b.setMoney(fromQQ, int(int64(ms.Ms.Money)))
		b.setExp(fromQQ, int(int64(ms.Ms.Exp)))
		ms.IsPlaying = false
		ms.Ms = nil
	}

	b.setPersonValue("Mission", fromQQ, ms)

	return msg
}

func (b *Bot) startMission(fromQQ uint64, msg string) string {
	magic := b.getMagic(fromQQ)
	if magic < 50 {
		return "灵力不足"
	}

	b.setMagic(fromQQ, -50)

	mg := mission.NewRandomMission(missionPath)
	b.setPersonValue("Mission", fromQQ, &MissionState{true, mg})
	return mg.ShowEvent(0)
}

func (b *Bot) showMissions() string {
	return mission.ListMission(missionPath)
}
