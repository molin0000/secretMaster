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
	EventCnt          uint64
	AdventureLog      string
	LevelState        string
}

func NewPetStore() *PetStore {
	return &PetStore{}
}

func atoi(msg string) uint64 {
	n, _ := strconv.Atoi(msg)
	return uint64(n)
}

func (ps *PetStore) LoadPets(path string) {
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
		fmt.Printf("RealPet:%+v\n", *pet)
		realPets = append(realPets, pet)
	}
	ps.RealPets = realPets

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
		fmt.Printf("SpiritPet:%+v\n", *pet)
		spiritPets = append(spiritPets, pet)
	}
	ps.SpiritPets = spiritPets
}

func (ps *PetStore) GetStorePets() string {
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

func (ps *PetStore) State(pet *Pet) string {
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

	adv := "未探险"
	if len(pet.AdventureState) > 0 {
		adv = pet.AdventureState
	}

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
生命：%d
攻击：%d
防御: %d
闪避：%d
魅力：%d
技能：%s
探险：%s(事件：%d)
========`,
		pet.Nick, pet.Class, stars, pet.Level, FightLevel[myFightIndex]+sReLive,
		pet.Exp, pet.LevelState, pet.Money, pet.HPNow, pet.Attack, pet.Defense, pet.Speed, pet.Charm, pet.Skill, adv, pet.EventCnt,
	)
	return info
}

func (ps *PetStore) LevelExp(level uint64) uint64 {
	//exp = 100+1.2^level + level*100
	return uint64(100) + uint64(math.Pow(float64(1.2), float64(level))) + level*100
}

func (ps *PetStore) LevelUp(pet *Pet) string {
	if pet.Exp < ps.LevelExp(pet.Level) {
		return "你的宠物未达到升级标准"
	}

	pet.Level++
	if pet.Level%20 == 0 {
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
