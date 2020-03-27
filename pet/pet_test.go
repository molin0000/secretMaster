package pet

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestPet(t *testing.T) {
	ps := NewPetStore()
	// ps.LoadPets("/Users/molin/coolq/data/app/me.cqp.molin.secretMaster/pets.xlsx")
	fmt.Println(ps.GetStorePets())

	fmt.Println(ps.Buy(1000, "布偶猫"))
	fmt.Println(ps.Buy(500, "美人鱼"))
	fmt.Println(ps.Buy(1000, "美人鱼"))
	fmt.Println(ps.Buy(10000, "战争天使"))
	fmt.Println(ps.GetStorePets())

	fmt.Println(ps.Buy(1000, "金毛犬"))
	pet, msg := ps.Buy(10000, "六翼石像鬼")
	fmt.Println(msg)

	var list = []string{"AAA", "BBB", "CCC"}

	fmt.Println(ps.State(pet, list))

	fmt.Println(ps.GetStorePets())

	fmt.Println(ps.LevelUp(pet))
	fmt.Println(ps.LevelUp(pet))
	fmt.Println(ps.LevelUp(pet))
	fmt.Println(ps.LevelUp(pet))
	fmt.Println(ps.LevelUp(pet))

	for i := 0; i < 30; i++ {
		pet.Exp += uint64(rand.Intn(1000))
		fmt.Println(ps.LevelUp(pet))
		fmt.Println(ps.State(pet, list))
	}

	cnt := 0
	pet0 := pet
	for i := 0; i < 30; i++ {
		pet, msg = ps.Summon()
		fmt.Println(msg)
		if pet != nil {
			cnt++
		}
	}
	fmt.Println(cnt)

	pet = pet0

	pet.Charm = 10

	fmt.Println(ps.StartAdv(pet))
	fmt.Println(ps.StartAdv(pet))

	fmt.Println(ps.State(pet, list))

	fmt.Println(ps.StopAdv(pet, list))

	fmt.Println(ps.StartAdv(pet))
	pet.AdvStartTime -= 3600 * 24
	pet.AdvUpdateTime -= 3600 * 24
	fmt.Println(ps.State(pet, list))

	fmt.Println(ps.StopAdv(pet, list))

	fmt.Println(ps.State(pet, list))

	fmt.Println(ps.StartAdv(pet))
	pet.AdvStartTime -= 3600 * 24
	pet.AdvUpdateTime -= 3600 * 24
	fmt.Println(ps.State(pet, list))
	pet.WeakStartTime -= 4 * 24 * 3600
	fmt.Println(ps.StopAdv(pet, list))

}

// func TestIntn(t *testing.T) {
// 	fmt.Println(rand.Intn(1 / 2))
// }

func TestSplit(t *testing.T) {
	msg := `
@木棉 
==================
值夜者小队：(胜：3，负：6)
0) 风筝 (正常) 金镑5646, 经验2823
1) 永夜劇作家 (正常) 金镑4225, 经验2112
2) 木棉 (正常) 金镑343, 经验171
3) 之也ﺴ࿈ (正常) 金镑452, 经验226
==================
赏金猎人：(胜：6，负：3)
0) 沉默门徒丶 (正常) 金镑4419, 经验2209
1) 繁夏 (正常) 金镑331, 经验165
2) Cyan (正常) 金镑225, 经验112
==================
`
	strs := strings.Split(msg, "\n")
	length := len(strs)
	cnt := 0
	for {
		info := ""
		for i := 0; i < 5; i++ {
			if i < length {
				info += strs[cnt] + "\n"
				cnt++
			} else {
				break
			}
		}
		info = strings.TrimRight(info, "\n")
		fmt.Println(info)
		fmt.Println(cnt)
		if cnt >= length {
			break
		}
	}
}
