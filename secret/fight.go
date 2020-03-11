package secret

import "fmt"

func (b *Bot) getBattleField() string {
	b.updateBattleField()

	bf := b.getGroupValue("BattleField", &BattleField{}).(*BattleField)
	info := "\n==================\n值夜者小队阵营："

	for i, v := range bf.Nightwatches {
		state := ""
		if v.State == 0 {
			state = "正常"
		} else {
			state = "休整"
		}
		info += fmt.Sprintf("\n%d) %s (%s) 金镑%d, 经验%d", i, v.Nick, state, v.Money, v.Exp)
	}

	info += "\n==================\n赏金猎人阵营："
	for i, v := range bf.MoneyHunter {
		state := ""
		if v.State == 0 {
			state = "正常"
		} else {
			state = "休整"
		}
		info += fmt.Sprintf("\n%d) %s (%s) 金镑%d, 经验%d", i, v.Nick, state, v.Money, v.Exp)
	}
	info += "\n=================="

	return ""
}

func (b *Bot) updateBattleField() {

}

func (b *Bot) joinBattleField() {

}
