package secret

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/molin0000/secretMaster/rlp"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func (b *Bot) setMaster(fromQQ uint64, msg string) string {
	cfg := b.getGroupValue("Config", &Config{})
	if cfg.(*Config).HaveMaster && !cqpCall.IsGroupMaster(b.Group, fromQQ) {
		return "Master已经配置，只有群主可以修改"
	}

	if !cqpCall.IsGroupAdmin(b.Group, fromQQ) {
		return "只有群主或管理员可以设置.master"
	}

	str0 := strings.Split(msg, "@")
	msg = str0[0]

	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return fmt.Sprintf("当前master:%+v", cfg)
	}

	masterQQ, _ := strconv.ParseUint(strs[1], 10, 64)
	cfg.(*Config).MasterQQ = masterQQ
	cfg.(*Config).HaveMaster = true
	if masterQQ == 0 {
		cfg.(*Config).HaveMaster = false
	}
	b.setGroupValue("Config", cfg)
	return fmt.Sprintf("成功设置%d为插件master", masterQQ)
}

func (b *Bot) setSuperMaster(fromQQ uint64, msg string) string {
	cfg := GetGlobalValue("Supermaster", &Config{}).(*Config)

	if cfg.HaveMaster && fromQQ != cfg.MasterQQ {
		return "supermaster只能配置1次，然后只有supermaster可以修改"
	}

	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return fmt.Sprintf("当前supermaster:%+v", cfg)
	}

	masterQQ, _ := strconv.ParseUint(strs[1], 10, 64)
	cfg.MasterQQ = masterQQ
	cfg.HaveMaster = true
	if masterQQ == 0 {
		cfg.HaveMaster = false
	}
	SetGlobalValue("Supermaster", cfg)
	return fmt.Sprintf("成功设置%d为插件supermaster", masterQQ)
}

func (b *Bot) GetMaster() uint64 {
	cfg := b.getGroupValue("Config", &Config{}).(*Config)
	return cfg.MasterQQ
}

func (b *Bot) isMaster(fromQQ uint64) bool {
	cfg := b.getGroupValue("Config", &Config{}).(*Config)
	cfgSuper := GetGlobalValue("Supermaster", &Config{}).(*Config)
	return (cfg.HaveMaster && (fromQQ == cfg.MasterQQ)) || (fromQQ == cfgSuper.MasterQQ)
}

func (b *Bot) isSuperMaster(fromQQ uint64) bool {
	cfgSuper := GetGlobalValue("Supermaster", &Config{}).(*Config)
	return (fromQQ == cfgSuper.MasterQQ)
}

func (b *Bot) notGM() string {
	return "对不起，你不是GM，别想欺骗机器人"
}

func (b *Bot) moneyUpdate(fromQQ uint64, update bool) string {
	if !b.isMaster(fromQQ) {
		return b.notGM()
	}

	bind := b.getMoneyBind()
	if update {
		if bind.HasUpdate {
			return "已经升级过了，请不要重复升级"
		}
		bind.HasUpdate = true
		b.setMoneyBind(bind)
		return "升级成功"
	}

	bind.HasUpdate = false
	b.setMoneyBind(bind)
	return "降级成功，配置保留，但不再读取ini文件"
}

func (b *Bot) moneyMap(fromQQ uint64, msg string) string {
	if !b.isMaster(fromQQ) {
		return b.notGM()
	}

	str0 := strings.Split(msg, "@")
	msg = str0[0]

	strs := strings.Split(msg, ";")
	bind := &MoneyBind{}
	bind.IniPath = strs[1]
	bind.IniSection = strs[2]
	bind.IniKey = strs[3]
	if len(strs) == 5 {
		bind.Encode = strs[4]
	}
	b.setMoneyBind(bind)
	return fmt.Sprintf("映射成功, Path:%s, Section:%s, Key:%s %+v\n", strs[1], strs[2], strs[3], *bind)
}

func (b *Bot) gmCmd(fromQQ uint64, msg string) string {
	if !b.isMaster(fromQQ) {
		return b.notGM()
	}

	str0 := strings.Split(msg, "@")
	msg = str0[0]

	strs := strings.Split(msg, ";")

	n1, err1 := strconv.Atoi(strs[2])
	n2, err2 := strconv.ParseUint(strs[3], 10, 64)

	if err1 != nil || err2 != nil {
		return fmt.Sprintf("参数解析错误: 0:%s, 1:%s, 2:%s, 3:%s, %+v, %+v", strs[0], strs[1], strs[2], strs[3], err1, err2)
	}
	switch strs[1] {
	case "money":
		b.setMoney(n2, n1)
		return fmt.Sprintf("%d 金镑：%d", n2, n1)
	case "exp":
		b.setExp(n2, n1)
		return fmt.Sprintf("%d 经验：%d", n2, n1)
	case "magic":
		b.setMagic(n2, n1)
		return fmt.Sprintf("%d 灵力：%d", n2, n1)
	case "god":
		b.setGodToDb(uint64(n1-1), &n2)
		return fmt.Sprintf("设置途径%d 神灵：%d", n1, n2)
	case "luck":
		b.setLuck(n2, n1)
		return fmt.Sprintf("%d 幸运：%d", n2, n1)
	case "skill":
		for m := 0; m < n1; m++ {
			b.allSkillLevelUp(n2)
		}
		return fmt.Sprintf("%d 所有技能升%d级", n2, n1)
	case "medal":
		cfg := GetGlobalValue("Supermaster", &Config{}).(*Config)
		if fromQQ != cfg.MasterQQ {
			return "只有机器人主人可以颁发勋章🎖"
		}

		b.setMedal(n2, n1)

		return fmt.Sprintf("%d 勋章🎖%d", n2, n1)
	case "level":
		p := b.getPersonFromDb(n2)
		if p.SecretID > 22 {
			return "请先选途径"
		}
		p.SecretLevel = uint64(n1)
		b.setPersonToDb(n2, p)
		return fmt.Sprintf("%d level to: %d", n2, n1)
	case "way":
		p := b.getPersonFromDb(n2)
		if n1 > 0 {
			p.SecretID = uint64(n1 - 1)
			b.setPersonToDb(n2, p)
			return fmt.Sprintf("%d way to: %d", n2, n1)
		}
		return "数值异常"
	case "bank":
		p := b.getPersonValue("Bank", n2, &Bank{}).(*Bank)
		if n1 >= 0 {
			p.Date += uint64(n1)
		} else {
			p.Date -= uint64(-1 * n1)
		}
		b.setPersonValue("Bank", n2, p)
		return fmt.Sprintf("%d bank date to: %d, %+v", n2, n1, p)
	default:
		return "参数解析错误"
	}
}

func (b *Bot) GetSwitch() bool {
	return b.getSwitch()
}

func (b *Bot) botSwitch(fromQQ uint64, enable bool) string {
	if !b.isMaster(fromQQ) {
		return b.notGM()
	}

	b.setSwitch(enable)
	if enable {
		return fmt.Sprintf("已在群%d开启《序列战争》诡秘之主背景小游戏插件。", b.Group)
	}

	return fmt.Sprintf("已在群%d关闭《序列战争》诡秘之主背景小游戏插件。", b.Group)
}

func (b *Bot) IsSilent() bool {
	s := b.getGroupValue("Silence", &SilenceState{}).(*SilenceState)
	if !s.IsSilence {
		return false
	}

	t1Strs := strings.Split(s.OpenStartTime, ":")
	t2Strs := strings.Split(s.OpenEndTime, ":")

	if len(t1Strs) != 2 || len(t2Strs) != 2 {
		return false
	}

	t1Hour, err1 := strconv.Atoi(t1Strs[0])
	t1Minute, err2 := strconv.Atoi(t1Strs[1])
	t2Hour, err3 := strconv.Atoi(t2Strs[0])
	t2Minute, err4 := strconv.Atoi(t2Strs[1])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return false
	}

	t := time.Now()

	dayMinute := t.Hour()*60 + t.Minute()
	startMinute := t1Hour*60 + t1Minute
	endMinute := t2Hour*60 + t2Minute

	fmt.Println("Slient check:", dayMinute, startMinute, endMinute)

	if dayMinute >= startMinute && dayMinute <= endMinute {
		return false
	}

	return true
}

func (b *Bot) SetSilentTime(fromQQ uint64, msg string) string {
	if !b.isMaster(fromQQ) {
		return b.notGM()
	}

	s := b.getGroupValue("Silence", &SilenceState{}).(*SilenceState)

	if strings.Contains(msg, "off") {
		s.IsSilence = false
		b.setGroupValue("Silence", s)
		return "关闭静默功能"
	}

	strs := strings.Split(msg, ";")
	if len(strs) == 4 {
		if strs[1] == "on" {
			s.IsSilence = true
			s.OpenStartTime = strs[2]
			s.OpenEndTime = strs[3]
			b.setGroupValue("Silence", s)
			return "静默功能开启." + fmt.Sprintf("%+v", strs)
		}
	}
	return fmt.Sprintf("%+v", s)
}

func (b *Bot) groupMap(fromQQ uint64, msg string) string {
	if !b.isMaster(fromQQ) {
		return b.notGM()
	}

	gp := b.getGroupValue("GroupMap", &GroupMap{}).(*GroupMap)
	strs := strings.Split(msg, ";")
	if len(strs) != 3 {
		return fmt.Sprintf("当前群数据映射参数：%+v", gp)
	}

	g1, err1 := strconv.ParseUint(strs[1], 10, 64)
	g2, err2 := strconv.ParseUint(strs[2], 10, 64)
	if err1 != nil || err2 != nil {
		return "数字格式异常"
	}

	gp.Mapped = true
	gp.MapFromGroup = g1
	gp.MapToGroup = g2
	b.setGroupValue("GroupMap", gp)
	return "映射成功"
}

func (b *Bot) setDelay(fromQQ uint64, msg string) string {
	if !b.isMaster(fromQQ) {
		return b.notGM()
	}

	gp := GetGlobalValue("ReplyDelay", &ReplyDelay{300}).(*ReplyDelay)
	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return fmt.Sprintf("当前群数据映射参数：%+v", gp)
	}

	g1, err1 := strconv.ParseUint(strs[1], 10, 64)
	if err1 != nil {
		return "数字格式异常"
	}

	gp.DelayMs = g1

	SetGlobalValue("ReplyDelay", gp)

	return "回复延迟配置成功"
}

func (b *Bot) fixNumber(fromQQ uint64) string {
	if !b.isMaster(fromQQ) {
		return b.notGM()
	}

	iter := getDb().NewIterator(util.BytesPrefix(b.getKeyPrefix()), nil)
	for iter.Next() {
		verify := iter.Value()
		var v Person
		rlp.DecodeBytes(verify, &v)
		b.setLuck(v.QQ, 0)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println(err)
	}
	return "数值修复完成，GM附加幸运全部归零了。"
}

func (b *Bot) imgMode(fromQQ uint64, msg string) string {
	if !b.isSuperMaster(fromQQ) {
		return b.notGM()
	}

	img := GetGlobalValue("ImgMode", &ImgMode{Enable: false, Lines: 0}).(*ImgMode)

	strs := strings.Split(msg, ";")
	if len(strs) < 2 {
		return fmt.Sprintf("当前参数：%+v", *img)
	}

	lines, err := strconv.Atoi(strs[1])
	if err != nil {
		return err.Error()
	}

	img.Enable = true
	img.Lines = uint64(lines)

	SetGlobalValue("ImgMode", img)

	return "图片模式修改完成" + fmt.Sprintf("当前参数：%+v", *img)
}

func (b *Bot) foldLineMode(fromQQ uint64, msg string) string {
	if !b.isSuperMaster(fromQQ) {
		return b.notGM()
	}

	fold := GetGlobalValue("FoldLineMode", &FoldLineMode{Enable: true, Lines: 5}).(*FoldLineMode)

	strs := strings.Split(msg, ";")
	if len(strs) < 2 {
		return fmt.Sprintf("当前参数：%+v", *fold)
	}

	lines, err := strconv.Atoi(strs[1])
	if err != nil {
		return err.Error()
	}

	fold.Lines = uint64(lines)

	SetGlobalValue("FoldLineMode", fold)

	return "文字分段修改完成" + fmt.Sprintf("当前参数：%+v", *fold)
}

func GetVersion() *Version {
	return version
}

func GetSuperMaster() uint64 {
	cfgSuper := GetGlobalValue("Supermaster", &Config{}).(*Config)
	return cfgSuper.MasterQQ
}

func GetMoneyMap(group uint64) *MoneyBind {
	b := &Bot{}
	b.Group = group
	bind := b.getMoneyBind()
	return bind
}

func SetMoneyMap(group uint64, bind *MoneyBind) {
	b := &Bot{}
	b.Group = group
	b.setMoneyBind(bind)
}
