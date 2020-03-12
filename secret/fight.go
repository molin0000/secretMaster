package secret

import (
	"fmt"
	"math/rand"
	"time"
)

func (b *Bot) getBattleField() string {
	b.updateBattleField()

	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)
	info := "\n==================\n值夜者小队阵营："

	for i, v := range bf.Nightwatches {
		state := ""
		if v.State == 0 {
			state = "正常"
		} else {
			state = "休整"
		}
		info += fmt.Sprintf("\n%d) %s (%s) 金镑%d, 经验%d", i, v.Nick, state, v.Money, v.Exp)
	}

	info += "\n==================\n赏金猎人阵营："
	for i, v := range bf.MoneyHunter {
		state := ""
		if v.State == 0 {
			state = "正常"
		} else {
			state = "休整"
		}
		info += fmt.Sprintf("\n%d) %s (%s) 金镑%d, 经验%d", i, v.Nick, state, v.Money, v.Exp)
	}
	info += "\n=================="

	return ""
}

func (b *Bot) updateBattleField() {
	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)
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
		b.setPersonValue("BattleInfo", v.QQ, bf.MoneyHunter[i])
	}

	b.setGroupValue("BattleField", bf)
}

func (b *Bot) joinBattleField(fromQQ uint64, fieldType int) (string, bool) {
	p := b.getPersonFromDb(fromQQ)

	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{fromQQ, p.Name, 0, 0, 0, 0, uint64(time.Now().Unix()), 0, 15, 0, 0}).(*BattleInfo)
	if bi.ExitTime != 0 {
		if uint64(time.Now().Unix()/3600/24)-bi.ExitTime == 0 {
			return "今天加入阵营次数已经用尽，请明天再试。", false
		}
	}

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
	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{fromQQ, p.Name, 0, 0, 0, 0, uint64(time.Now().Unix()), 0, 15, 0, 0}).(*BattleInfo)
	return fmt.Sprintf("\n===============\n个人战绩: 胜%d次，败%d次\n当前收入：%d金镑, %d经验\n===============",
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

	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{fromQQ, b.CurrentNick, 0, 0, 0, 0, uint64(time.Now().Unix()), 0, 15, 0, 0}).(*BattleInfo)
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

	return `============================
战斗结算：经验/技能/装备/人品/答题
空想之喵：15/5/15/10/5 总计：50
敌人XXX：3/10/12/15/5 总计：45
战斗胜利！！收获：100金镑，20经验
============================`
}

func (b *Bot) getBattleScore(fromQQ uint64) (exp, skill, item, rp, speed, total int) {
	rp = rand.Intn(10)
	bi := b.getPersonValue("BattleInfo", fromQQ, &BattleInfo{fromQQ, b.CurrentNick, 0, 0, 0, 0, uint64(time.Now().Unix()), 0, 15, 0, 0}).(*BattleInfo)
	speed = int(bi.Speed)
	p := b.getPersonFromDb(fromQQ)
	exp = int(p.ChatCount) / 1000
	if exp > 30 || exp < 0 {
		exp = 30
	}

	skillTree := b.getPersonValue("SkillTree", fromQQ, &SkillTree{}).(*SkillTree)
	church := b.getPersonValue("Church", fromQQ, &ChurchInfo{}).(*ChurchInfo)
	skill = len(skillTree.Skills)*3 + len(church.Skills)*3
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

	total = exp + skill + item + rp + speed
	return
}

func (b *Bot) getPlayerSpeed(fromQQ uint64) {

}
