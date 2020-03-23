package pet

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type PetStore struct {
	RealPets        []*Pet
	SpiritPets      []*Pet
	StorePets       []*Pet
	StoreUpdateTime uint64
	PetFilePath     string
}

type Pet struct {
	Class             string
	Price             uint64
	From              string
	Commit            string
	Skill             string
	HP                uint64
	Attack            uint64
	Defense           uint64
	Speed             uint64
	Charm             uint64
	HPGrowthType      uint64
	AttackGrowthType  uint64
	DefenseGrowthType uint64
	SpeedGrowthType   uint64
	Nick              string
	Star              uint64
	Level             uint64
	Money             uint64
	Exp               uint64
	HPNow             uint64
	AdventureState    string
	AdvStartTime      uint64
	WeakStartTime     uint64
	AdvUpdateTime     uint64
	EventCnt          uint64
	AdventureLog      string
	LevelState        string
}

type Food struct {
	Name  string
	Exp   int
	Money int
}

var _realPets []*Pet
var _spiritPets []*Pet

func init() {
	LoadPets(PetFilePath)
}

func NewPetStore() *PetStore {
	ps := &PetStore{}
	ps.RealPets = _realPets
	ps.SpiritPets = _spiritPets
	return ps
}

func atoi(msg string) uint64 {
	n, _ := strconv.Atoi(msg)
	return uint64(n)
}

func LoadPets(path string) {
	fmt.Println("加载宠物表格...")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("加载宠物表格...失败")
			fmt.Println(err)
		}
	}()

	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows := f.GetRows("现世宠物")
	realPets := make([]*Pet, 0)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		pet := &Pet{
			Class:             row[0],
			Price:             atoi(row[1]),
			From:              row[2],
			Commit:            row[3],
			Skill:             row[4],
			HP:                atoi(row[5]),
			Attack:            atoi(row[6]),
			Defense:           atoi(row[7]),
			Speed:             atoi(row[8]),
			Charm:             atoi(row[9]),
			HPGrowthType:      atoi(row[10]),
			AttackGrowthType:  atoi(row[11]),
			DefenseGrowthType: atoi(row[12]),
			SpeedGrowthType:   atoi(row[12]),
			Nick:              row[0],
			Star:              1,
			Level:             1,
			Money:             0,
			Exp:               0,
			HPNow:             atoi(row[5]),
		}
		realPets = append(realPets, pet)
	}
	_realPets = realPets

	// Get all the rows in the Sheet2.
	rows = f.GetRows("灵界宠物")
	spiritPets := make([]*Pet, 0)
	for i, row := range rows {
		if i == 0 {
			continue
		}

		pet := &Pet{
			Class:             row[0],
			Price:             atoi(row[1]),
			From:              row[2],
			Commit:            row[3],
			Skill:             row[4],
			HP:                atoi(row[5]),
			Attack:            atoi(row[6]),
			Defense:           atoi(row[7]),
			Speed:             atoi(row[8]),
			Charm:             atoi(row[9]),
			HPGrowthType:      atoi(row[10]),
			AttackGrowthType:  atoi(row[11]),
			DefenseGrowthType: atoi(row[12]),
			SpeedGrowthType:   atoi(row[12]),
			Nick:              row[0],
			Star:              1,
			Level:             1,
			Money:             0,
			Exp:               0,
			HPNow:             atoi(row[5]),
		}
		spiritPets = append(spiritPets, pet)
	}
	_spiritPets = spiritPets

	// Get all the rows in the Sheet3.
	rows = f.GetRows("宠物投食")

	for i, row := range rows {
		if i == 0 {
			continue
		}
		food := Food{
			Name:  row[0],
			Exp:   int(atoi(row[1])),
			Money: int(atoi(row[2])),
		}

		foodList = append(foodList, food)
	}
	fmt.Printf("已加载:%d现世宠物，%d灵界宠物，%d宠物投食\n", len(realPets), len(spiritPets), len(foodList))
}

func (ps *PetStore) GetStorePets() string {
	fmt.Println("GetStorePets", len(ps.RealPets), len(ps.SpiritPets))
	if len(ps.RealPets) == 0 {
		if len(_realPets) == 0 {
			LoadPets(PetFilePath)
		}
		ps.RealPets = _realPets
	}

	if len(ps.SpiritPets) == 0 {
		ps.SpiritPets = _spiritPets
	}

	nowTime := uint64(time.Now().Unix())
	updateHour := uint64(4)
	petCnt := 4
	needFresh := false
	d := nowTime - ps.StoreUpdateTime
	if d > updateHour*3600 {
		needFresh = true
		d = 0
		ps.StoreUpdateTime = nowTime
	}
	info := fmt.Sprintf("\n欢迎光临，宠物货架每4小时自动刷新，刷新剩余：%d秒", 3600*updateHour-d)

	if needFresh {
		ps.StorePets = make([]*Pet, 0)
		fmt.Println("GetStorePets2", len(ps.RealPets), len(ps.SpiritPets))

		for i := 0; i < petCnt; i++ {
			n := rand.Intn(len(ps.RealPets))
			ps.StorePets = append(ps.StorePets, ps.RealPets[n])
		}
	}

	if len(ps.StorePets) == 0 {
		info += "\n对不起，宠物已经全被领养走了呢"
		return info
	}

	for _, p := range ps.StorePets {
		info += "\n---------------------\n"
		info += fmt.Sprintf("名称：%s\n价格：%d\n简介：%s", p.Nick, p.Price, p.Commit)
	}

	info += "\n---------------------\n"
	return info
}

func (ps *PetStore) Buy(money uint64, name string) (*Pet, string) {
	var pet *Pet
	for i, v := range ps.StorePets {
		if v.Class == name {
			if money < v.Price {
				return nil, "对不起，你的钱还不够领养这只宠物。（你的贫穷遭到了宠物们的鄙视）"
			}
			pet = ps.StorePets[i]
			if len(ps.StorePets) > 1 {
				ps.StorePets[i] = ps.StorePets[len(ps.StorePets)-1]
				ps.StorePets = ps.StorePets[:len(ps.StorePets)-1]
			} else {
				ps.StorePets = nil
			}

			return pet, fmt.Sprintf("在店长的鉴证下，你和%s达成了神圣的宠物契约，并把它带回了家。", v.Class)
		}
	}
	return nil, "店里没有你要的宠物。"
}

func (ps *PetStore) ChangeNick(nick string, pet *Pet) string {
	if len(nick) > 40 {
		return "超出最大长度(10/40)"
	}
	pet.Nick = nick
	return "改名成功：" + nick
}

func (ps *PetStore) State(pet *Pet, list []string) string {
	info := ""
	stars := ""

	for i := uint64(0); i < pet.Star; i++ {
		stars += "⭐️"
	}
	myFightIndex := pet.Level
	reLive := uint64(0)
	sReLive := ""
	if myFightIndex > 99 {
		myFightIndex = pet.Level % 100
		reLive = pet.Level / uint64(100)
		if reLive > 0 {
			sReLive = fmt.Sprintf(" - 转生+%d", reLive)
		}
	}

	adv := ps.checkPetState(pet, list)

	if pet.Exp > ps.LevelExp(pet.Level) {
		pet.LevelState = "(可晋级)"
	}

	info += fmt.Sprintf(`
=========
昵称：%s
种类：%s
品质：%s
等级：lv%d(%s)
经验：%d%s
金镑：%d
生命：%d/%d
攻击：%d
防御: %d
闪避：%d
魅力：%d
技能：%s
历练：%s(事件：%d)
========`,
		pet.Nick, pet.Class, stars, pet.Level, FightLevel[myFightIndex]+sReLive,
		pet.Exp, pet.LevelState, pet.Money, pet.HPNow, pet.HP, pet.Attack, pet.Defense, pet.Speed, pet.Charm, pet.Skill, adv, pet.EventCnt,
	)
	return info
}

func (ps *PetStore) LevelExp(level uint64) uint64 {
	if level > 100 {
		level100 := uint64(100) + uint64(math.Pow(float64(1.2), float64(100))) + 100*100
		return level100 + uint64(100) + uint64(math.Pow(float64(1.2), float64(level%101))) + level%101*100
	}
	//exp = 100+1.2^level + level*100
	return uint64(100) + uint64(math.Pow(float64(1.2), float64(level))) + level*100
}

func (ps *PetStore) LevelUp(pet *Pet) string {
	if pet.Exp < ps.LevelExp(pet.Level) {
		return "你的宠物未达到升级标准"
	}

	pet.Level++
	if pet.Level%20 == 0 && pet.Level <= 100 {
		pet.Star++
	}

	h := pet.Level*10 + uint64(rand.Intn(10))
	pet.HP += h
	pet.HPNow = pet.HP
	a := pet.Level + uint64(rand.Intn(5))
	d := pet.Level + uint64(rand.Intn(5))
	s := pet.Level + uint64(rand.Intn(5))

	pet.Attack += a
	pet.Defense += d
	pet.Speed += s
	pet.LevelState = ""
	return "宠物升级成功!" + fmt.Sprintf("生命+%d, 攻击+%d, 防御+%d, 敏捷+%d", h, a, d, s)
}

func (ps *PetStore) Summon() (*Pet, string) {
	baseNum := uint64(100000)
	totalNum := uint64(0)

	for i, p := range ps.SpiritPets {
		totalNum += uint64(float64(baseNum) / float64(p.Price))
		if p.HP == 0 || p.Attack == 0 {
			fmt.Println("err", i)
		}
	}

	fmt.Println(totalNum)
	scope := totalNum * 5
	r := uint64(rand.Intn(int(scope)))
	fmt.Println("R:", r)

	if r > totalNum {
		return nil, "你没找到宠物，还在灵界迷路，历尽千辛万苦，终于找到归途。"
	}

	totalNum = 0
	for i, p := range ps.SpiritPets {
		totalNum += uint64(float64(baseNum) / float64(p.Price))
		if int(r)-int(totalNum) <= 0 {
			return ps.SpiritPets[i], "你找到了宠物！" + ps.State(ps.SpiritPets[i], nil)
		}
	}

	return nil, "你没找到宠物，还在灵界迷路，历尽千辛万苦，终于找到归途。"
}

func (ps *PetStore) StartAdv(pet *Pet) string {
	if pet.AdvStartTime != 0 {
		return fmt.Sprintf("\n你的宠物%s已经在努力的探险了。", pet.Nick)
	}

	if pet.WeakStartTime != 0 {
		return fmt.Sprintf("\n%s: 主人，我都快死了，还探什么险啊", pet.Nick)
	}

	pet.AdvStartTime = uint64(time.Now().Unix())
	pet.AdvUpdateTime = uint64(time.Now().Unix())
	pet.HPNow = pet.HP
	pet.WeakStartTime = 0
	pet.EventCnt = 0
	pet.AdventureLog = ""
	return fmt.Sprintf("\n%s走上了历练探险的旅途，它一脸坚毅，充满了决心。", pet.Nick)
}

func (ps *PetStore) StopAdv(pet *Pet, list []string) (ret string, die bool) {
	info := pet.AdventureLog
	if pet.AdvStartTime == 0 {
		return info + fmt.Sprintf("\n%s: 主人，我在家呢，不用召回。", pet.Nick), false
	}

	ps.checkPetState(pet, list)

	pet.AdvStartTime = 0

	nowTime := uint64(time.Now().Unix())
	if pet.WeakStartTime != 0 && ((nowTime - pet.WeakStartTime) < 3600*24*3) {
		pet.WeakStartTime = 0
		pet.HPNow = pet.HP
		return info + fmt.Sprintf("\n%s: 主人，你终于来救我了，差一点就挂了，呜呜呜", pet.Nick), false
	}

	if pet.WeakStartTime != 0 {
		return info + fmt.Sprintf("\n%s走的很安详，你为它收敛了尸骨，并立了一个墓碑。", pet.Nick), true
	}

	pet.AdvStartTime = 0
	pet.WeakStartTime = 0
	pet.HPNow = pet.HP
	pet.AdvUpdateTime = 0
	pet.EventCnt = 0
	pet.AdventureLog = ""
	return info + "\n成功召回！", false
}

func (ps *PetStore) pk(pet *Pet, enemyType int) string {
	var enemyList []*Pet
	if enemyType == 0 {
		enemyList = ps.RealPets
	} else {
		enemyList = ps.SpiritPets
	}

	enemy := *enemyList[rand.Intn(len(enemyList))]

	info := "\n"
	r := rand.Intn(7) - 5
	if pet.Level < 4 && r < 0 {
		r = 0
	}
	enemy.Level = pet.Level + uint64(r)
	info += fmt.Sprintf("%s在探险的途中遭遇了%s(lv%d)\n", pet.Nick, enemy.Nick, enemy.Level)
	xp := rand.Int63n(int64(ps.LevelExp(pet.Level)))
	win := false
	if uint64(xp) < pet.Exp/4 || pet.Level < 10 {
		win = true
	} else {
		win = false
	}

	if win {
		e := pet.Level/2 + uint64(5+rand.Intn(int(pet.Level)/2+1))
		m := uint64(5 + rand.Intn(int(pet.Level)/2+1))
		h := uint64(rand.Intn(int(pet.Attack)))
		if pet.HPNow > h {
			info += fmt.Sprintf("%s战胜了%s(lv%d), HP:-%d，获得：%d经验，%d金镑", pet.Nick, enemy.Nick, enemy.Level, h, e, m)
			pet.Exp += e
			pet.Money += m
			pet.HPNow -= h
		} else {
			pet.HPNow = 1
			info += fmt.Sprintf("你失败了，进入濒死状态。HP:-%d", h)
			pet.WeakStartTime = uint64(time.Now().Unix())
		}
	} else {
		pet.HPNow = 1
		info += "你失败了，进入濒死状态。"
		pet.WeakStartTime = uint64(time.Now().Unix())
	}
	return info
}

func (ps *PetStore) getEvent(pet *Pet, list []string) string {
	if pet.WeakStartTime != 0 {
		return ""
	}
	pet.EventCnt++
	charm := pet.Charm + 5
	if charm > 50 {
		charm = 50
	}
	info := "\n"
	max := 75 + charm
	r := uint64(rand.Intn(int(max)))
	man := list[rand.Intn(len(list))]
	food := foodList[rand.Intn(len(foodList))]
	if r < charm {
		info += fmt.Sprintf("%s在路上偶遇了%s，%s摸了摸%s的头并投喂了%s。（经验：%d, 金镑：%d）", pet.Nick, man, man, pet.Nick, food.Name, food.Exp, food.Money)
		if food.Exp >= 0 {
			pet.Exp += uint64(food.Exp)
		} else {
			if pet.Exp > uint64(-1*food.Exp) {
				pet.Exp -= uint64(-1 * food.Exp)
			} else {
				pet.Exp = 0
			}
		}
		if food.Money >= 0 {
			pet.Money += uint64(food.Money)
		} else {
			if pet.Money > uint64(-1*food.Money) {
				pet.Money -= uint64(-1 * food.Money)
			} else {
				pet.Money = 0
			}
		}
		return info
	}

	if r < charm+25 {
		info += fmt.Sprintf("%s在探险的路上迷失了方向，历尽千辛万苦才找到了回家的路。（经验: +5)", pet.Nick)
		pet.Exp += 5
		return info
	}

	if r < charm+50 {
		return ps.pk(pet, 0)
	}

	if r <= charm+75 {
		return ps.pk(pet, 1)
	}
	return info + "遭遇了bug"
}

func (ps *PetStore) checkPetState(pet *Pet, list []string) string {
	//几种状态，迷路25%，投喂?%，现世pk25%，灵界pk25%
	if pet.AdvStartTime == 0 {
		return "未历练"
	}

	nowTime := uint64(time.Now().Unix())
	if pet.WeakStartTime != 0 && ((nowTime - pet.WeakStartTime) < 3600*24*3) {
		return "濒死"
	}

	if pet.WeakStartTime != 0 {
		return "死亡"
	}

	for {
		rTime := uint64(rand.Intn(60 * 10))
		if nowTime-pet.AdvUpdateTime > (60*10 + rTime) {
			pet.AdvUpdateTime += (60*10 + rTime)
			pet.AdventureLog += ps.getEvent(pet, list)
		} else {
			break
		}
	}

	if pet.WeakStartTime != 0 {
		return "濒死"
	}

	return "历练中"
}
