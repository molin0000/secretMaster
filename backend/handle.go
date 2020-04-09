package backend

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/labstack/echo"
	"github.com/molin0000/secretMaster/competition"
	"github.com/molin0000/secretMaster/interact"
	"github.com/molin0000/secretMaster/mission"
	"github.com/molin0000/secretMaster/pet"
	"github.com/molin0000/secretMaster/secret"
)

func response(e echo.Context, data interface{}, err error) error {
	ret := map[string]interface{}{}

	if err == nil {
		ret["status"] = 0
		ret["desc"] = "success"
		ret["data"] = data
	} else {
		ret["status"] = -1
		ret["desc"] = err.Error()
	}

	return e.JSONPretty(http.StatusOK, ret, "  ")
}

func GetVersion(e echo.Context) (err error) {
	v := secret.GetVersion()
	return response(e, v, err)
}

type Count struct {
	MissionCnt   int
	QuestionCnt  int
	RealPetCnt   int
	SpiritPetCnt int
	PetFoodCnt   int
}

func GetCount(e echo.Context) (err error) {
	v := &Count{}
	v.MissionCnt = mission.GetMissionCount()
	v.QuestionCnt = competition.GetQuestionCount()
	v.RealPetCnt = pet.GetRealPetCount()
	v.SpiritPetCnt = pet.GetSpiritPetCount()
	v.PetFoodCnt = pet.GetFoodCount()

	return response(e, v, err)
}

func GetInterfaceList(e echo.Context) (err error) {
	strs := []string{
		"/version",
		"/count",
		"/group",
		"/password",
		"/supermaster",
		"/delay",
		"/chat",
		"/moneyMap",
		"/imageMode",
		"/textSegment",
		"/activities",
		"/locked",
	}

	return response(e, strs, err)
}

func GetSuperMaster(e echo.Context) (err error) {
	superMaster := secret.GetSuperMaster()
	return response(e, superMaster, err)
}

func GetDelay(e echo.Context) (err error) {
	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	return response(e, gp, err)
}

func GetImageMode(e echo.Context) (err error) {
	img := secret.GetGlobalValue("ImgMode", &secret.ImgMode{}).(*secret.ImgMode)
	return response(e, img, err)
}

func GetTextSegment(e echo.Context) (err error) {
	foldLine := secret.GetGlobalValue("FoldLineMode", &secret.FoldLineMode{Enable: true, Lines: 5}).(*secret.FoldLineMode)
	return response(e, foldLine, err)
}

func GetMoneyMap(e echo.Context) (err error) {
	var req struct {
		Group uint64 `json:"group"    query:"group"`
	}

	err = e.Bind(&req)
	fmt.Printf("%+v", req)

	mp := secret.GetMoneyMap(req.Group)
	return response(e, mp, err)
}

type GroupInfo struct {
	Key     uint64 `json:"key" xml:"key" form:"key" query:"key"`
	Group   uint64 `json:"group" xml:"group" form:"group" query:"group"`
	Member  string `json:"member" xml:"member" form:"member" query:"member"`
	Master  uint64 `json:"master" xml:"master" form:"master" query:"master"`
	Switch  bool   `json:"switch" xml:"switch" form:"switch" query:"switch"`
	Silence bool   `json:"silence" xml:"silence" form:"silence" query:"silence"`
}

var GetGroupInfoList func() []*GroupInfo

type GetGroupRet struct {
	GlobalSwitch  bool         `json:"globalSwitch" xml:"globalSwitch" form:"globalSwitch" query:"globalSwitch"`
	GlobalSilence bool         `json:"globalSilence" xml:"globalSilence" form:"globalSilence" query:"globalSilence"`
	Groups        []*GroupInfo `json:"groups" xml:"groups" form:"groups" query:"groups"`
}

func GetGroup(e echo.Context) (err error) {
	ret := &GetGroupRet{}
	ret.Groups = GetGroupInfoList()
	sw := secret.GetGlobalValue("GlobalSwitch", &secret.GlobalSwitch{Enable: true}).(*secret.GlobalSwitch)
	si := secret.GetGlobalValue("GlobalSilence", &secret.GlobalSilence{}).(*secret.GlobalSilence)
	ret.GlobalSwitch = sw.Enable
	ret.GlobalSilence = si.Enable
	return response(e, ret, err)
}

func GetActivities(e echo.Context) (err error) {
	acts := secret.GetActivities()
	return response(e, acts, err)
}

func GetLocked(e echo.Context) (err error) {
	value := secret.GetGlobalPersonValue("Password", 0, &secret.Password{QQ: 0, Password: ""}).(*secret.Password)
	return response(e, len(value.Password) > 0, err)
}

func PostPassword(e echo.Context) (err error) {
	p := &secret.Password{}
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}

	fmt.Printf("PostPassword:%+v\n", *p)
	password := p.Password
	qq := p.QQ

	value := secret.GetGlobalPersonValue("Password", p.QQ, &secret.Password{QQ: p.QQ, Password: ""}).(*secret.Password)
	if value.QQ == 0 && len(value.Password) == 0 && len(password) > 0 {
		value.Password = password
		secret.SetGlobalPersonValue("Password", 0, value)
		return response(e, true, err)
	}

	if qq == value.QQ && password == value.Password {
		return response(e, true, err)
	}
	return response(e, "密码错误", err)
}

func PostSuperMaster(e echo.Context) (err error) {
	qqStr := e.FormValue("supermaster")
	password := e.FormValue("password")
	if !verifyPassword(0, password) {
		return response(e, "密码错误", err)
	}

	qq, _ := strconv.ParseUint(qqStr, 10, 64)

	cfgSuper := secret.GetGlobalValue("Supermaster", &secret.Config{}).(*secret.Config)
	cfgSuper.HaveMaster = true
	cfgSuper.MasterQQ = qq
	secret.SetGlobalValue("Supermaster", cfgSuper)

	return response(e, true, err)
}

func verifyPassword(qq uint64, password string) bool {
	value := secret.GetGlobalPersonValue("Password", qq, &secret.Password{QQ: 0, Password: ""}).(*secret.Password)
	return password == value.Password
}

func PostDelay(e echo.Context) (err error) {
	delayStr := e.FormValue("delay")
	password := e.FormValue("password")
	if !verifyPassword(0, password) {
		return response(e, "密码错误", err)
	}

	delay, _ := strconv.ParseUint(delayStr, 10, 64)

	gp := secret.GetGlobalValue("ReplyDelay", &secret.ReplyDelay{DelayMs: 300}).(*secret.ReplyDelay)
	gp.DelayMs = delay

	secret.SetGlobalValue("ReplyDelay", gp)

	return response(e, true, err)
}

func PostImageMode(e echo.Context) (err error) {
	p := new(secret.ImgMode)
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}

	secret.SetGlobalValue("ImgMode", p)
	return response(e, true, err)
}

func PostTextSegment(e echo.Context) (err error) {
	p := new(secret.FoldLineMode)
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}

	secret.SetGlobalValue("FoldLineMode", p)
	return response(e, true, err)
}

type MoneyBindPost struct {
	Group      uint64 `json:"group" xml:"group" form:"group" query:"group"`
	IniPath    string `json:"iniPath" xml:"iniPath" form:"iniPath" query:"iniPath"`
	IniSection string `json:"iniSection" xml:"iniSection" form:"iniSection" query:"iniSection"`
	IniKey     string `json:"iniKey" xml:"iniKey" form:"iniKey" query:"iniKey"`
	HasUpdate  bool   `json:"hasUpdate" xml:"hasUpdate" form:"hasUpdate" query:"hasUpdate"`
	Encode     string `json:"encode" xml:"encode" form:"encode" query:"encode"`
}

func PostMoneyMap(e echo.Context) (err error) {
	p := new(MoneyBindPost)
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}

	pp := &secret.MoneyBind{}
	pp.IniPath = p.IniPath
	pp.IniSection = p.IniSection
	pp.IniKey = p.IniKey
	pp.HasUpdate = p.HasUpdate
	pp.Encode = p.Encode

	secret.SetMoneyMap(p.Group, pp)

	return response(e, true, err)
}

func PostActivities(e echo.Context) (err error) {
	p := []*secret.Activity{}
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}
	secret.SetActivities(p)
	return response(e, true, err)
}

type ChatMsg struct {
	QQ       uint64 `json:"qq" xml:"qq" form:"qq" query:"qq"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Msg      string `json:"msg" xml:"msg" form:"msg" query:"msg"`
}

type ChatReply struct {
	Msg string
}

type PersonState struct {
	Group  uint64
	Switch bool
	State  uint64 // 0:Start, 1:GroupSelect, 2:Play
}

func GetGroupNickName(info *cqp.GroupMember) string {
	if len(info.Card) > 0 {
		return info.Card
	}

	return info.Name
}

func GetGroupMemberInfo(group, qq int64, noCatch bool) cqp.GroupMember {
	return cqp.GetGroupMemberInfo(group, qq, noCatch)
}

func PostChat(e echo.Context) (err error) {
	p := &ChatMsg{}
	if err := e.Bind(p); err != nil {
		return response(e, "数据错误", err)
	}

	if !verifyPassword(p.QQ, p.Password) {
		return response(e, &ChatReply{Msg: "口令错误"}, err)
	}

	msg := p.Msg
	fromQQ := int64(p.QQ)
	state := secret.GetGlobalPersonValue("State", uint64(fromQQ), &PersonState{0, false, 0}).(*PersonState)

	fromGroup := int64(state.Group)

	info := GetGroupMemberInfo(fromGroup, fromQQ, false)
	selfQQ := cqp.GetLoginQQ()
	selfInfo := GetGroupMemberInfo(fromGroup, selfQQ, false)
	bot := secret.NewSecretBot(uint64(selfQQ), uint64(fromGroup), selfInfo.Name, true, &interact.Interact{})
	ret := ""

	update := func() {
		if len(msg) > 9 {
			fmt.Println(msg, "大于3", len(msg))
			ret = bot.Update(uint64(fromQQ), GetGroupNickName(&info))
		} else {
			fmt.Println(msg, "小于3", len(msg))
		}
	}

	update()
	ret = bot.RunPrivate(msg, uint64(fromQQ), GetGroupNickName(&info))

	fmt.Printf("\nSend private msg:%d, %s\n", fromGroup, ret)

	return response(e, &ChatReply{Msg: ret}, err)
}

type GroupOperate struct {
	Group    uint64 `json:"group" xml:"group" form:"group" query:"group"`
	Value    bool   `json:"value" xml:"value" form:"value" query:"value"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}

func PostGlobalSwitch(e echo.Context) (err error) {
	p := &GroupOperate{}
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}

	if verifyPassword(0, p.Password) {
		return response(e, "密码错误", err)
	}

	sw := secret.GetGlobalValue("GlobalSwitch", &secret.GlobalSwitch{Enable: true}).(*secret.GlobalSwitch)
	if sw.Enable != p.Value {
		sw.Enable = p.Value
		secret.SetGlobalValue("GlobalSwitch", sw)
	}

	return response(e, true, err)
}

func PostGlobalSilent(e echo.Context) (err error) {
	p := &GroupOperate{}
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}
	if verifyPassword(0, p.Password) {
		return response(e, "密码错误", err)
	}

	si := secret.GetGlobalValue("GlobalSilence", &secret.GlobalSilence{}).(*secret.GlobalSilence)
	if si.Enable != p.Value {
		si.Enable = p.Value
		secret.SetGlobalValue("GlobalSilence", si)
	}

	return response(e, true, err)
}

func PostGroupSwitch(e echo.Context) (err error) {
	p := &GroupOperate{}
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}
	if verifyPassword(0, p.Password) {
		return response(e, "密码错误", err)
	}

	b := &secret.Bot{}
	b.Group = p.Group
	b.SetSwitch(p.Value)
	return response(e, true, err)
}

func PostGroupSilent(e echo.Context) (err error) {
	p := &GroupOperate{}
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}
	if verifyPassword(0, p.Password) {
		return response(e, "密码错误", err)
	}

	b := &secret.Bot{}
	b.Group = p.Group

	s := b.GetGroupValue("Silence", &secret.SilenceState{}).(*secret.SilenceState)
	if p.Value {
		s.IsSilence = true
		s.OpenEndTime = "23:59"
		s.OpenStartTime = "23:58"
	} else {
		s.IsSilence = false
		b.SetGroupValue("Silence", s)
	}
	return response(e, true, err)
}

func PostGroupExit(e echo.Context) (err error) {
	p := &GroupOperate{}
	if err := e.Bind(p); err != nil {
		return response(e, false, err)
	}
	if verifyPassword(0, p.Password) {
		return response(e, "密码错误", err)
	}

	cqp.SetGroupLeave(int64(p.Group), false)
	return response(e, true, err)
}
