package secret

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

func GetActivities() []*Activity {
	return GetGlobalValue("Activities", []*Activity{}).([]*Activity)
}

func SetActivities(activites []*Activity) {
	SetGlobalValue("Activities", activites)
}
