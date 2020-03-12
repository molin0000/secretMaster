package secret

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/molin0000/secretMaster/calculator"
)

func (b *Bot) checkCalc(fromQQ uint64, msg string) string {
	ms := b.getPersonValue("Calc", fromQQ, &CalcState{false, nil}).(*CalcState)
	if !ms.IsPlaying {
		return ""
	}

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

	msg, finish := ms.Calc.GiveResult(uint64(selection))

	if finish {
		b.setMoney(fromQQ, int(int64(ms.Calc.Money)))
		ms.IsPlaying = false
		ms.Calc = nil
	}

	b.setPersonValue("Calc", fromQQ, ms)

	return msg
}

func (b *Bot) startCalc(fromQQ uint64, msg string) string {
	magic := b.getMagic(fromQQ)
	if magic < 10 {
		return "灵性不足"
	}

	b.setMagic(fromQQ, -10)

	mg := calculator.NewCalcGame()
	info := mg.Start()
	b.setPersonValue("Calc", fromQQ, &CalcState{true, mg})
	fmt.Printf("%+v", *mg)
	return info
}
