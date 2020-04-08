package backend

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/molin0000/secretMaster/competition"
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

// func Get(e echo.Context) (err error) {
// 	// data := {}
// 	// return response(e, data, err)
// }

// func Post(e echo.Context) (err error) {
// 	// var market string
// 	// err = e.Bind(&market)
// 	// if err != nil {
// 	// 	utils.Debugf("bind param error: %v, params:%v", err, e.Request().Body)
// 	// 	return response(e, nil, err)
// 	// }
// 	// err = models.MarketDao.InsertMarket(&market)
// 	// return response(e, nil, err)
// }

// func GetOrdersHandler(e echo.Context) (err error) {
// 	var req struct {
// 		Address  string `json:"address"   query:"address"   validate:"required"`
// 		MarketID string `json:"market_id" query:"market_id" validate:"required"`
// 		Status   string `json:"status"    query:"status"`
// 		Offset   int    `json:"offset"    query:"offset"`
// 		Limit    int    `json:"limit "    query:"limit"`
// 	}

// 	var orders []*models.Order
// 	var count int64

// 	err = e.Bind(&req)
// 	if err == nil {
// 		count, orders = models.OrderDao.FindByAccount(req.Address, req.MarketID, req.Status, req.Offset, req.Limit)
// 	}

// 	return response(e, map[string]interface{}{"count": count, "orders": orders}, err)
// }

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
	Key     uint64 `json:"key"     `
	Group   uint64 `json:"group"   `
	Member  string `json:"member"  `
	Master  uint64 `json:"master"  `
	Switch  bool   `json:"switch"  `
	Silence bool   `json:"silence" `
}

var GetGroupInfoList func() []*GroupInfo

type GetGroupRet struct {
	GlobalSwitch  bool         `json:"globalSwitch"     `
	GlobalSilence bool         `json:"globalSilence"     `
	Groups        []*GroupInfo `json:"groups"     `
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
	qqStr := e.FormValue("qq")
	qq, _ := strconv.ParseUint(qqStr, 10, 64)
	password := e.FormValue("password")

	value := secret.GetGlobalPersonValue("Password", qq, &secret.Password{QQ: qq, Password: ""}).(*secret.Password)
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
	if !verifyAdminPassword(password) {
		return response(e, "密码错误", err)
	}

	qq, _ := strconv.ParseUint(qqStr, 10, 64)

	cfgSuper := secret.GetGlobalValue("Supermaster", &secret.Config{}).(*secret.Config)
	cfgSuper.HaveMaster = true
	cfgSuper.MasterQQ = qq
	secret.SetGlobalValue("Supermaster", cfgSuper)

	return response(e, true, err)
}

func verifyAdminPassword(password string) bool {
	value := secret.GetGlobalPersonValue("Password", 0, &secret.Password{QQ: 0, Password: ""}).(*secret.Password)
	return password == value.Password
}

func PostDelay(e echo.Context) (err error) {
	delayStr := e.FormValue("delay")
	password := e.FormValue("password")
	if !verifyAdminPassword(password) {
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
	return response(e, true, err)
}

func PostChat(e echo.Context) (err error) {
	return response(e, true, err)
}

func PostGroup(e echo.Context) (err error) {
	return response(e, true, err)
}
