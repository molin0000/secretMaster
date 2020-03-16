package secret

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func calcInterest(bank *Bank) *Bank {
	today := uint64(time.Now().Unix() / (3600 * 24))
	if bank.Date != 0 && bank.Date != today {
		bank.Amount = uint64(float64(bank.Amount) * math.Pow(1.03, float64(today-bank.Date)))
		bank.Date = today
	}
	fmt.Printf("%+v", bank)
	return bank
}

func (b *Bot) bank(fromQQ uint64, msg string) string {
	strs := strings.Split(msg, ";")
	if len(strs) < 2 {
		return "指令错误" + fmt.Sprintf("%+v", strs)
	}

	switch strs[1] {
	case "存款":
		if len(strs) < 3 {
			return "指令错误" + fmt.Sprintf("%+v", strs)
		}
		money, err := strconv.Atoi(strs[2])
		if err != nil {
			return "指令错误" + fmt.Sprintf("%+v", err)
		}
		if money < 100 {
			return "每次存款最少100金镑。"
		}
		if uint64(money) > b.getMoney(fromQQ) {
			return "你并没有那么多钱。"
		}

		bank := b.getPersonValue("Bank", fromQQ, &Bank{}).(*Bank)
		bank = calcInterest(bank)
		bank.Amount += uint64(money)
		b.setPersonValue("Bank", fromQQ, bank)
		b.setMoney(fromQQ, -1*money)
		return fmt.Sprintf("成功存入%d金镑，当前存款%d金镑。", money, bank.Amount)
	case "取款":
		if len(strs) < 3 {
			return "指令错误" + fmt.Sprintf("%+v", strs)
		}
		money, err := strconv.Atoi(strs[2])
		if err != nil {
			return "指令错误" + fmt.Sprintf("%+v", err)
		}
		if money <= 0 {
			return "你想取个西瓜？"
		}
		bank := b.getPersonValue("Bank", fromQQ, &Bank{}).(*Bank)
		bank = calcInterest(bank)
		if uint64(money) > bank.Amount {
			return "你没有那么多钱。"
		}

		bank.Amount -= uint64(money)
		b.setPersonValue("Bank", fromQQ, bank)
		b.setMoney(fromQQ, money)
		return fmt.Sprintf("成功取出%d金镑", money)
	case "查账":
		bank := b.getPersonValue("Bank", fromQQ, &Bank{}).(*Bank)
		bank = calcInterest(bank)
		b.setPersonValue("Bank", fromQQ, bank)
		return fmt.Sprintf("当前存款%d金镑。日利率3％", bank.Amount)
	default:
		return "指令错误" + fmt.Sprintf("%+v", strs)
	}
}

func (b *Bot) work(fromQQ uint64, msg string) string {
	strs := strings.Split(msg, ";")
	w := b.getPersonValue("Work", fromQQ, &Work{}).(*Work)
	if len(strs) < 2 {
		info := "\n当前支持的工作种类有："
		for _, v := range workList {
			info += fmt.Sprintf("\n" + v.Name)
		}
		return info
	}

	if len(w.Name) > 0 && strs[1] != "停止" {
		return "你必须停止当前工作才能进行下一项工作, 当前正在：" + w.Name
	}

	if strs[1] == "停止" {
		if len(w.Name) == 0 {
			return "你没有工作。"
		}
		if w.Name == "值夜者小队" || w.Name == "赏金猎人" {
			b.removePersonValue("Work", fromQQ)
			return b.exitBattleField(fromQQ)
		}
		today := uint64(time.Now().Unix() / (3600 * 24))
		time := today - w.Date
		m := uint64(0)
		e := uint64(0)
		if time > 0 {
			for _, v := range workList {
				if v.Name == w.Name {
					m = time * v.MoneyAdd
					e = time * v.ExpAdd
					break
				}
			}
			b.setMoney(fromQQ, int(m))
			b.setExp(fromQQ, int(e))
		}
		b.removePersonValue("Work", fromQQ)
		return fmt.Sprintf("\n你工作了%d天，内容为:%s，收获%d金镑，%d经验", time, w.Name, m, e)
	}

	for i, v := range workList {
		if strs[1] == v.Name {
			today := uint64(time.Now().Unix() / (3600 * 24))
			v.Date = today
			if i == 0 || i == 1 {
				str, ret := b.joinBattleField(fromQQ, int64(i))
				if ret {
					b.setPersonValue("Work", fromQQ, v)
					return "你开始了工作：" + v.Name + str
				} else {
					return str
				}
			} else {
				b.setPersonValue("Work", fromQQ, v)
			}
			return "你开始了工作：" + v.Name
		}
	}

	return "未知工作内容。"
}

func (b *Bot) fishing(fromQQ uint64) string {
	if b.getMagic(fromQQ) < 5 {
		return "你已经灵性枯竭了，再去钓鱼，怕会被鱼吃了。"
	}

	b.setMagic(fromQQ, -5)

	totalProp := 0
	for _, v := range fishList {
		totalProp += int(v.Property)
	}

	luckNum := rand.Intn(totalProp) + 1

	for _, v := range fishList {
		luckNum -= int(v.Property)
		if luckNum <= 0 {
			if v.Money == 1000 {
				p := b.getPersonValue("Person", fromQQ, &Person{}).(*Person)
				if p.SecretLevel < 7 || p.SecretLevel > 10 {
					b.setExp(fromQQ, -100)
					return "你钓鱼钓到了可怕的" + v.Name + ", 你历经九死一生，终于逃脱了，经验-100."
				} else {
					b.setMoney(fromQQ, 1000)
					return "你钓鱼钓到了可怕的" + v.Name + ", 它历经九死一生，终究不是你的对手，你卖掉它，得到了" + fmt.Sprintf("%d", v.Money) + "金镑。"
				}
			}
			b.setMoney(fromQQ, int(v.Money))
			return "你钓鱼钓到了" + v.Name + ", 你卖掉它，得到了" + fmt.Sprintf("%d", v.Money) + "金镑。"
		}
	}

	return "你一无所获。"
}

func (b *Bot) lottery(fromQQ uint64) string {
	if b.getMoney(fromQQ) < 2 {
		return "你身无分文，许愿不投币怎么行。"
	}
	b.setMoney(fromQQ, -2)

	if rand.Intn(500) == 0 {
		b.setMoney(fromQQ, 1000)
		return "天啊！！女神保佑！你的屋顶上有东西落下来了！1000金镑！！"
	}

	return "很遗憾，你的许愿没有回应。"
}

func (b *Bot) redPack(fromQQ uint64, msg string) string {
	p := b.getPersonFromDb(fromQQ)
	if p.SecretLevel < 2 || p.SecretLevel > 9 {
		return "序列7以下不可发红包"
	}

	strs := strings.Split(msg, ";")
	if len(strs) != 3 {
		return "指令格式错误，应为：红包;金额;QQ"
	}

	n1, err1 := strconv.Atoi(strs[1])
	if n1 <= 0 || err1 != nil {
		return "金额数值异常"
	}

	if n1 > int(100*math.Pow(2, float64(p.SecretLevel))) {
		return "单次金额超出限制：" + fmt.Sprintf("%d", uint64(100*math.Pow(2, float64(p.SecretLevel))))
	}

	n2, err2 := strconv.ParseUint(strs[2], 10, 64)
	if err2 != nil {
		return "QQ数值异常"
	}

	m := b.getMoney(fromQQ)
	if m < uint64(n1) {
		return "对不起，你的钱不够呢"
	}

	b.setMoney(fromQQ, -1*n1)

	b.setMoney(n2, n1)
	return fmt.Sprintf("你成功向%d发送了%d金镑的红包", n2, n1)
}
