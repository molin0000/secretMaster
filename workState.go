package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/molin0000/secretMaster/backend"
	"github.com/molin0000/secretMaster/secret"
)

type PersonState struct {
	Group  uint64
	Switch bool
	State  uint64 // 0:Start, 1:GroupSelect, 2:Play
}

func getGroupInfo(group int64, noCach bool) (detail *cqp.GroupDetail) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("获取群信息失败", group)
			detail = nil
		}
	}()
	getInfo := cqp.GetGroupInfo(group, noCach)
	return &getInfo
}

func getGroupList() (ret string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("getGroupList panic")
		}
	}()
	groups := secret.GetGroups()
	fmt.Println("getGroupList", groups)
	ret = "\n"
	for i, v := range groups {
		detail := getGroupInfo(int64(v), false)
		if detail != nil {
			fmt.Printf("%+v", detail)
			ret += fmt.Sprintf("%d) %s [%d] (%d/%d);\n", i, detail.Name, v, detail.MembersNum, detail.MaxMemberNum)
			fmt.Println(ret)
		} else {
			secret.RemoveGroup(uint64(i))
		}
	}
	return ret
}

func GetGroupInfoList() []*backend.GroupInfo {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("getGroupList panic")
		}
	}()
	groups := secret.GetGroups()
	fmt.Println("getGroupList", groups)
	gps := make([]*backend.GroupInfo, 0)
	key := uint64(0)
	for i, v := range groups {
		detail := getGroupInfo(int64(v), false)
		if detail != nil {
			gi := &backend.GroupInfo{}
			gi.Key = key
			key++
			gi.Group = v
			gi.Member = fmt.Sprintf("%d/%d", detail.MembersNum, detail.MaxMemberNum)
			b := &secret.Bot{}
			b.Group = v
			gi.Master = b.GetMaster()
			gi.Silence = b.IsSilent()
			gi.Switch = b.GetSwitch()
			gps = append(gps, gi)
		} else {
			secret.RemoveGroup(uint64(i))
		}
	}
	return gps
}

func switchState(fromQQ int64, msg string) {
	state := secret.GetGlobalPersonValue("State", uint64(fromQQ), &PersonState{0, false, 0}).(*PersonState)
	if !state.Switch && !strings.Contains(msg, "序列战争") && msg != "关闭私聊" {
		return
	}

	fmt.Printf("switchState:%+v", state)

	if msg == "关闭私聊" {
		state.Switch = false
		state.State = 0
		secret.SetGlobalPersonValue("State", uint64(fromQQ), state)
		cqp.SendPrivateMsg(fromQQ, "序列战争私聊功能已关闭，下次开启请说【序列战争】。")
		return
	}

	if msg == "@" {
		state.State = 1
		secret.SetGlobalPersonValue("State", uint64(fromQQ), state)
		cqp.SendPrivateMsg(fromQQ,
			`请在下面列表中选择你的默认QQ群（回复序号选择）：`+getGroupList())
		return
	}

	switch state.State {
	case 0:
		cqp.SendPrivateMsg(fromQQ,
			`⭐️⭐️⭐️⭐️⭐️⭐️⭐️⭐️⭐️
欢迎来到序列战争游戏世界
这是基于《诡秘之主》小说世界观的同人QQ群小游戏。它将可以根据你在群聊中的消息数量进行经验积累和升级。
你可以在22条成神途径中选取你喜爱的非凡途径进行扮演，并在积累足够的经验后进行序列晋升。
游戏现包含6大功能：资料、途径、教会、探险、商店、生活。
在游玩过程中，可随时回复 [帮助] 查看可用指令列表，回复 [@] 可切换默认群聊，回复 [关闭私聊] 可关闭私聊游戏功能，下次想开启时再输入 [序列战争] 即可。`)

		groupInfo := getGroupList()
		cqp.SendPrivateMsg(fromQQ,
			"首先请在下面列表中选择你的默认QQ群（回复序号选择）：")
		fmt.Println("groupInfo", groupInfo)
		cqp.SendPrivateMsg(fromQQ, groupInfo)
		state.State = 1
		state.Switch = true
		secret.SetGlobalPersonValue("State", uint64(fromQQ), state)
		return
	case 1:
		selection, err := strconv.Atoi(msg)
		if err != nil {
			cqp.SendPrivateMsg(fromQQ,
				`数值错误，请在下面列表中选择你的默认QQ群（回复序号选择）：`+getGroupList())
			return
		}
		groupList := secret.GetGroups()
		if selection >= len(groupList) {
			cqp.SendPrivateMsg(fromQQ,
				`数值错误，请在下面列表中选择你的默认QQ群（回复序号选择）：`+getGroupList())
			return
		}
		if !isPersonInGroup(groupList[selection], uint64(fromQQ)) {
			cqp.SendPrivateMsg(fromQQ,
				`你不在这个QQ群中，请在下面列表中选择你的默认QQ群（回复序号选择）：`+getGroupList())
			return
		}
		state.Group = groupList[selection]
		state.State = 2
		secret.SetGlobalPersonValue("State", uint64(fromQQ), state)
		cqp.SendPrivateMsg(fromQQ, "⭐️默认群聊设置成功！请回复[帮助]看看都有哪些功能吧~")
		return
	case 2:
		procOldPrivateMsg(fromQQ, fmt.Sprintf("%s@%d", msg, state.Group))
	}

	return
}

func isPersonInGroup(group, qq uint64) (exist bool) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			exist = false
		}
	}()
	cqp.GetGroupMemberInfo(int64(group), int64(qq), false)
	return true
}
