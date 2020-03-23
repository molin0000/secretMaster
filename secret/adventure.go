package secret

import (
	"fmt"
	"math/rand"
	"time"
)

func (b *Bot) adventure(fromQQ uint64, limit bool) string {
	a := b.getAdvFromDb(fromQQ)
	if limit && a.DayCnt >= (3+b.getAdditionAdventure(fromQQ)) {
		return "对不起，您今日奇遇探险机会已经用完"
	}

	money := b.getMoney(fromQQ)
	if !limit {
		if money > 50 {
			b.setMoney(fromQQ, -50)
		} else {
			return "钱包空空，买不起了哦"
		}
	}

	a.DayCnt++
	a.Days = uint64(time.Now().Unix() / (3600 * 24))
	b.setAdvToDb(fromQQ, a)

	rand.Seed(time.Now().UnixNano())

	luck := b.getLuck(fromQQ)
	if luck >= 19 {
		luck = 5
	}

	i := int(luck*5) + rand.Intn(100-int(luck*5))
	info := ""
	m := 0
	e := 0
	fmt.Println("i:", i)

	for p := 0; p < len(adventureEvents); p++ {
		i = i - adventureEvents[p].Probability
		fmt.Println("i:", i)
		if i < 0 {
			switch adventureEvents[p].Type {
			case 1:
				m = -1 * (10 + rand.Intn(41))
				e = -1 * (10 + rand.Intn(41))
				info = fmt.Sprintf("%s经验:%d, 金镑:%d", adventureEvents[p].Messages[rand.Intn(len(adventureEvents[p].Messages))], e, m)
			case 2:
				m = 0
				e = -1 * (10 + rand.Intn(41))
				info = fmt.Sprintf("%s经验:%d", adventureEvents[p].Messages[rand.Intn(len(adventureEvents[p].Messages))], e)
			case 3:
				m = -1 * (10 + rand.Intn(41))
				e = 0
				info = fmt.Sprintf("%s金镑:%d", adventureEvents[p].Messages[rand.Intn(len(adventureEvents[p].Messages))], m)
			case 4:
				m = (20 + rand.Intn(81))
				e = (20 + rand.Intn(81))
				info = fmt.Sprintf("%s经验:%d, 金镑:%d", adventureEvents[p].Messages[rand.Intn(len(adventureEvents[p].Messages))], e, m)
			case 5:
				m = (20 + rand.Intn(81))
				e = 0
				info = fmt.Sprintf("%s金镑:%d", adventureEvents[p].Messages[rand.Intn(len(adventureEvents[p].Messages))], m)
			case 6:
				m = 0
				e = (20 + rand.Intn(81))
				info = fmt.Sprintf("%s经验:%d", adventureEvents[p].Messages[rand.Intn(len(adventureEvents[p].Messages))], e)
			default:
				m = 0
				e = 0
				info = fmt.Sprintf("%s", adventureEvents[p].Messages[rand.Intn(len(adventureEvents[p].Messages))])
			}
			break
		}
	}

	b.setMoney(fromQQ, m)
	b.setExp(fromQQ, e)
	return info
}
