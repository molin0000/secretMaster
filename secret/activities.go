package secret

type Activity struct {
	Key     uint64
	KeyWord string
	Reply   string
	Type    string
	Exp     uint64
	Money   uint64
	Magic   uint64
	Enable  bool
}

func GetActivities() []*Activity {
	return GetGlobalValue("Activities", []*Activity{}).([]*Activity)
}

func SetActivities(activites []*Activity) {
	SetGlobalValue("Activities", activites)
}
