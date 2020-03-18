package secret

import (
	"fmt"
	"strings"

	"github.com/molin0000/secretMaster/pet"
)

func (b *Bot) petState(fromQQ uint64) string {
	pet := b.getPersonValue("Pet", fromQQ, &pet.Pet{}).(*pet.Pet)
	if len(pet.Class) == 0 {
		return "你还没有宠物"
	}

	ps := b.getPetStore()

	if len(b.RankNames) == 0 {
		b.getRank(fromQQ)
	}
	ret := ps.State(pet, b.RankNames)

	b.savePet(fromQQ, pet)
	return ret
}

func (b *Bot) getPet(fromQQ uint64) *pet.Pet {
	pet := b.getPersonValue("Pet", fromQQ, &pet.Pet{}).(*pet.Pet)
	if len(pet.Class) == 0 {
		return nil
	}
	return pet
}

func (b *Bot) savePet(fromQQ uint64, pet *pet.Pet) {
	b.setPersonValue("Pet", fromQQ, pet)
}

func (b *Bot) petRename(fromQQ uint64, msg string) string {
	pet := b.getPet(fromQQ)
	if pet == nil {
		return "你没有宠物"
	}

	if b.getMoney(fromQQ) < 100 {
		return "钱不够"
	}
	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return "格式不对"
	}
	b.setMoney(fromQQ, -100)

	ret := b.getPetStore().ChangeNick(strs[1], pet)
	b.savePet(fromQQ, pet)
	return ret
}

func (b *Bot) petStore() string {
	ps := b.getPetStore()
	ret := ps.GetStorePets()
	b.savePetStore(ps)
	return ret
}

func (b *Bot) petBuy(fromQQ uint64, msg string) string {
	p := b.getPet(fromQQ)
	if p != nil {
		return "你已经有宠物了，请一心一意待它。"
	}
	ps := b.getPetStore()

	strs := strings.Split(msg, ";")
	if len(strs) != 2 {
		return "格式错误"
	}
	pet, ret := ps.Buy(b.getMoney(fromQQ), strs[1])
	if pet == nil {
		return ret
	}
	b.savePet(fromQQ, pet)
	return ret
}

func (b *Bot) petSummon(fromQQ uint64) string {
	p := b.getPet(fromQQ)
	if p != nil {
		return "你已经有宠物了，请一心一意待它。"
	}

	if b.getMagic(fromQQ) < 20 {
		return "灵性不足"
	}

	b.setMagic(fromQQ, -20)
	have := b.useItem(fromQQ, "灵性材料")
	if !have {
		return "你缺少灵性材料"
	}
	ps := b.getPetStore()
	pet, ret := ps.Summon()
	if pet == nil {
		return ret
	}
	b.savePet(fromQQ, pet)
	return ret
}

func (b *Bot) petLevelUp(fromQQ uint64) string {
	ps := b.getPetStore()
	pet := b.getPet(fromQQ)
	ret := ps.LevelUp(pet)
	b.savePet(fromQQ, pet)
	return ret
}

func (b *Bot) petSell(fromQQ uint64) string {
	p := b.getPet(fromQQ)
	info := fmt.Sprintf("你释放了你的宠物%s, 它送给你%d金镑。", p.Nick, p.Money/2)
	b.savePet(fromQQ, &pet.Pet{})
	b.setMoney(fromQQ, int(p.Money/2))
	return info
}

func (b *Bot) petEat(fromQQ uint64) string {
	p := b.getPet(fromQQ)
	info := fmt.Sprintf("你融合了你的宠物%s, 得到%d经验。", p.Nick, p.Exp/5)
	b.savePet(fromQQ, &pet.Pet{})
	b.setExp(fromQQ, int(p.Exp/5))
	return info
}

func (b *Bot) petAdventure(fromQQ uint64) string {
	ps := b.getPetStore()
	p := b.getPet(fromQQ)
	ret := ps.StartAdv(p)
	b.savePet(fromQQ, p)
	return ret
}

func (b *Bot) petGoHome(fromQQ uint64) string {
	ps := b.getPetStore()
	p := b.getPet(fromQQ)
	if len(b.RankNames) == 0 {
		b.getRank(fromQQ)
	}
	ret, die := ps.StopAdv(p, b.RankNames)
	if die {
		b.savePet(fromQQ, &pet.Pet{})
	} else {
		b.savePet(fromQQ, p)
	}
	return ret
}

func (b *Bot) petBeauty(fromQQ uint64) string {
	pet := b.getPersonValue("Pet", fromQQ, &pet.Pet{}).(*pet.Pet)
	if len(pet.Class) == 0 {
		return "你还没有宠物"
	}

	if b.getMoney(fromQQ) < 500 {
		return "你钱不够"
	}

	b.setMoney(fromQQ, -500)

	pet.Charm++

	b.setPersonValue("Pet", fromQQ, pet)

	return "你的宠物变得美美哒、萌萌哒、闪闪惹人爱"
}

func (b *Bot) getPetStore() *pet.PetStore {
	ps := b.getGroupValue("PetStore", pet.NewPetStore()).(*pet.PetStore)
	return ps
}

func (b *Bot) savePetStore(ps *pet.PetStore) {
	b.setGroupValue("PetStore", ps)
}
