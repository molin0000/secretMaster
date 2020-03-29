package secret

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (b *Bot) createChurch(fromQQ uint64, msg string) string {
	strs := strings.Split(msg, ";")
	if len(strs) != 5 {
		return "\n指令格式不正确！别想欺骗机器人！" + fmt.Sprintf("%+v", strs)
	}

	enterPrice, err := strconv.Atoi(strs[4])

	if err != nil {
		return "\n指令格式不正确！别想欺骗机器人！" + fmt.Sprintf("%+v %+v", strs, err)
	}

	p := b.getPersonFromDb(fromQQ)

	if p.SecretLevel < 5 {
		return "\n你的能力和声望不足以服众，没人愿意听你说什么。"
	}

	if b.getMoney(fromQQ) < 1000 {
		return "\n你查了查贝克兰德的地价，发现钱似乎还不够。"
	}

	newChurch := &ChurchInfo{strs[1], strs[2], nil, fromQQ, b.CurrentNick, uint64(enterPrice), b.getMoney(fromQQ) / 200, 1, b.getMoney(fromQQ), 1}

	info := "\n"

	churchs := b.getGroupValue("Churchs", &Churchs{}).(*Churchs)

	for _, church := range churchs.ChurchList {
		if church.CreatorQQ == newChurch.CreatorQQ || church.Name == newChurch.Name {
			return "请不要重复创建。"
		}
	}

	findSkill := false
	for _, skill := range skillList {
		if strs[3] == skill.Name {
			newChurch.Skills = append(newChurch.Skills, &skill)
			findSkill = true
			break
		}
	}

	if !findSkill {
		return "\n抱歉，没听说过这个技能。"
	}

	if !b.useItem(fromQQ, "至高权杖") {
		return "\n你似乎缺少了什么信物，没法吸引信徒。"
	}

	b.setMoney(fromQQ, -1000)

	churchs.ChurchList = append(churchs.ChurchList, newChurch)
	b.setGroupValue("Churchs", churchs)

	b.setPersonValue("Church", fromQQ, newChurch)

	info += fmt.Sprintf("创建成功！\n名称:%s\n介绍:%s\n尊神/教主:%s\n技能:%s\n入会费:%d\n最大人数:%d\n等级:%d级\n注册资本:%d金镑",
		newChurch.Name, newChurch.Commit, newChurch.CreatorNick, strs[3], newChurch.Money, b.getMoney(fromQQ)/200, newChurch.Level, b.getMoney(fromQQ))

	return info
}

func (b *Bot) deleteChurch(fromQQ uint64, msg string) string {
	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return "格式不正确" + fmt.Sprintf("%+v", strs)
	}

	churchs := b.getGroupValue("Churchs", &Churchs{}).(*Churchs)
	for i, c := range churchs.ChurchList {
		if c == nil || (c.Name == strs[1] && c.CreatorQQ == fromQQ) {
			if len(churchs.ChurchList) > 1 {
				churchs.ChurchList[i] = churchs.ChurchList[len(churchs.ChurchList)-1]
				churchs.ChurchList = churchs.ChurchList[:len(churchs.ChurchList)-1]
			} else {
				churchs.ChurchList = nil
			}
			b.setGroupValue("Churchs", churchs)
		}
	}

	return "解散成功"
}

func (b *Bot) listChurch() string {
	info := ""
	churchs := b.getGroupValue("Churchs", &Churchs{}).(*Churchs)
	if len(churchs.ChurchList) == 0 {
		return "你没有找到任何教会。"
	}

	for i, c := range churchs.ChurchList {
		update := false

		money := b.getMoney(c.CreatorQQ)
		if int64(money)-int64(c.CreateMoney) > 0 {
			c.Level = (money - c.CreateMoney) / 10000
			if c.Level > 3 {
				c.Level = 3
			}

			if c.Level > uint64(len(c.Skills)) {
				add := c.Level - uint64(len(c.Skills))
				for i, s := range skillList {
					have := false
					for _, skill := range c.Skills {
						if s.ID == skill.ID {
							have = true
						}
					}
					if !have {
						c.Skills = append(c.Skills, &skillList[i])
						add--
						if add == 0 {
							break
						}
					}
				}
				update = true
			}
		}
		p := b.getPersonFromDb(c.CreatorQQ)
		skillStr := ""
		for m := 0; m < len(c.Skills); m++ {
			skillStr += fmt.Sprintf("%s lv%d; ", c.Skills[m].Name, p.SecretLevel-4)
			if c.Skills[m].Level < p.SecretLevel-4 {
				c.Skills[m].Level = p.SecretLevel - 4
				update = true
			}
		}

		info += fmt.Sprintf("\n名称:%s\n介绍:%s\n尊神/教主:%s\n技能:%s\n入会费:%d\n人数:%d/%d\n等级:%d级\n注册资本:%d金镑\n",
			c.Name, c.Commit, c.CreatorNick, skillStr, c.Money, int64(c.Members), b.getMoney(c.CreatorQQ)/200, c.Level, c.CreateMoney)

		if update {
			churchs.ChurchList[i] = c
			b.setGroupValue("Churchs", churchs)
		}
	}

	return info
}

func (b *Bot) joinChurch(fromQQ uint64, msg string) string {
	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return "格式不正确" + fmt.Sprintf("%+v", strs)
	}

	cc := b.getPersonValue("Church", fromQQ, &ChurchInfo{}).(*ChurchInfo)
	if cc.Name == strs[1] {
		return "你不是已经加入了么？别开玩笑。"
	}

	if len(cc.Name) > 0 {
		return "还不快走开！你个二五仔！你已经有组织了！"
	}

	churchs := b.getGroupValue("Churchs", &Churchs{}).(*Churchs)
	for i, c := range churchs.ChurchList {
		if c != nil && c.Name == strs[1] {
			if c.Members >= b.getMoney(c.CreatorQQ)/200 {
				return "\n对不起，满员了。"
			}

			if b.getMoney(fromQQ) < c.Money {
				return "\n你囊中羞涩，交不起入会费，被委婉的劝退了。"
			}

			b.setMoney(fromQQ, -1*int(c.Money))
			b.setMoney(c.CreatorQQ, int(c.Money))
			churchs.ChurchList[i].Members++
			b.setPersonValue("Church", fromQQ, c)
			b.setGroupValue("Churchs", churchs)
			return "街头巷尾的流言、古怪神秘的仪式、一支未知的信仰正在悄然兴起。你遵循道听途说的消息找到了这座默默无闻的教堂，衣冠朴素的教士微笑着对你张开双臂：欢迎，我的孩子！欢迎加入" + c.Name + "，在" + c.CreatorNick + "的见证下，我们将步入永恒。"
		}
	}

	return "\n你没有找到这个教会的据点，你确定有这个教会？"
}

func (b *Bot) exitChurch(fromQQ uint64, msg string) string {
	cc := b.getPersonValue("Church", fromQQ, &ChurchInfo{}).(*ChurchInfo)
	if len(cc.Name) == 0 {
		return "\n你并未加入教会/组织，无需退出。"
	}

	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return "\n退出指令格式为：退出;教会名"
	}

	if cc.Name != strs[1] {
		return "\n你的教会名称输入错了。"
	}

	getDb().Delete(b.personKey("Church", fromQQ), nil)

	churchs := b.getGroupValue("Churchs", &Churchs{}).(*Churchs)
	if churchs.ChurchList == nil {
		return "未找到教会组织。"
	}

	for i, v := range churchs.ChurchList {
		if v != nil && v.Name == cc.Name {
			if v.CreatorQQ == fromQQ {
				return "教主不能退出，只能解散。"
			}

			if int64(churchs.ChurchList[i].Members) > 0 {
				churchs.ChurchList[i].Members--
			}
			b.setGroupValue("Churchs", churchs)
			return "你退出了" + cc.Name
		}
	}

	return "你的教会早已消亡。"
}

func (b *Bot) pray(fromQQ uint64) string {
	today := uint64(time.Now().Unix() / (3600 * 24))
	ps := b.getPersonValue("Pray", fromQQ, &PrayState{}).(*PrayState)
	if ps.Date != 0 && (today-ps.Date) < 3 {
		return "你已经处于被庇护状态，无需重复祈祷。"
	}

	cc := b.getPersonValue("Church", fromQQ, &ChurchInfo{}).(*ChurchInfo)
	// if fromQQ == cc.CreatorQQ {
	// 	return "你试着向自己祈祷，并无效果。"
	// }

	if !b.useItem(fromQQ, "灵力材料") {
		return "连灵力材料都没有，瞎祈祷个什么。"
	}

	ps.Date = today
	b.setPersonValue("Pray", fromQQ, ps)

	b.listChurch()

	churchs := b.getGroupValue("Churchs", &Churchs{}).(*Churchs)
	find := false
	for i, v := range churchs.ChurchList {
		if cc.Name == v.Name {
			cc = churchs.ChurchList[i]
			find = true
			break
		}
	}

	if find {
		b.setPersonValue("Church", fromQQ, cc)
		b.setExp(cc.CreatorQQ, 10)
		b.setMagic(fromQQ, int(b.getChurchAdditionInfo(fromQQ, "灵力协调", 50)))
		return "你摆出精心准备的灵力材料，双手合十，认真祈祷……一阵清风拂过，你感觉自己似乎变强了。"
	}

	return "你没有加入教会，请不要胡乱祈祷"
}

func (b *Bot) getAdditionInfo(fromQQ uint64, skill string, baseCount uint64) uint64 {
	return b.getPersonalAdditionInfo(fromQQ, skill, baseCount) + b.getChurchAdditionInfo(fromQQ, skill, baseCount)
}

func (b *Bot) getPersonalAdditionInfo(fromQQ uint64, skill string, baseCount uint64) uint64 {
	addition := uint64(0)
	tree := b.getPersonValue("SkillTree", fromQQ, &SkillTree{}).(*SkillTree)
	if len(tree.Skills) > 0 {
		for _, v := range tree.Skills {
			if v.Name == skill {
				addition += baseCount * v.Level
				break
			}
		}
	}

	return addition
}

func (b *Bot) getChurchAdditionInfo(fromQQ uint64, skill string, baseCount uint64) uint64 {
	addition := uint64(0)
	today := uint64(time.Now().Unix() / (3600 * 24))
	pray := b.getPersonValue("Pray", fromQQ, &PrayState{}).(*PrayState)

	if today-pray.Date < 3 {
		cc := b.getPersonValue("Church", fromQQ, &ChurchInfo{}).(*ChurchInfo)
		if len(cc.Skills) > 0 {
			for _, v := range cc.Skills {
				if v.Name == skill {
					addition += baseCount * v.Level
					break
				}
			}
		}
	}

	return addition
}

func (b *Bot) getAdditionMagic(fromQQ uint64) int {
	m := b.getAdditionInfo(fromQQ, "灵力协调", 50)
	w := b.getPersonValue("Work", fromQQ, &Work{}).(*Work)

	return int(m) - int(w.MagicMinus)
}

func (b *Bot) getAdditionAdventure(fromQQ uint64) uint64 {
	return b.getAdditionInfo(fromQQ, "精力充沛", 1)
}

func (b *Bot) getAdditionLucky(fromQQ uint64) uint64 {
	v := b.getAdditionInfo(fromQQ, "幸运光环", 1)
	return v
}
