package secret

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/molin0000/secretMaster/rlp"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func (b *Bot) setRespectName(msg string, fromQQ uint64) string {
	p := b.getPersonFromDb(fromQQ)

	if p.SecretLevel < 6 || p.SecretLevel > 10 {
		return "只有序列3以上才可拥有尊名，因为随便的尊名可能会引起高位存在的注意"
	}

	rname := msg[strings.Index(msg, "尊名")+len("尊名"):]

	b.setRNameToDb(fromQQ, &RespectName{Group: b.Group, QQ: fromQQ, Name: rname})

	return "成功设置尊名"
}

func (b *Bot) deletePerson(fromQQ uint64) string {

	k := b.getPersonValue("KillSelf", fromQQ, &KillSelf{}).(*KillSelf)
	if k.DayCnt >= 3 && k.Days == uint64(time.Now().Day()) {
		return "你今天已经自杀太多次了。"
	}

	if k.Days != uint64(time.Now().Day()) {
		k.DayCnt = 0
		k.Days = uint64(time.Now().Day())
	}

	p := b.getPersonFromDb(fromQQ)

	if b.getGodFromDb(p.SecretID) == p.QQ {
		getDb().Delete(b.godKey(p.SecretID), nil)
	}

	getDb().Delete(b.keys(fromQQ), nil)
	getDb().Delete(b.ruleKey(fromQQ), nil)
	getDb().Delete(b.moneyKey(fromQQ), nil)
	getDb().Delete(b.advKey(fromQQ), nil)
	getDb().Delete(b.externKey(fromQQ), nil)
	getDb().Delete(b.rnameKey(fromQQ), nil)

	getDb().Delete(b.personKey("Bag", fromQQ), nil)
	getDb().Delete(b.personKey("Potion", fromQQ), nil)
	getDb().Delete(b.personKey("SkillTree", fromQQ), nil)
	getDb().Delete(b.personKey("Church", fromQQ), nil)
	getDb().Delete(b.personKey("Pray", fromQQ), nil)
	getDb().Delete(b.personKey("Bank", fromQQ), nil)
	getDb().Delete(b.personKey("Person", fromQQ), nil)
	m := b.getMoney(fromQQ)
	b.setMoney(fromQQ, -1*int(m))
	getDb().Delete(b.personKey("Money", fromQQ), nil)
	getDb().Delete(b.personKey("Water", fromQQ), nil)
	getDb().Delete(b.personKey("Adventure", fromQQ), nil)
	getDb().Delete(b.personKey("Work", fromQQ), nil)
	getDb().Delete(b.personKey("Calc", fromQQ), nil)
	getDb().Delete(b.personKey("Mission", fromQQ), nil)

	b.exitBattleField(fromQQ)
	getDb().Delete(b.personKey("BattleInfo", fromQQ), nil)

	churchs := b.getGroupValue("Churchs", &Churchs{}).(*Churchs)
	for i, c := range churchs.ChurchList {
		if c == nil || c.CreatorQQ == fromQQ {
			if len(churchs.ChurchList) > 1 {
				churchs.ChurchList[i] = churchs.ChurchList[len(churchs.ChurchList)-1]
				churchs.ChurchList = churchs.ChurchList[:len(churchs.ChurchList)-1]
			} else {
				churchs.ChurchList = nil
			}
			b.setGroupValue("Churchs", churchs)
			iter := getDb().NewIterator(util.BytesPrefix(b.getKeyPrefix()), nil)
			for iter.Next() {
				verify := iter.Value()
				var v Person
				rlp.DecodeBytes(verify, &v)
				qq := v.QQ
				cc := b.getPersonValue("Church", qq, &ChurchInfo{}).(*ChurchInfo)
				if cc.Name == c.Name {
					b.removePersonValue("Church", qq)
				}
			}
			iter.Release()
			err := iter.Error()
			if err != nil {
				fmt.Println(err)
			}
			break
		}
	}

	// --remove person in
	p = &Person{
		Group:       b.Group,
		QQ:          fromQQ,
		Name:        b.CurrentNick,
		JoinTime:    uint64(time.Now().Unix()),
		LastChat:    uint64(time.Now().Unix()),
		LevelDown:   uint64(time.Now().Unix()),
		SecretID:    99,
		SecretLevel: 99,
		ChatCount:   1,
	}

	b.setPersonToDb(fromQQ, p)

	k.DayCnt++
	b.setPersonValue("KillSelf", fromQQ, k)

	return "人物删除成功"
}

func (b *Bot) getProperty(fromQQ uint64) string {
	v := b.getPersonFromDb(fromQQ)

	var secretName string
	var secretLevelName string
	var startTime string

	if v.SecretID > 22 {
		secretName = "普通人"
	} else {
		secretName = secretInfo[v.SecretID].SecretName
	}

	if v.SecretLevel > 10 {
		secretLevelName = "普通人"
	} else {
		secretLevelName = fmt.Sprintf("序列%d：%s", 9-v.SecretLevel, secretInfo[v.SecretID].SecretLevelName[v.SecretLevel])
		// secretLevelName = fmt.Sprintf("序列%d", 9-v.SecretLevel)
	}
	if v.JoinTime == 0 {
		v.JoinTime = uint64(time.Now().Unix())
	}
	startTime = fmt.Sprintf("%d小时", (time.Now().Unix()-time.Unix(int64(v.JoinTime), 0).Unix())/3600)
	// startTime = time.Unix(int64(v.JoinTime), 0).Format("2006-01-02 15:04:05")
	exp := b.getExp(fromQQ)
	myFightIndex := exp / 100
	reLive := uint64(0)
	sReLive := ""
	if myFightIndex > 99 {
		myFightIndex = exp / 100 % 100
		reLive = exp / uint64(10000)
		if reLive > 0 {
			sReLive = fmt.Sprintf("(转生+%d)", reLive)
		}
	}

	cc := b.getPersonValue("Church", fromQQ, &ChurchInfo{}).(*ChurchInfo)
	if len(cc.Name) == 0 {
		cc.Name = "无"
	}

	w := b.getPersonValue("Work", fromQQ, &Work{}).(*Work)
	workStr := "无业游民"
	if len(w.Name) > 0 {
		workStr = w.Name
	}

	if len(v.Name) == 0 {
		v.Name = b.CurrentNick
	}

	medalStr := ""
	medalCnt := b.getMedal(fromQQ)
	for md := uint64(0); md < medalCnt; md++ {
		medalStr += "🎖"
	}
	if len(medalStr) == 0 {
		medalStr = "无"
	}

	info := ""
	info = fmt.Sprintf("\n昵称：%s\n途径：%s\n序列：%s\n勋章：%s\n经验：%d\n金镑：%d\n幸运：%d\n灵性：%d\n修炼时间：%s\n战力评价：%s%s\n教会/组织：%s\n工作：%s\n尊名：%s",
		v.Name, secretName, secretLevelName,
		medalStr,
		int64(b.getExp(fromQQ)), int64(b.getMoney(fromQQ)),
		int64(b.getLuck(fromQQ)),
		int64(b.getMagic(fromQQ)),
		startTime, FightLevel[myFightIndex], sReLive,
		cc.Name,
		workStr,
		b.getRNameFromDb(fromQQ),
	)

	return info
}

func (b *Bot) getRank(fromQQ uint64) string {
	iter := getDb().NewIterator(util.BytesPrefix(b.getKeyPrefix()), nil)
	persons := make([]Person, 0)
	cnt := 0
	for iter.Next() {
		verify := iter.Value()
		var v Person
		rlp.DecodeBytes(verify, &v)
		persons = append(persons, v)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println(err)
	}

	retValue := ""

	sort.Sort(Persons(persons))

	b.Rank = make([]uint64, 0)
	b.RankNames = make([]string, 0)
	for i := 0; i < len(persons); i++ {
		v := persons[i]
		retValue = fmt.Sprintf("%s\n第%d名：%s，经验：%d", retValue, i+1, v.Name, v.ChatCount)
		b.Rank = append(b.Rank, v.QQ)
		b.RankNames = append(b.RankNames, v.Name)
		cnt++
		if cnt > 30 {
			break
		}
	}

	b.setMoney(fromQQ, -2)
	return retValue
}

func (b *Bot) getItems(fromQQ uint64) string {
	bag := b.getPersonValue("Bag", fromQQ, &Bag{})
	if len(bag.(*Bag).Items) == 0 {
		return "你的背包空空如也，干净的没有一丝尘土。"
	}

	info := "\n"
	for i := 0; i < len(bag.(*Bag).Items); i++ {
		info += fmt.Sprintf("%s x%d; ", bag.(*Bag).Items[i].Name, bag.(*Bag).Items[i].Count)
	}
	return info
}

func (b *Bot) setExp(fromQQ uint64, v int) {
	person := b.getPersonFromDb(fromQQ)
	if v >= 0 {
		person.ChatCount += uint64(v)
	} else {
		if person.ChatCount > uint64(-1*v) {
			person.ChatCount -= uint64(-1 * v)
		} else {
			person.ChatCount = 0
		}
	}
	b.setPersonToDb(fromQQ, person)
}

func (b *Bot) getExp(fromQQ uint64) uint64 {
	person := b.getPersonFromDb(fromQQ)
	return person.ChatCount
}

func (b *Bot) setMoney(fromQQ uint64, v int) {
	money := b.getMoneyFromDb(fromQQ, 0)
	if v >= 0 {
		money.Money += uint64(v)
	} else {
		if money.Money > uint64(-1*v) {
			money.Money -= uint64(-1 * v)
		} else {
			money.Money = 0
		}
	}
	b.setMoneyToDb(fromQQ, money)
}

func (b *Bot) getMoney(fromQQ uint64) uint64 {
	money := b.getMoneyFromDb(fromQQ, 0)
	return money.Money
}

func (b *Bot) setMagic(fromQQ uint64, v int) {
	e := b.getExternFromDb(fromQQ)
	if v >= 0 {
		e.Magic += uint64(v)
	} else {
		if e.Magic > uint64(-1*v) {
			e.Magic -= uint64(-1 * v)
		} else {
			e.Magic = 0
		}
	}
	b.setExternToDb(fromQQ, e)
}

func (b *Bot) setLuck(fromQQ uint64, v int) {
	e := b.getExternFromDb(fromQQ)
	if v >= 0 {
		e.BaseLuck += uint64(v)
		if e.BaseLuck > 20 {
			e.BaseLuck = 20
		}
	} else {
		if e.BaseLuck > uint64(-1*v) {
			e.BaseLuck -= uint64(-1 * v)
		} else {
			e.BaseLuck = 0
		}
	}
	b.setExternToDb(fromQQ, e)
}

func (b *Bot) getLuck(fromQQ uint64) uint64 {
	e := b.getExternFromDb(fromQQ)
	if int64(e.Luck) < 0 || e.Luck > 20 {
		e.Luck = 0
		b.setExternToDb(fromQQ, e)
	}
	return e.Luck
}

func (b *Bot) getMagic(fromQQ uint64) uint64 {
	e := b.getExternFromDb(fromQQ)
	if int64(e.Magic) < 0 {
		e.Magic = 200
		b.setExternToDb(fromQQ, e)
	}
	return e.Magic
}

func (b *Bot) getGod() string {
	info := "\n"
	for i := 0; i < 22; i++ {
		god := b.getGodFromDb(uint64(i))
		godName := ""
		if god == 0 {
			godName = "空"
		} else {
			p := b.getPersonFromDb(god)
			godName = p.Name
		}

		info += fmt.Sprintf("%d - %s: %s\n", i+1, secretInfo[i].SecretName, godName)
	}
	return info
}

func (b *Bot) getSkill(fromQQ uint64) string {
	tree := b.getPersonValue("SkillTree", fromQQ, &SkillTree{}).(*SkillTree)
	if len(tree.Skills) == 0 {
		return "\n你没有任何技能，努力吧少年(少女)。\n"
	}

	info := "\n"
	for i := 0; i < len(tree.Skills); i++ {
		info += fmt.Sprintf("%s lv%d; ", tree.Skills[i].Name, tree.Skills[i].Level)
	}
	return info
}

func (b *Bot) setMedal(fromQQ uint64, v int) {
	m := GetGlobalPersonValue("Medal", fromQQ, &Medal{}).(*Medal)
	if v >= 0 {
		m.MedalCnt++
	} else {
		if m.MedalCnt > uint64(-1*v) {
			m.MedalCnt -= uint64(-1 * v)
		} else {
			m.MedalCnt = 0
		}
	}
	SetGlobalPersonValue("Medal", fromQQ, m)
}

func (b *Bot) getMedal(fromQQ uint64) uint64 {
	m := GetGlobalPersonValue("Medal", fromQQ, &Medal{}).(*Medal)
	return m.MedalCnt
}
