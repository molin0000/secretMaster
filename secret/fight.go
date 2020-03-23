package secret

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/molin0000/secretMaster/pet"
)

func (b *Bot) getBattleField() string {
	b.updateBattleField()

	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)
	info := fmt.Sprintf("\n==================\n值夜者小队：(胜：%d，负：%d)", bf.NightwatchWinCnt, bf.NightwatchLoseCnt)

	for i, v := range bf.Nightwatches {
		state := ""
		if v.State == 0 {
			state = "正常"
		} else {
			state = "休整"
		}
		info += fmt.Sprintf("\n%d) %s (%s) 金镑%d, 经验%d", i, v.Nick, state, v.Money/10000, v.Exp/10000)
	}

	info += fmt.Sprintf("\n==================\n赏金猎人：(胜：%d，负：%d)", bf.MoneyHunterWinCnt, bf.MoneyHunterLoseCnt)
	for i, v := range bf.MoneyHunter {
		state := ""
		if v.State == 0 {
			state = "正常"
		} else {
			state = "休整"
		}
		info += fmt.Sprintf("\n%d) %s (%s) 金镑%d, 经验%d", i, v.Nick, state, v.Money/10000, v.Exp/10000)
	}
	info += "\n=================="

	return info
}

func (b *Bot) updateBattleField() {
	fmt.Println("updateBattleField")
	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)
	fmt.Printf("%+v", bf)
	timeNow := uint64(time.Now().Unix())
	for i, v := range bf.Nightwatches {
		duration := timeNow - v.UpdateTime
		if v.State == 2 {
			if v.UpdateTime < v.FailTime {
				duration = v.FailTime - v.UpdateTime
			} else {
				duration = 0
			}
		}
		bf.Nightwatches[i].UpdateTime = timeNow
		bf.Nightwatches[i].Money += duration * workList[0].MoneyAdd * 10000 / 3600 // Zoom in 10000 times to calc
		bf.Nightwatches[i].Exp += duration * workList[0].ExpAdd * 10000 / 3600
		if v.State == 1 {
			if timeNow-v.FailTime > 3600 {
				bf.Nightwatches[i].State = 0
			}
		}
		b.setPersonValue("BattleInfo", v.QQ, bf.Nightwatches[i])
	}

	for i, v := range bf.MoneyHunter {
		duration := timeNow - v.UpdateTime
		if v.State == 2 {
			if v.UpdateTime < v.FailTime {
				duration = v.FailTime - v.UpdateTime
			} else {
				duration = 0
			}
		}
		bf.MoneyHunter[i].UpdateTime = timeNow
		bf.MoneyHunter[i].Money += duration * workList[0].MoneyAdd * 10000 / 3600 // Zoom in 10000 times to calc
		bf.MoneyHunter[i].Exp += duration * workList[0].ExpAdd * 10000 / 3600
		if v.State == 1 {
			if timeNow-v.FailTime > 3600 {
				bf.MoneyHunter[i].State = 0
			}
		}
		b.setPersonValue("BattleInfo", v.QQ, bf.MoneyHunter[i])
	}

	b.setGroupValue("BattleField", bf)
}

func (b *Bot) joinBattleField(fromQQ uint64, fieldType int64) (string, bool) {
	p := b.getPersonFromDb(fromQQ)

	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{fromQQ, p.Name, 0, 0, 0, 0, uint64(time.Now().Unix()), 0, 15, 0, 0, uint64(fieldType)}).(*BattleInfo)
	if bi.ExitTime != 0 {
		if uint64(time.Now().Unix()/3600/24)-bi.ExitTime == 0 {
			return "今天加入阵营次数已经用尽，请明天再试。", false
		}
	}
	bi.State = 0
	bi.Money = 0
	bi.Exp = 0
	bi.FieldType = uint64(fieldType)

	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)
	if fieldType == 0 {
		bf.Nightwatches = append(bf.Nightwatches, bi)
	} else {
		bf.MoneyHunter = append(bf.MoneyHunter, bi)
	}
	b.setPersonValue("BattleInfo", fromQQ, bi)
	b.setGroupValue("BattleField", bf)
	return "\n加入阵营成功", true
}

func (b *Bot) getBattleInfo(fromQQ uint64) string {
	p := b.getPersonFromDb(fromQQ)
	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{fromQQ, p.Name, 0, 0, 0, 0, uint64(time.Now().Unix()), 0, 15, 0, 0, 0}).(*BattleInfo)
	return fmt.Sprintf("\n===============\n个人战绩: 胜%d次，负%d次\n当前收入：%d金镑, %d经验\n===============",
		bi.WinCnt, bi.FailCnt, bi.Money/10000, bi.Exp/10000)
}

func (b *Bot) exitBattleField(fromQQ uint64) string {
	b.updateBattleField()
	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)
	for i, v := range bf.Nightwatches {
		if v.QQ == fromQQ {
			if len(bf.Nightwatches) > 1 {
				bf.Nightwatches[i] = bf.Nightwatches[len(bf.Nightwatches)-1]
				bf.Nightwatches = bf.Nightwatches[:len(bf.Nightwatches)-1]
			} else {
				bf.Nightwatches = nil
			}

			break
		}
	}

	for i, v := range bf.MoneyHunter {
		if v.QQ == fromQQ {
			if len(bf.MoneyHunter) > 1 {
				bf.MoneyHunter[i] = bf.MoneyHunter[len(bf.MoneyHunter)-1]
				bf.MoneyHunter = bf.MoneyHunter[:len(bf.MoneyHunter)-1]
			} else {
				bf.MoneyHunter = nil
			}

			break
		}
	}
	b.setGroupValue("BattleField", bf)

	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{fromQQ, b.CurrentNick, 0, 0, 0, 0, uint64(time.Now().Unix()), 0, 15, 0, 0, 0}).(*BattleInfo)
	info := fmt.Sprintf("你成功停止了这份危险的工作，从战场上退了下来，获得%d金镑，%d经验", bi.Money/10000, bi.Exp/10000)
	b.setMoney(fromQQ, int(bi.Money/10000))
	b.setExp(fromQQ, int(bi.Exp/10000))
	bi.Money = 0
	bi.Exp = 0
	bi.State = 2
	bi.ExitTime = uint64(time.Now().Unix() / 3600 / 24)
	b.setPersonValue("BattleInfo", fromQQ, bi)
	return info
}

func (b *Bot) pk(fromQQ uint64, msg string) string {
	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return fmt.Sprintf("指令格式错误:%+v", strs)
	}

	aim, err := strconv.Atoi(strs[1])
	if err != nil {
		return fmt.Sprintf("指令格式错误:%+v", strs)
	}

	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{}).(*BattleInfo)
	if bi.QQ == 0 || bi.State == 2 {
		return "请先加入阵营"
	}

	if bi.State == 1 {
		return "对不起，你正在休整之中，不能挑战别人。"
	}

	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)

	if bi.FieldType == 0 && aim >= len(bf.MoneyHunter) {
		return fmt.Sprintf("找不到这个人:%+v", strs)
	}

	if bi.FieldType == 1 && aim >= len(bf.Nightwatches) {
		return fmt.Sprintf("找不到这个人:%+v", strs)
	}

	aimInfo := &BattleInfo{}
	if bi.FieldType == 0 {
		aimInfo = bf.MoneyHunter[aim]
	} else {
		aimInfo = bf.Nightwatches[aim]
	}

	if aimInfo.State == 1 {
		return "对不起，对方正在休整之中，不能接受挑战。"
	}

	aimQQ := aimInfo.QQ

	exp, skill, item, rp, speed, pt, total := b.getBattleScore(fromQQ)
	exp1, skill1, item1, rp1, speed1, pt1, total1 := b.getBattleScore(aimQQ)
	resultStr := "平局。收获："
	moneyNum := int64(0)
	expNum := int64(0)

	if total > total1 {
		resultStr = "战斗胜利！!收获："
		moneyNum, expNum = b.battleFailed(aimQQ)
		b.battleSuccess(fromQQ, moneyNum, expNum)
	}

	if total < total1 {
		resultStr = "战斗失败...损失："
		moneyNum, expNum = b.battleFailed(fromQQ)
		b.battleSuccess(aimQQ, moneyNum, expNum)
	}

	info := fmt.Sprintf(`
============================
战斗结算：经验/技能/装备/人品/答题/宠物
%s：%d/%d/%d/%d/%d/%d 总计：%d
%s：%d/%d/%d/%d/%d/%d 总计：%d
%s%d金镑，%d经验
============================`,
		bi.Nick, exp, skill, item, rp, speed, pt, total,
		aimInfo.Nick, exp1, skill1, item1, rp1, speed1, pt1, total1, resultStr, moneyNum/10000, expNum/10000)

	return info
}

func (b *Bot) getBattleScore(fromQQ uint64) (exp, skill, item, rp, speed, pt, total int64) {
	rp = int64(rand.Intn(10))
	speed = b.getPlayerSpeed(fromQQ)
	p := b.getPersonFromDb(fromQQ)
	exp = int64(p.ChatCount) / 1000
	if exp > 30 || exp < 0 {
		exp = 30
	}

	skillTree := b.getPersonValue("SkillTree", fromQQ, &SkillTree{}).(*SkillTree)
	church := b.getPersonValue("Church", fromQQ, &ChurchInfo{}).(*ChurchInfo)
	skill = int64(len(skillTree.Skills)*3 + len(church.Skills)*3)
	if skill > 20 || skill < 0 {
		skill = 20
	}

	bag := b.getPersonValue("Bag", fromQQ, &Bag{}).(*Bag)
	for _, v := range bag.Items {
		if v.Name == "左轮手枪" {
			item += 4
		}
		if v.Name == "精致礼帽" {
			item += 4
		}
		if v.Name == "勇者护盾" {
			item += 4
		}
		if v.Name == "旅行者靴子" {
			item += 4
		}
		if v.Name == "旅行者手链" {
			item += 4
		}
	}

	ptv := b.getPersonValue("Pet", fromQQ, &pet.Pet{}).(*pet.Pet)
	pt = int64(ptv.Level) / 2

	total = exp + skill + item + rp + speed + pt
	return
}

func (b *Bot) getPlayerSpeed(fromQQ uint64) int64 {
	ms := b.getPersonValue("Calc", fromQQ, &CalcState{false, nil}).(*CalcState)
	fmt.Printf("speed:%+v", ms)

	if ms.Calc == nil {
		return 0
	}

	if ms.Calc.Speed > 20 || ms.Calc.Speed == 0 {
		ms.Calc.Speed = 20
	}

	s1 := int64(20-ms.Calc.Speed) / 2
	s2 := uint64(0)

	qs := b.getPersonValue("Competition", fromQQ, &CompetitionState{}).(*CompetitionState)
	max := b.getGroupValue("MaxVictory", &MaxVictory{}).(*MaxVictory)

	s2 = qs.MaxVictoryCnt * 10 / max.VictoryCnt

	return s1 + int64(s2)
}

func (b *Bot) battleFailed(fromQQ uint64) (money, exp int64) {
	b.updateBattleField()
	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{}).(*BattleInfo)
	money = int64(bi.Money) / 2
	bi.Money = bi.Money / 2
	exp = int64(bi.Exp) / 2
	bi.Exp = bi.Exp / 2
	bi.FailCnt++
	bi.State = 1
	bi.FailTime = uint64(time.Now().Unix())
	b.setPersonValue("BattleInfo", fromQQ, bi)

	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)
	if bi.FieldType == 0 {
		for i, v := range bf.Nightwatches {
			if v.QQ == fromQQ {
				bf.Nightwatches[i] = bi
				break
			}
		}
	} else {
		for i, v := range bf.MoneyHunter {
			if v.QQ == fromQQ {
				bf.MoneyHunter[i] = bi
				break
			}
		}
	}
	b.setGroupValue("BattleField", bf)
	b.updateBattleField()
	return
}

func (b *Bot) battleSuccess(fromQQ uint64, money, exp int64) {
	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{}).(*BattleInfo)
	bi.Money += uint64(money)
	bi.Exp += uint64(exp)
	bi.WinCnt++
	b.setPersonValue("BattleInfo", fromQQ, bi)

	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)
	if bi.FieldType == 0 {
		for i, v := range bf.Nightwatches {
			if v.QQ == fromQQ {
				bf.Nightwatches[i] = bi
				break
			}
		}
		bf.NightwatchWinCnt++
		bf.MoneyHunterLoseCnt++
	} else {
		for i, v := range bf.MoneyHunter {
			if v.QQ == fromQQ {
				bf.MoneyHunter[i] = bi
				break
			}
		}
		bf.MoneyHunterWinCnt++
		bf.NightwatchLoseCnt++
	}
	b.setGroupValue("BattleField", bf)
	b.updateBattleField()
}
