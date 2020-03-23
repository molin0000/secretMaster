package secret

import (
	"fmt"
	"strings"
	"time"

	"github.com/molin0000/secretMaster/competition"
)

func (b *Bot) startCompetition(fromQQ uint64) string {
	magic := b.getMagic(fromQQ)
	if magic < 5 {
		return "灵性不足"
	}

	b.setMagic(fromQQ, -5)

	qs := b.getPersonValue("Competition", fromQQ, &CompetitionState{true, nil, uint64(time.Now().Unix()), 0, 0}).(*CompetitionState)

	q := competition.GetRandomQuestion()
	if q == nil {
		return "题库为空，请先导入competition.xlsx文件"
	}

	qs.Q = q
	qs.IsPlaying = true
	qs.StartTime = uint64(time.Now().Unix())
	b.setPersonValue("Competition", fromQQ, qs)

	fmt.Println("学识竞猜开始", fromQQ, *qs.Q)

	info := fmt.Sprintf(`
============
题目：%s
选项/提示：%s
答题时限：%d秒
作者：%s
============
请回答：
`, qs.Q.Question, qs.Q.Option, qs.Q.Time, qs.Q.Author)
	return info
}

func (b *Bot) checkCompetion(fromQQ uint64, msg string) string {
	qs := b.getPersonValue("Competition", fromQQ, &CompetitionState{false, nil, 0, 0, 0}).(*CompetitionState)
	if !qs.IsPlaying {
		return ""
	}

	fmt.Println("msg", msg, "qs", qs)

	if qs.Q == nil {
		qs.VictoryCnt = 0
		qs.IsPlaying = false
		b.setPersonValue("Competition", fromQQ, qs)
		return "数据异常，答题终止"
	}

	now := uint64(time.Now().Unix())
	if now-qs.StartTime > qs.Q.Time {
		qs.VictoryCnt = 0
		qs.IsPlaying = false
		b.setPersonValue("Competition", fromQQ, qs)
		return "对不起，答题超时！"
	}

	var answer string
	strs := strings.Split(msg, "] ")
	fmt.Println(strs)
	if len(strs) > 1 {
		answer = strs[1]
	} else {
		answer = msg
	}

	if strings.ToLower(qs.Q.Answer) != strings.ToLower(answer) {
		qs.VictoryCnt = 0
		qs.IsPlaying = false
		b.setPersonValue("Competition", fromQQ, qs)
		return "对不起，回答错误！"
	}

	b.setMoney(fromQQ, 10)
	b.setExp(fromQQ, 10)

	qs.VictoryCnt++
	if qs.VictoryCnt > qs.MaxVictoryCnt {
		qs.MaxVictoryCnt = qs.VictoryCnt
	}
	qs.IsPlaying = false

	b.setPersonValue("Competition", fromQQ, qs)

	max := b.getGroupValue("MaxVictory", &MaxVictory{0, "无"}).(*MaxVictory)
	if qs.MaxVictoryCnt > max.VictoryCnt {
		max.VictoryCnt = qs.MaxVictoryCnt
		max.Name = b.CurrentNick
		b.setGroupValue("MaxVictory", max)
	}

	return fmt.Sprintf("\n恭喜你！回答正确！经验+10，金镑+10\n当前连胜:%d, 最佳连胜:%d, 本群最佳:%d（%s)", qs.VictoryCnt, qs.MaxVictoryCnt, max.VictoryCnt, max.Name)
}
