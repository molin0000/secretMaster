package secret

import (
	"github.com/molin0000/secretMaster/calculator"
	"github.com/molin0000/secretMaster/mission"
)

type Person struct {
	Group       uint64
	QQ          uint64
	Name        string
	JoinTime    uint64
	LastChat    uint64
	LevelDown   uint64
	SecretID    uint64
	SecretLevel uint64
	ChatCount   uint64
}

type Persons []Person

func (a Persons) Len() int           { return len(a) }
func (a Persons) Less(i, j int) bool { return a[i].ChatCount > a[j].ChatCount }
func (a Persons) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type SecretInfo struct {
	SecretName      string
	SecretLevelName [10]string
}

type Bot struct {
	QQ        uint64
	Group     uint64
	Name      string
	Rank      []uint64
	RankNames []string
	Private   bool

	CurrentNick string

	information  *Information
	career       *Career
	organization *Organization
	store        *Store
}

type WaterRule struct {
	Group  uint64
	QQ     uint64
	DayCnt uint64
	Days   uint64
}

type Money struct {
	Group uint64
	QQ    uint64
	Money uint64
}

type Adventure struct {
	Group  uint64
	QQ     uint64
	DayCnt uint64
	Days   uint64
}

type RespectName struct {
	Group uint64
	QQ    uint64
	Name  string
}

type ExternProperty struct {
	Luck     uint64
	Magic    uint64
	Days     uint64
	BaseLuck uint64
}

type AdventureEvent struct {
	// 类型：1）-钱＆经验5%；2）-经验10%；3）-钱10%；4）+钱＆经验10%；5）+钱15%；6）+经验15%；7）无事发生35%
	Type int
	// 概率：百分比1~100
	Probability int
	// 文案
	Messages []string
}

type Menu struct {
	// 菜单级别：0，1，2，3
	ID      int
	Title   string
	Info    string
	Commit  string
	SubMenu []Menu
}

type Information struct {
	b *Bot
}

type Career struct {
	b *Bot
}

type Organization struct {
	b *Bot
}

type Store struct {
	b *Bot
}

type BotSwitch struct {
	Group  uint64
	Enable bool
}

type MoneyBind struct {
	IniPath    string
	IniSection string
	IniKey     string
	HasUpdate  bool
}

type Config struct {
	HaveMaster bool
	MasterQQ   uint64
}

type Item struct {
	Name  string
	Count uint64
}

type Bag struct {
	Items []*Item
}

type Potion struct {
	DayCnt uint64
	Days   uint64
}

type Skill struct {
	ID       uint64
	Name     string
	Level    uint64
	MaxLevel uint64
}
type SkillTree struct {
	Skills []*Skill
}

type ChurchInfo struct {
	Name        string
	Commit      string
	Skills      []*Skill
	CreatorQQ   uint64
	CreatorNick string
	Money       uint64
	MaxMember   uint64
	Level       uint64
	CreateMoney uint64
	Members     uint64
}

type Churchs struct {
	ChurchList []*ChurchInfo
}

type Version struct {
	Name    string
	Version string
	Date    string
}

type PrayState struct {
	Date uint64
}

type DbUpdate struct {
	HasUpdate bool
}

type Bank struct {
	Amount uint64
	Date   uint64
}

type Work struct {
	ID         uint64
	Name       string
	MagicMinus uint64
	MoneyAdd   uint64
	ExpAdd     uint64
	Date       uint64
}

type Fish struct {
	ID       uint64
	Name     string
	Property uint64
	Money    uint64
}

type Groups struct {
	Groups []uint64
}

type MissionState struct {
	IsPlaying bool
	Ms        *mission.MissionGame
}

type CalcState struct {
	IsPlaying bool
	Calc      *calculator.CalcGame
}

type Medal struct {
	MedalCnt uint64
}

type BattleField struct {
	Nightwatches       []*BattleInfo
	MoneyHunter        []*BattleInfo
	NightwatchWinCnt   uint64
	NightwatchLoseCnt  uint64
	MoneyHunterWinCnt  uint64
	MoneyHunterLoseCnt uint64
}

type BattleInfo struct {
	QQ         uint64
	Nick       string
	State      uint64 // 0:normal, 1:failed, 2:exited
	FailTime   uint64
	Money      uint64
	Exp        uint64
	UpdateTime uint64
	ExitTime   uint64
	Speed      uint64
	WinCnt     uint64
	FailCnt    uint64
	FieldType  uint64
}

type SilenceState struct {
	IsSilence     bool
	OpenStartTime string
	OpenEndTime   string
}

type KillSelf struct {
	DayCnt uint64
	Days   uint64
}

type GroupMap struct {
	Mapped       bool
	MapFromGroup uint64
	MapToGroup   uint64
}
