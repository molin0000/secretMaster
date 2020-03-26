package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/molin0000/secretMaster/interact"
	"github.com/molin0000/secretMaster/secret"
)

//go:generate cqcfg -c .
// cqp: 名称: 序列战争
// cqp: 版本: 3.1.5:1
// cqp: 作者: molin
// cqp: 简介: 专为诡秘之主粉丝序列群开发的小游戏
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "me.cqp.molin.secretmaster" // TODO: 修改为这个插件的ID
	cqp.PrivateMsg = onPrivateMsg
	cqp.GroupMsg = onGroupMsg
	rand.Seed(time.Now().Unix())
}

func sendSplitPrivateMsg(qq int64, msg string) {
	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	strs := strings.Split(msg, "\n")
	length := len(strs)
	cnt := 0
	for {
		info := ""
		for i := 0; i < 5; i++ {
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
		fmt.Printf("\nSend finish id:%d\n", id)
		if cnt >= length {
			break
		}
	}
}

func procOldPrivateMsg(fromQQ int64, msg string) int {
	strArray := strings.Split(msg, "@")
	if len(strArray) != 2 {
		return 0
	}

	value, _ := strconv.ParseUint(strArray[1], 10, 64)
	fromGroup := int64(value)
	msg = strArray[0]

	info := cqp.GetGroupMemberInfo(fromGroup, fromQQ, true)
	selfQQ := cqp.GetLoginQQ()
	selfInfo := cqp.GetGroupMemberInfo(fromGroup, selfQQ, false)
	bot := secret.NewSecretBot(uint64(cqp.GetLoginQQ()), uint64(fromGroup), selfInfo.Name, true, &interact.Interact{})
	ret := ""

	send := func() {
		if len(ret) > 0 {
			fmt.Printf("\nSend private msg:%d, %s\n", fromGroup, ret)
			sendSplitPrivateMsg(fromQQ, ret)
		}
	}

	update := func() {
		if len(msg) > 9 {
			fmt.Println(msg, "大于3", len(msg))
			ret = bot.Update(uint64(fromQQ), GetGroupNickName(&info))
		} else {
			fmt.Println(msg, "小于3", len(msg))
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
	fmt.Println(n0, n1)
	fmt.Println(isPersonInGroup(n0, n1))
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()

	fmt.Println("Private msg:", msg, fromQQ)

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

func sendSplitGroupMsg(group int64, msg string) {
	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	strs := strings.Split(msg, "\n")
	length := len(strs)
	cnt := 0
	for {
		info := ""
		for i := 0; i < 5; i++ {
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
		fmt.Printf("\nSend finish id:%d\n", id)
		if cnt >= length {
			break
		}
	}
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()

	fmt.Println("Group msg:", msg)
	info := cqp.GetGroupMemberInfo(fromGroup, fromQQ, true)
	selfQQ := cqp.GetLoginQQ()
	selfInfo := cqp.GetGroupMemberInfo(fromGroup, selfQQ, false)
	bot := secret.NewSecretBot(uint64(cqp.GetLoginQQ()), uint64(fromGroup), selfInfo.Name, false, &interact.Interact{})
	ret := ""

	send := func() {
		if len(ret) > 0 {
			fmt.Printf("\nSend group msg:%d, %s\n", fromGroup, ret)
			if !bot.IsSilent() {
				sendSplitGroupMsg(fromGroup, "@"+GetGroupNickName(&info)+" "+ret)
			} else {
				fmt.Println("It's silent time.")
			}
		}
	}

	update := func() {
		if len(msg) > 9 {
			fmt.Println(msg, "大于3", len(msg))
			ret = bot.Update(uint64(fromQQ), GetGroupNickName(&info))
		} else {
			fmt.Println(msg, "小于3", len(msg))
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
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()

	fmt.Println("broadcast", msg)

	if fromQQ != 67939461 {
		return
	}

	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return
	}

	groups := secret.GetGroups()
	for _, v := range groups {
		fmt.Println("Ready to send:", v, strs[1])
		gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
		time.Sleep(time.Millisecond * time.Duration(gp.DelayMs))
		cqp.SendGroupMsg(int64(v), strs[1])
		fmt.Println("Send finish:", v)
	}
}
