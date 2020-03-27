package secret

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/molin0000/secretMaster/rlp"
	"github.com/syndtr/goleveldb/leveldb"
)

var botMap map[uint64]*Bot

func init() {
	fmt.Println("加载技能表格...")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("加载技能表格...失败")
			fmt.Println(err)
		}
	}()

	if !fileExists(careerSkillPath) {
		fmt.Println("文件不存在")
	}

	f, err := excelize.OpenFile(careerSkillPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows := f.GetRows("技能列表")
	for i, row := range rows {
		if i == 0 {
			continue
		}

		cs := &CareerSkill{
			Name:  row[3],
			Times: atoi(row[4]),
			Type:  row[5],
			Desc:  row[6],
		}

		careerSkills = append(careerSkills, cs)
	}

	fmt.Printf("共加载技能：%d\n", len(careerSkills))
}

func atoi(msg string) uint64 {
	n, _ := strconv.Atoi(msg)
	return uint64(n)
}

func NewSecretBot(qq, group uint64, groupNick string, private bool, api interface{}) *Bot {
	setCqpCall(api.(CqpCall))
	if botMap == nil {
		botMap = make(map[uint64]*Bot)
	}
	bot, ok := botMap[group]
	if ok {
		bot.Private = private
		return bot
	}
	bot = &Bot{QQ: qq, Group: group, Name: groupNick, Private: private}
	bot.information = &Information{bot}
	bot.career = &Career{bot}
	bot.organization = &Organization{bot}
	bot.store = &Store{bot}

	botMap[group] = bot
	return bot
}

func (b *Bot) Run(msg string, fromQQ uint64, nick string) string {
	if !b.talkToMe(msg) {
		return ""
	}
	gp := b.getGroupValue("GroupMap", &GroupMap{}).(*GroupMap)
	if gp.Mapped {
		b.Group = gp.MapToGroup
	}

	b.CurrentNick = nick

	ret := b.checkMission(fromQQ, msg)
	if len(ret) > 0 {
		return ret
	}

	ret = b.checkCalc(fromQQ, msg)
	if len(ret) > 0 {
		return ret
	}

	ret = b.checkCompetion(fromQQ, msg)
	if len(ret) > 0 {
		return ret
	}

	return b.searchMenu(msg, fromQQ, &menus)
}

func (b *Bot) RunPrivate(msg string, fromQQ uint64, nick string) string {
	gp := b.getGroupValue("GroupMap", &GroupMap{}).(*GroupMap)
	if gp.Mapped {
		b.Group = gp.MapToGroup
	}

	b.CurrentNick = nick
	ret := b.checkMission(fromQQ, msg)
	if len(ret) > 0 {
		return ret
	}

	ret = b.checkCalc(fromQQ, msg)
	if len(ret) > 0 {
		return ret
	}

	ret = b.checkCompetion(fromQQ, msg)
	if len(ret) > 0 {
		return ret
	}

	return b.searchMenu(msg, fromQQ, &menus)
}

func (b *Bot) UpdateFromOldVersion(fromQQ uint64) string {
	up := b.getPersonValue("Update", fromQQ, &DbUpdate{}).(*DbUpdate)
	if up.HasUpdate {
		return ""
	}
	info := ""
	if dirExists("secret.db") {
		fmt.Println("老版本数据库已找到，准备升级", fromQQ)
		_db, err := leveldb.OpenFile("secret.db", nil)
		if err != nil {
			fmt.Printf("open db error: %+v", err)
			return ""
		}
		defer _db.Close()
		verify, _ := _db.Get(b.keys(fromQQ), nil)
		var v Person
		rlp.DecodeBytes(verify, &v)
		p := b.getPersonFromDb(fromQQ)
		p.ChatCount = v.ChatCount
		b.setPersonToDb(fromQQ, p)

		m, _ := _db.Get(b.moneyKey(fromQQ), nil)
		var money Money
		rlp.DecodeBytes(m, &money)
		b.setMoney(fromQQ, int(money.Money))
		fmt.Println("继承经验:", p.ChatCount, "继承金钱:", money.Money)
		info += fmt.Sprintf("\n继承经验:%d, 继承金钱:%d\n", p.ChatCount, money.Money)
	} else {
		fmt.Println("老版本数据库不存在")
	}
	up.HasUpdate = true
	b.setPersonValue("Update", fromQQ, up)
	fmt.Println("升级完成")
	return info
}

func (b *Bot) Update(fromQQ uint64, nick string) string {
	if !b.getSwitch() {
		return ""
	}

	gp := b.getGroupValue("GroupMap", &GroupMap{}).(*GroupMap)
	if gp.Mapped {
		b.Group = gp.MapToGroup
	}

	key := b.personKey("Person", fromQQ)
	value, err := getDb().Get(key, nil)
	fmt.Println("value:", value)
	fmt.Println("err:", err)
	ret := ""
	if err != nil {
		if err.Error() == "leveldb: not found" {
			fmt.Println("a new man.")
			p := &Person{
				Group:       b.Group,
				QQ:          fromQQ,
				Name:        nick,
				JoinTime:    uint64(time.Now().Unix()),
				LastChat:    uint64(time.Now().Unix()),
				LevelDown:   uint64(time.Now().Unix()),
				SecretID:    99,
				SecretLevel: 99,
				ChatCount:   1,
			}

			b.setPersonToDb(fromQQ, p)

			e := b.getExternFromDb(fromQQ)
			b.setExternToDb(fromQQ, e)
		}
	} else {
		ret += b.UpdateFromOldVersion(fromQQ)

		v := b.getPersonFromDb(fromQQ)
		if v.Name != nick && len(nick) > 0 {
			v.Name = nick
		}

		magic := b.getMagic(fromQQ)

		if magic > 0 {
			b.setMagic(fromQQ, -1)
		}

		normalHumanStop := false
		if v.SecretID > 22 && v.ChatCount > 400 {
			normalHumanStop = true
		}

		if int64(magic) > 0 && !normalHumanStop {
			v.ChatCount++
			b.setMoney(fromQQ, 1)
			if v.ChatCount%100 == 0 {
				ret += "\n恭喜！你的战力评价升级了！"
			}
		}

		v.LastChat = uint64(time.Now().Unix())
		v.LevelDown = uint64(time.Now().Unix())
		b.setPersonToDb(fromQQ, v)
	}

	return ret
}

func (b *Bot) printMenu(menu *Menu) string {
	if menu == nil {
		return ""
	}

	if menu.ID == 7 && !b.Private {
		return ""
	}

	if !b.getSwitch() {
		return ""
	}

	info := fmt.Sprintf("\n%s: %s \n", menu.Title, menu.Info)
	if menu.SubMenu != nil && len(menu.SubMenu) > 0 {
		for _, mu := range menu.SubMenu {
			if mu.ID == 7 && !b.Private {
				continue
			}

			info += fmt.Sprintf("%s: %s \n", mu.Title, mu.Info)
		}
	}
	if len(menu.Commit) > 0 {
		info += menu.Commit
	}
	return info
}

func (b *Bot) searchMenu(msg string, fromQQ uint64, menu *Menu) string {
	if strings.Contains(msg, menu.Title) {
		if menu.SubMenu != nil && len(menu.SubMenu) > 0 {
			return b.printMenu(menu)
		}
		return b.cmdRun(msg, fromQQ)
	}

	if menu.SubMenu != nil && len(menu.SubMenu) > 0 {
		for _, mu := range menu.SubMenu {
			info := b.searchMenu(msg, fromQQ, &mu)
			if len(info) > 0 {
				return info
			}
		}
	}
	return ""
}

func (b *Bot) cmdRun(msg string, fromQQ uint64) string {
	fmt.Println("发现指令触发")

	msg = strings.ReplaceAll(msg, "；", ";")

	if strings.Contains(msg, "序列战争关") {
		return b.botSwitch(fromQQ, false)
	}

	if strings.Contains(msg, "序列战争开") {
		return b.botSwitch(fromQQ, true)
	}

	if !b.getSwitch() {
		return ""
	}

	if strings.Contains(msg, "探险卷轴") {
		return b.adventure(fromQQ, false)
	}

	if strings.Contains(msg, "红剧场门票") {
		return b.redTheater(fromQQ)
	}

	if strings.Contains(msg, "属性") {
		return b.getProperty(fromQQ)
	}

	if strings.Contains(msg, "查询") {
		rankStr := msg[strings.Index(msg, "查询")+len("查询"):]
		fmt.Println("查询排名：", rankStr)
		rank, err := strconv.Atoi(rankStr)
		if err != nil {
			return "是找我吗？查询排名要查询加数字哦。如果不是，请不要艾特我。"
		}
		if len(b.Rank) > rank-1 {
			return b.getProperty(b.Rank[rank-1])
		}
		return "请先查看最新排行"
	}

	if strings.Contains(msg, "途径") {
		return b.getSecretList()
	}

	if strings.Contains(msg, "更换") {
		return b.changeSecretList(msg, fromQQ)
	}

	if strings.Contains(msg, "排行") {
		return b.getRank(fromQQ)
	}

	if strings.Contains(msg, "探险") {
		return b.adventure(fromQQ, true)
	}

	if strings.Contains(msg, "自杀") {
		return b.deletePerson(fromQQ)
	}

	if strings.Contains(msg, "尊名") {
		return b.setRespectName(msg, fromQQ)
	}

	if strings.Contains(msg, "GM") {
		return b.gmCmd(fromQQ, msg)
	}

	if strings.Contains(msg, "货币升级") {
		return b.moneyUpdate(fromQQ, true)
	}

	if strings.Contains(msg, "货币降级") {
		return b.moneyUpdate(fromQQ, false)
	}

	if strings.Contains(msg, "货币映射") {
		return b.moneyMap(fromQQ, msg)
	}

	if strings.Contains(msg, "查看映射") {
		bind := b.getMoneyBind()
		return fmt.Sprintf("%+v\n", bind)
	}

	if strings.Contains(msg, ".master") {
		return b.setMaster(fromQQ, msg)
	}

	if strings.Contains(msg, ".supermaster") {
		return b.setSuperMaster(fromQQ, msg)
	}

	if strings.Contains(msg, "道具") {
		return b.getItems(fromQQ)
	}

	if strings.Contains(msg, "神位") {
		return b.getGod()
	}

	if strings.Contains(msg, "灵性药剂") {
		return b.buyMagicPotion(fromQQ)
	}

	if strings.Contains(msg, "灵性材料") {
		return b.buyMagicItem(fromQQ)
	}

	if strings.Contains(msg, "至高权杖") {
		return b.buyMace(fromQQ)
	}

	if strings.Contains(msg, "左轮手枪") {
		return b.buyItem(fromQQ, "左轮手枪", 200)
	}

	if strings.Contains(msg, "精致礼帽") {
		return b.buyItem(fromQQ, "精致礼帽", 100)
	}

	if strings.Contains(msg, "勇者护盾") {
		return b.buyItem(fromQQ, "勇者护盾", 200)
	}

	if strings.Contains(msg, "旅行者靴子") {
		return b.buyItem(fromQQ, "旅行者靴子", 200)
	}

	if strings.Contains(msg, "旅行者手链") {
		return b.buyItem(fromQQ, "旅行者手链", 200)
	}

	if strings.Contains(msg, "技能") {
		return b.getSkill(fromQQ)
	}

	if strings.Contains(msg, "晋升") {
		return b.promotion(fromQQ)
	}

	if strings.Contains(msg, "创建") {
		return b.createChurch(fromQQ, msg)
	}

	if strings.Contains(msg, "解散") {
		return b.deleteChurch(fromQQ, msg)
	}

	if strings.Contains(msg, "寻访") {
		return b.listChurch()
	}

	if strings.Contains(msg, "版本") {
		return b.getVersion()
	}

	if strings.Contains(msg, "加入") {
		return b.joinChurch(fromQQ, msg)
	}

	if strings.Contains(msg, "退出") {
		return b.exitChurch(fromQQ, msg)
	}

	if strings.Contains(msg, "祈祷") {
		return b.pray(fromQQ)
	}

	if strings.Contains(msg, "银行") {
		return b.bank(fromQQ, msg)
	}

	if strings.Contains(msg, "工作") {
		return b.work(fromQQ, msg)
	}

	if strings.Contains(msg, "钓鱼") {
		return b.fishing(fromQQ)
	}

	if strings.Contains(msg, "许愿") {
		return b.lottery(fromQQ)
	}

	if strings.Contains(msg, "速算") {
		return b.startCalc(fromQQ, msg)
	}

	if strings.Contains(msg, "学识") {
		return b.startCompetition(fromQQ)
	}

	if strings.Contains(msg, "副本") {
		if strings.Contains(msg, "列表") {
			return b.showMissions()
		}
		return b.startMission(fromQQ, msg)
	}

	if strings.Contains(msg, "silent") {
		return b.SetSilentTime(fromQQ, msg)
	}

	if strings.Contains(msg, "阵营") {
		fmt.Println("阵营触发")
		return b.getBattleField()
	}

	if strings.Contains(msg, "挑战") {
		return b.pk(fromQQ, msg)
	}

	if strings.Contains(msg, "战绩") {
		return b.getBattleInfo(fromQQ)
	}

	if strings.Contains(msg, "紫包") {
		return b.redPack(fromQQ, msg)
	}

	if strings.Contains(msg, "map") {
		return b.groupMap(fromQQ, msg)
	}

	if strings.Contains(msg, "宠物状态") {
		return b.petState(fromQQ)
	}

	if strings.Contains(msg, "宠物改名") {
		return b.petRename(fromQQ, msg)
	}

	if strings.Contains(msg, "宠物货架") {
		return b.petStore()
	}

	if strings.Contains(msg, "宠物领养") {
		return b.petBuy(fromQQ, msg)
	}

	if strings.Contains(msg, "灵界法阵") {
		return b.petSummon(fromQQ)
	}

	if strings.Contains(msg, "宠物晋级") {
		return b.petLevelUp(fromQQ)
	}

	if strings.Contains(msg, "宠物放生") {
		return b.petSell(fromQQ)
	}

	if strings.Contains(msg, "宠物融合") {
		return b.petEat(fromQQ)
	}

	if strings.Contains(msg, "宠物历练") {
		return b.petAdventure(fromQQ)
	}

	if strings.Contains(msg, "宠物召回") {
		return b.petGoHome(fromQQ)
	}

	if strings.Contains(msg, "宠物美妆") {
		return b.petBeauty(fromQQ)
	}

	if strings.Contains(msg, "delay") {
		return b.setDelay(fromQQ, msg)
	}

	if strings.Contains(msg, "数值修复") {
		return b.fixNumber(fromQQ)
	}
	return ""
}

func (b *Bot) talkToMe(msg string) bool {
	if len(msg) == 0 {
		return false
	}

	cp := fmt.Sprintf("CQ:at,qq=%d", b.QQ)

	if strings.Index(msg, cp) != -1 {
		return true
	}

	return false
}

func (b *Bot) getVersion() string {
	v := b.getGroupValue("Version", &Version{"", "", ""}).(*Version)
	if v.Version == "" {
		//Update

	}

	return fmt.Sprintf("\n%s %s %s", version.Name, version.Version, version.Date)
}
