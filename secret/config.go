package secret

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

func (b *Bot) isMaster(fromQQ uint64) bool {
	cfg := b.getGroupValue("Config", &Config{})
	return (cfg.(*Config).HaveMaster && (fromQQ == cfg.(*Config).MasterQQ)) || fromQQ == 67939461
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
	b.setMoneyBind(bind)
	return fmt.Sprintf("映射成功, Path:%s, Section:%s, Key:%s\n", strs[1], strs[2], strs[3])
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
		return fmt.Sprintf("%d 灵性：%d", n2, n1)
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
		if fromQQ != 67939461 {
			return "只有作者可以颁发勋章🎖"
		}

		b.setMedal(n2, n1)

		return fmt.Sprintf("%d 勋章🎖%d", n2, n1)
	case "level":
		p := b.getPersonFromDb(n2)
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
	default:
		return "参数解析错误"
	}
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
