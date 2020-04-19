package calculator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/molin0000/secretMaster/qlog"
)

type CalcGame struct {
	Level     uint64
	Result    uint64
	StartTime uint64
	Money     uint64
	Success   bool
	Speed     uint64
}

func NewCalcGame() *CalcGame {
	return &CalcGame{}
}

func (c *CalcGame) Start() string {

	info := "\n你去替朋友参加知识教会的入门考试，他说考过了奖金全归你，考不过奖金归他。"
	info += "\n要求每一题在20秒钟内作答，但是监考老师慢悠悠的发卷就耗去了不少时间，留给你的时间已经不多了。"
	info += c.ShowQuestion(c.Level)

	return info
}

func (c *CalcGame) ShowQuestion(level uint64) string {
	a := rand.Intn(899) + 100
	b := rand.Intn(899) + 100
	c.Result = uint64(a + b)
	c.StartTime = uint64(time.Now().Unix())
	return fmt.Sprintf("\n第%d题：%d + %d = ?", level+1, a, b)
}

func (c *CalcGame) GiveResult(result uint64) (msg string, finish bool) {
	ret := c.Result == result
	timeNow := time.Now().Unix()
	info := ""
	c.Speed = uint64(timeNow - int64(c.StartTime))
	if timeNow-int64(c.StartTime) > 20 {
		qlog.Println(timeNow, c.StartTime)
		return info + "很遗憾，你超时了。没有通过考试。你朋友对你的表现十分失望。", true
	}
	if ret {
		c.Level++
		info += "\n恭喜你，答对了！" + fmt.Sprintf("监考老师递给你%d金镑奖金！", c.Level*25)
		if c.Level >= 3 {
			c.Money += 150
			return info + "\n你全部答对了，获得了共计150金镑的奖金，你的朋友以为荣！", true
		}
		info += c.ShowQuestion(c.Level)
		return info, false
	} else {
		return "很遗憾，你答错了，没有通过考试。你朋友对你的表现十分失望", true
	}
}
