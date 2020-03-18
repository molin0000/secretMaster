package pet

import (
	"fmt"
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

}
