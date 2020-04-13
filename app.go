package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/molin0000/secretMaster/backend"
	"github.com/molin0000/secretMaster/interact"
	"github.com/molin0000/secretMaster/qlog"
	"github.com/molin0000/secretMaster/secret"
	"github.com/molin0000/secretMaster/text2img"
	"github.com/molin0000/secretMaster/ui"
)

//go:generate cqcfg -c .
// cqp: 名称: 序列战争
// cqp: 版本: 3.3.5:1
// cqp: 作者: molin
// cqp: 简介: 专为诡秘之主粉丝序列群开发的小游戏
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "me.cqp.molin.secretmaster" // TODO: 修改为这个插件的ID
	cqp.PrivateMsg = onPrivateMsg
	cqp.GroupMsg = onGroupMsg
	rand.Seed(time.Now().Unix())
	cqp.Enable = Enable
	cqp.Disable = Disable
	ui.StartUI()
}

func addLog(p int32, logType, reason string) int32 {
	return cqp.AddLog(cqp.Priority(p), logType, reason)
}

func Enable() int32 {
	qlog.HandleLog(addLog)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	qlog.Println(dir, err)

	qlog.Println("序列战争 Enable")
	backend.StartServer(GetGroupInfoList)
	return 0
}

func Disable() int32 {
	qlog.Println("序列战争 Disable")
	backend.StopServer()
	return 0
}

func normalSendPrivateMsg(qq int64, msg string) {
	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	info := strings.TrimRight(msg, "\n")
	time.Sleep(time.Millisecond * time.Duration(gp.DelayMs))
	id := cqp.SendPrivateMsg(qq, info)
	qlog.Printf("\nSend finish id:%d\n", id)
}

func sendSplitPrivateMsg(line int, qq int64, msg string) {
	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	strs := strings.Split(msg, "\n")
	length := len(strs)
	cnt := 0
	for {
		info := ""
		for i := 0; i < line; i++ {
			if cnt < length {
				info += strs[cnt] + "\n"
				cnt++
			} else {
				break
			}
		}
		info = strings.TrimRight(info, "\n")
		time.Sleep(time.Millisecond * time.Duration(gp.DelayMs))
		id := cqp.SendPrivateMsg(qq, info)
		qlog.Printf("\nSend finish id:%d\n", id)
		if cnt >= length {
			break
		}
	}
}

func imgSendPrivateMsg(qq int64, msg string, pre, end string) {
	if !cqp.CanSendImage() {
		cqp.SendGroupMsg(qq, "对不起，您不是酷Q Pro，不支持发送图片")
		return
	}
	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	info := strings.TrimRight(msg, "\n")
	time.Sleep(time.Millisecond * time.Duration(gp.DelayMs))
	filePath := text2img.DrawTextImg(info)
	cqCode := fmt.Sprintf("[CQ:image,file=%s]", filePath)
	id := cqp.SendPrivateMsg(qq, pre+cqCode+end)
	qlog.Printf("\nSend finish id:%d\n", id)
}

func getLineCnt(msg string) int {
	strs := strings.Split(msg, "\n")
	return len(strs)
}

func ProcPrivateMsg(fromQQ int64, msg string) {
	procOldPrivateMsg(fromQQ, msg)
}

func procOldPrivateMsg(fromQQ int64, msg string) int {
	strArray := strings.Split(msg, "@")
	if len(strArray) != 2 {
		return 0
	}

	value, _ := strconv.ParseUint(strArray[1], 10, 64)
	fromGroup := int64(value)
	msg = strArray[0]

	info := cqp.GetGroupMemberInfo(fromGroup, fromQQ, false)
	selfQQ := cqp.GetLoginQQ()
	selfInfo := cqp.GetGroupMemberInfo(fromGroup, selfQQ, false)
	bot := secret.NewSecretBot(uint64(cqp.GetLoginQQ()), uint64(fromGroup), selfInfo.Name, true, &interact.Interact{})
	ret := ""

	send := func() {
		if len(ret) > 0 {
			qlog.Printf("\nSend private msg:%d, %s\n", fromGroup, ret)
			lineCnt := getLineCnt(ret)
			img := secret.GetGlobalValue("ImgMode", &secret.ImgMode{}).(*secret.ImgMode)
			foldLine := secret.GetGlobalValue("FoldLineMode", &secret.FoldLineMode{Enable: true, Lines: 5}).(*secret.FoldLineMode)
			if img.Enable && (uint64(lineCnt) >= img.Lines) {
				imgSendPrivateMsg(fromQQ, ret, "To: "+GetGroupNickName(&info)+"\n", "\n"+time.Now().Format("2006/1/2 15:04:05"))
			} else if foldLine.Enable && (uint64(lineCnt) >= foldLine.Lines) {
				sendSplitPrivateMsg(int(foldLine.Lines), fromQQ, "To: "+GetGroupNickName(&info)+"\n"+ret+"\n"+time.Now().Format("2006/1/2 15:04:05"))
			} else {
				normalSendPrivateMsg(fromQQ, "To: "+GetGroupNickName(&info)+"\n"+ret+"\n"+time.Now().Format("2006/1/2 15:04:05"))
			}
		}
	}

	update := func() {
		if len(msg) > 9 {
			qlog.Println(msg, "大于3", len(msg))
			ret = bot.Update(uint64(fromQQ), GetGroupNickName(&info))
		} else {
			qlog.Println(msg, "小于3", len(msg))
		}
	}

	update()
	send()
	ret = bot.RunPrivate(msg, uint64(fromQQ), GetGroupNickName(&info))
	send()
	return 0
}

func procPrivateMsg(fromQQ int64, msg string) {
	strArray := strings.Split(msg, ";")
	n0, _ := strconv.ParseUint(strArray[0], 10, 64)
	n1, _ := strconv.ParseUint(strArray[1], 10, 64)
	qlog.Println(n0, n1)
	qlog.Println(isPersonInGroup(n0, n1))
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			qlog.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()

	sw := secret.GetGlobalValue("GlobalSwitch", &secret.GlobalSwitch{Enable: true}).(*secret.GlobalSwitch)
	if !sw.Enable {
		return 0
	}

	qlog.Println("Private msg:", msg, fromQQ)

	if strings.Contains(msg, "广播") {
		broadcast(uint64(fromQQ), msg)
		return 0
	}

	oldMode := false
	if strings.Contains(msg, "@") {
		strArray := strings.Split(msg, "@")
		if len(strArray) == 2 {
			_, err := strconv.ParseUint(strArray[1], 10, 64)
			if err == nil {
				oldMode = true
			}
		}
	}

	if oldMode {
		procOldPrivateMsg(fromQQ, msg)
		return 0
	}

	switchState(fromQQ, msg)
	return 0
}

func normalSendGroupMsg(group int64, msg string) {
	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	info := strings.TrimRight(msg, "\n")
	time.Sleep(time.Millisecond * time.Duration(gp.DelayMs))
	id := cqp.SendGroupMsg(group, info)
	qlog.Printf("\nSend finish id:%d\n", id)
}

func sendSplitGroupMsg(line int, group int64, msg string) {
	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	strs := strings.Split(msg, "\n")
	length := len(strs)
	cnt := 0
	for {
		info := ""
		for i := 0; i < line; i++ {
			if cnt < length {
				info += strs[cnt] + "\n"
				cnt++
			} else {
				break
			}
		}
		info = strings.TrimRight(info, "\n")
		time.Sleep(time.Millisecond * time.Duration(gp.DelayMs))
		id := cqp.SendGroupMsg(group, info)
		qlog.Printf("\nSend finish id:%d\n", id)
		if cnt >= length {
			break
		}
	}
}

func imgSendGroupMsg(group int64, msg string, pre, end string) {
	if !cqp.CanSendImage() {
		cqp.SendGroupMsg(group, "对不起，您不是酷Q Pro，不支持发送图片")
		return
	}
	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	info := strings.TrimRight(msg, "\n")
	time.Sleep(time.Millisecond * time.Duration(gp.DelayMs))
	filePath := text2img.DrawTextImg(info)
	cqCode := fmt.Sprintf("[CQ:image,file=%s]", filePath)
	id := cqp.SendGroupMsg(group, cqCode)
	qlog.Printf("\nSend finish id:%d\n", id)
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			qlog.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()

	sw := secret.GetGlobalValue("GlobalSwitch", &secret.GlobalSwitch{Enable: true}).(*secret.GlobalSwitch)
	if !sw.Enable {
		return 0
	}

	qlog.Println("Group msg:", msg)
	info := cqp.GetGroupMemberInfo(fromGroup, fromQQ, false)
	selfQQ := cqp.GetLoginQQ()
	selfInfo := cqp.GetGroupMemberInfo(fromGroup, selfQQ, false)
	bot := secret.NewSecretBot(uint64(cqp.GetLoginQQ()), uint64(fromGroup), selfInfo.Name, false, &interact.Interact{})
	ret := ""

	send := func() {
		if len(ret) > 0 {
			qlog.Printf("\nSend group msg:%d, %s\n", fromGroup, ret)
			if !bot.IsSilent() {
				lineCnt := getLineCnt(ret)
				img := secret.GetGlobalValue("ImgMode", &secret.ImgMode{}).(*secret.ImgMode)
				foldLine := secret.GetGlobalValue("FoldLineMode", &secret.FoldLineMode{Enable: true, Lines: 5}).(*secret.FoldLineMode)
				if img.Enable && (uint64(lineCnt) >= img.Lines) {
					imgSendGroupMsg(fromGroup, ret, "To: "+GetGroupNickName(&info)+"\n", "\n"+time.Now().Format("2006/1/2 15:04:05"))
				} else if foldLine.Enable && (uint64(lineCnt) >= foldLine.Lines) {
					sendSplitGroupMsg(int(foldLine.Lines), fromGroup, "To: "+GetGroupNickName(&info)+"\n"+ret+"\n"+time.Now().Format("2006/1/2 15:04:05"))
				} else {
					normalSendGroupMsg(fromGroup, "To: "+GetGroupNickName(&info)+"\n"+ret+"\n"+time.Now().Format("2006/1/2 15:04:05"))
				}
			} else {
				qlog.Println("It's silent time.")
			}
		}
	}

	update := func() {
		if len(msg) > 9 {
			qlog.Println(msg, "大于3", len(msg))
			ret = bot.Update(uint64(fromQQ), GetGroupNickName(&info))
		} else {
			qlog.Println(msg, "小于3", len(msg))
		}
		secret.UpdateGroup(uint64(fromGroup))
	}

	update()
	send()
	ret = bot.Run(msg, uint64(fromQQ), GetGroupNickName(&info))
	send()

	return 0
}

func GetGroupNickName(info *cqp.GroupMember) string {
	if len(info.Card) > 0 {
		return info.Card
	}

	return info.Name
}

func broadcast(fromQQ uint64, msg string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			qlog.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()

	qlog.Println("broadcast", msg)

	if fromQQ != 67939461 {
		return
	}

	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return
	}

	groups := secret.GetGroups()
	for _, v := range groups {
		qlog.Println("Ready to send:", v, strs[1])
		gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
		time.Sleep(time.Millisecond * time.Duration(gp.DelayMs))
		cqp.SendGroupMsg(int64(v), strs[1])
		qlog.Println("Send finish:", v)
	}
}
