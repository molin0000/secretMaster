package pet

import (
	"fmt"
	"math/rand"
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
