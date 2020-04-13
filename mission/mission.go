package mission

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/molin0000/secretMaster/config"
	"github.com/molin0000/secretMaster/qlog"
)

type Result struct {
	Key  json.Number `json:"key"`
	Type string      `json:"type"`
	Data json.Number `json:"data"`
}

type Option struct {
	Key     json.Number `json:"key"`
	Desc    string      `json:"desc"`
	Results []Result    `json:"results"`
}

type Event struct {
	Key     json.Number `json:"key"`
	Desc    string      `json:"desc"`
	Options []Option    `json:"options"`
}

type Mission struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Date   string  `json:"date"`
	Events []Event `json:"events"`
}

type MissionGame struct {
	Ms    *Mission
	Event uint64
	Money uint64
	Exp   uint64
	Magic uint64
	Msg   string
}

var missionList = []*Mission{}

func loadMissions(jsonPath string) {
	files, _ := ioutil.ReadDir(jsonPath)
	for _, f := range files {
		qlog.Println(f.Name())
		if !strings.Contains(f.Name(), ".json") {
			continue
		}
		str := readJSONFile(path.Join(jsonPath, f.Name()))
		var ms Mission
		err := json.Unmarshal([]byte(str), &ms)
		if err != nil {
			qlog.Println(err)
			continue
		}
		missionList = append(missionList, &ms)
	}
}

func GetMissionCount() int {
	if len(missionList) == 0 {
		loadMissions(config.MissionPath)
	}
	return len(missionList)
}

func readJSONFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return ""
	}
	return string(fd)
}

// NewRandomMission 随机获取一个新的任务副本
func NewRandomMission(jsonPath string) *MissionGame {
	if len(missionList) == 0 {
		loadMissions(jsonPath)
	}

	idx := rand.Intn(len(missionList))
	return GetMissionGame(idx)
}

func ListMission(jsonPath string) string {
	if len(missionList) == 0 {
		loadMissions(jsonPath)
	}

	info := "\n"
	for i, v := range missionList {
		info += fmt.Sprintf("%d) 副本：%s, 作者：%s, 时间：%s\n", i, v.Title, v.Author, v.Date)
	}
	return info
}

func GetMissionGame(idx int) *MissionGame {
	if idx >= len(missionList) {
		return nil
	}
	mg := &MissionGame{Ms: missionList[idx]}
	return mg
}

func (mg *MissionGame) ShowEvent(event int) string {
	msg := "\n"
	if event == 0 {
		msg = mg.GetInfo() + "\n--------------\n"
	}

	if event >= len(mg.Ms.Events) {
		return mg.failed(event, 0, "showEvent")
	}

	msg += mg.Ms.Events[event].Desc + "\n"
	for i, opt := range mg.Ms.Events[event].Options {
		msg += fmt.Sprintf("%d) %s\n", i, opt.Desc)
	}
	msg += "请行动(输入数字选择行动)：\n"
	return msg
}

func (mg *MissionGame) failed(event, selection int, tag string) string {
	return fmt.Sprintf("副本:%s, 事件:%d, 选项:%d, %s数值异常，副本强制退出", mg.Ms.Title, event, selection, tag)
}

func (mg *MissionGame) SelectOption(event int, selection int) (msg string, finish bool) {
	msg = "\n"

	if event >= len(mg.Ms.Events) {
		msg += mg.failed(event, selection, "event")
		return msg, true
	}

	if len(mg.Ms.Events[event].Options) == 0 {
		return msg, true
	}

	if selection >= len(mg.Ms.Events[event].Options) {
		msg += "选择数值异常，请重新选择\n"
		return msg, false
	}

	opt := mg.Ms.Events[event].Options[selection]

	if len(opt.Results) == 0 {
		return msg, true
	}
	hasRa := false
	raSuccess := false
	for _, v := range opt.Results {
		switch v.Type {
		case ".ra":
			hasRa = true
			n, err := strconv.Atoi(string(v.Data))
			if err != nil {
				msg += mg.failed(event, selection, ".ra data")
				return msg, true
			}

			r := rand.Intn(100)
			if r <= n {
				raSuccess = true
			} else {
				raSuccess = false
			}
			str := ""
			if raSuccess {
				str = "成功"
			} else {
				str = "失败"
			}
			msg += fmt.Sprintf(".ra%d = %d/%d %s\n", n, r, n, str)
		case "ra成功":
			if hasRa && raSuccess {
				n, err := strconv.Atoi(string(v.Data))
				if err != nil {
					msg += mg.failed(event, selection, v.Type)
					return msg, true
				}
				mg.Event = uint64(n)
				qlog.Println("ra成功", mg.Event)
				msg += mg.ShowEvent(int(mg.Event))
				return msg, false
			}
		case "ra失败":
			if hasRa && !raSuccess {
				n, err := strconv.Atoi(string(v.Data))
				if err != nil {
					msg += mg.failed(event, selection, v.Type)
					return msg, true
				}
				mg.Event = uint64(n)
				msg += mg.ShowEvent(int(mg.Event))
				return msg, false
			}
		case "描述":
			msg += string(v.Data) + "\n"
		case "金镑":
			n, err := strconv.Atoi(string(v.Data))
			if err != nil {
				msg += mg.failed(event, selection, v.Type)
				return msg, true
			}
			msg += fmt.Sprintf("金镑：%d\n", n)
			mg.Money = uint64(int64(mg.Money) + int64(n))
		case "经验":
			n, err := strconv.Atoi(string(v.Data))
			if err != nil {
				msg += mg.failed(event, selection, v.Type)
				return msg, true
			}
			msg += fmt.Sprintf("经验：%d\n", n)
			mg.Exp = uint64(int64(mg.Exp) + int64(n))
		case "灵性":
			n, err := strconv.Atoi(string(v.Data))
			if err != nil {
				msg += mg.failed(event, selection, v.Type)
				return msg, true
			}
			msg += fmt.Sprintf("灵力：%d\n", n)
			mg.Magic = uint64(int64(mg.Magic) + int64(n))
		case "跳转":
			n, err := strconv.Atoi(string(v.Data))
			if err != nil {
				msg += mg.failed(event, selection, v.Type)
				return msg, true
			}
			mg.Event = uint64(n)
			msg += mg.ShowEvent(int(mg.Event))
			return msg, false
		case "结束":
			msg += string(v.Data) + "\n"
			return msg, true
		}
	}

	return msg, true
}

func (mg *MissionGame) GetInfo() string {
	return fmt.Sprintf("副本：%s, 作者：%s, 时间：%s", mg.Ms.Title, mg.Ms.Author, mg.Ms.Date)
}

func (mg *MissionGame) Finish() string {
	mg.Event = 0
	return fmt.Sprintf("\n----------------\n副本结束! 金镑：%d, 经验：%d, 灵力: %d\n----------------\n", int64(mg.Money), int64(mg.Exp), int64(mg.Magic))
}
