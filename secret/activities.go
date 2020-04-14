package secret

import (
	"fmt"
	"time"

	"github.com/molin0000/secretMaster/qlog"
)

type Activity struct {
	Key     uint64 `json:"key" xml:"key" form:"key" query:"key"`
	KeyWord string `json:"keyWord" xml:"keyWord" form:"keyWord" query:"keyWord"`
	Reply   string `json:"reply" xml:"reply" form:"reply" query:"reply"`
	Type    string `json:"type" xml:"type" form:"type" query:"type"`
	Exp     uint64 `json:"exp" xml:"exp" form:"exp" query:"exp"`
	Money   uint64 `json:"money" xml:"money" form:"money" query:"money"`
	Magic   uint64 `json:"magic" xml:"magic" form:"magic" query:"magic"`
	Enable  bool   `json:"enable" xml:"enable" form:"enable" query:"enable"`
}

type ActivityState struct {
	GlobalSwitch bool        `json:"globalSwitch" xml:"globalSwitch" form:"globalSwitch" query:"globalSwitch"`
	Activities   []*Activity `json:"activities" xml:"activities" form:"activities" query:"activities"`
}

func GetActivities() *ActivityState {
	return GetGlobalValue("Activities", &ActivityState{}).(*ActivityState)
}

func SetActivities(activites *ActivityState) {
	SetGlobalValue("Activities", activites)
}

func CheckActivities(msg string, qq int64, group int64) (found bool, money, exp, magic uint64, reply string) {
	qlog.Println("CheckActivities", msg)
	as := GetActivities()
	qlog.Printf("%+v", *as)
	if !as.GlobalSwitch {
		return false, 0, 0, 0, ""
	}

	for _, v := range as.Activities {
		qlog.Printf("%+v", *v)
		if v.Enable && (msg == v.KeyWord) {
			if RecordActivities(msg, v, uint64(qq), uint64(group)) {
				return true, v.Money, v.Exp, v.Magic, v.Reply + fmt.Sprintf("\n经验：%d，金镑：%d，灵力：%d", v.Exp, v.Money, v.Magic)
			}
			return false, 0, 0, 0, ""
		}
	}

	return false, 0, 0, 0, ""
}

type PersonalActivities struct {
	QQ    uint64
	Group uint64
	Date  uint64
}

func RecordActivities(msg string, act *Activity, qq uint64, group uint64) bool {
	b := &Bot{}
	b.Group = group
	pa := b.getPersonValue("PersonalActivities_"+msg, qq, &PersonalActivities{QQ: qq, Group: group, Date: 0}).(*PersonalActivities)
	if pa.Date == 0 {
		pa.Date = uint64(time.Now().Unix() / (3600 * 24))
		b.setPersonValue("PersonalActivities_"+msg, qq, pa)
		return true
	}

	nowDay := uint64(time.Now().Unix() / (3600 * 24))
	ret := false
	switch act.Type {
	case "每日":
		if nowDay != pa.Date {
			pa.Date = nowDay
			ret = true
		}
	case "每周":
		if (nowDay / 7) != (pa.Date / 7) {
			pa.Date = nowDay
			ret = true
		}
	case "每月":
		if (nowDay / 30) != (pa.Date / 30) {
			pa.Date = nowDay
			ret = true
		}
	case "每年":
		if (nowDay / 365) != (pa.Date / 365) {
			pa.Date = nowDay
			ret = true
		}
	case "每人一次":
		ret = false
	default:
		qlog.Println("Error type", act.Type)
	}

	if ret {
		b.setPersonValue("PersonalActivities_"+msg, qq, pa)
		return true
	}
	return false
}
