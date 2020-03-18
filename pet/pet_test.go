package pet

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPet(t *testing.T) {
	ps := NewPetStore()
	ps.LoadPets("/Users/molin/coolq/data/app/me.cqp.molin.secretMaster/pets.xlsx")
	fmt.Println(ps.GetStorePets())

	fmt.Println(ps.Buy(1000, "布偶猫"))
	fmt.Println(ps.Buy(500, "美人鱼"))
	fmt.Println(ps.Buy(1000, "美人鱼"))
	fmt.Println(ps.Buy(10000, "战争天使"))
	fmt.Println(ps.GetStorePets())

	fmt.Println(ps.Buy(1000, "金毛犬"))
	pet, msg := ps.Buy(10000, "六翼石像鬼")
	fmt.Println(msg)

	fmt.Println(ps.State(pet))

	fmt.Println(ps.GetStorePets())

	fmt.Println(ps.LevelUp(pet))
	fmt.Println(ps.LevelUp(pet))
	fmt.Println(ps.LevelUp(pet))
	fmt.Println(ps.LevelUp(pet))
	fmt.Println(ps.LevelUp(pet))

	for i := 0; i < 30; i++ {
		pet.Exp += uint64(rand.Intn(1000))
		fmt.Println(ps.LevelUp(pet))
		fmt.Println(ps.State(pet))
	}

	cnt := 0
	for i := 0; i < 300; i++ {
		pet, msg = ps.Summon()
		fmt.Println(msg)
		if pet != nil {
			cnt++
		}
	}
	fmt.Println(cnt)
}
