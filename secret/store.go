package secret

import "time"

func (b *Bot) redTheater(fromQQ uint64) string {
	if b.getMoney(fromQQ) < 100 {
		return "你站在红剧场售票厅，盯着票价，攥紧了空荡荡的口袋，穷的掩面而去。"
	}

	if b.getMagic(fromQQ) < 40 {
		return "演出才刚刚开始，你突然一阵头晕目眩，感觉到灵力枯竭预警，立刻退票离开了。"
	}

	b.setMoney(fromQQ, -100)
	b.setMagic(fromQQ, -39)
	b.setExp(fromQQ, 100)
	return "你看了一场酣畅淋漓的演出，感觉自己对这个世界的了解更加深刻了。经验+100"
}

func (b *Bot) buyMagicPotion(fromQQ uint64) string {
	potion := b.getPersonValue("Potion", fromQQ, &Potion{0, uint64(time.Now().Unix() / 3600 / 24)})

	if potion.(*Potion).Days != uint64(time.Now().Unix()/3600/24) {
		potion.(*Potion).DayCnt = 0
		potion.(*Potion).Days = uint64(time.Now().Unix() / 3600 / 24)
	}

	if b.getMoney(fromQQ) < 70 {
		return "你盯着商品报价，攥紧了空荡荡的钱包，穷的掩面而去。"
	}

	if potion.(*Potion).DayCnt >= 3 && potion.(*Potion).Days == uint64(time.Now().Unix()/3600/24) {
		return "你今天喝了太多，实在是喝不下了。"
	}

	potion.(*Potion).DayCnt++

	b.setMagic(fromQQ, 30)
	b.setMoney(fromQQ, -70)
	b.setPersonValue("Potion", fromQQ, potion)

	return "喝下这瓶药剂后，你满嘴苦涩，灵力似乎恢复了一点。"
}

func (b *Bot) buyMagicItem(fromQQ uint64) string {
	if b.getMoney(fromQQ) < 50 {
		return "你盯着商品报价，攥紧了空荡荡的钱包，穷的掩面而去。"
	}

	bag := b.getPersonValue("Bag", fromQQ, &Bag{})
	find := false
	for i := 0; i < len(bag.(*Bag).Items); i++ {
		if bag.(*Bag).Items[i].Name == "灵力材料" {
			find = true
			bag.(*Bag).Items[i].Count++
			break
		}
	}

	if !find {
		bag.(*Bag).Items = append(bag.(*Bag).Items, &Item{"灵力材料", 1})
	}
	b.setPersonValue("Bag", fromQQ, bag)
	b.setMoney(fromQQ, -50)
	return "你精心挑选了一份外表精致的灵力材料，把它买回家。"
}

func (b *Bot) clearItem(fromQQ uint64) {
	b.setPersonValue("Bag", fromQQ, &Bag{})
}

func (b *Bot) buyMace(fromQQ uint64) string {
	if b.getMoney(fromQQ) < 1000 {
		return "你盯着商品报价，攥紧了空荡荡的钱包，穷的掩面而去。"
	}

	bag := b.getPersonValue("Bag", fromQQ, &Bag{})
	find := false
	for i := 0; i < len(bag.(*Bag).Items); i++ {
		if bag.(*Bag).Items[i] != nil && bag.(*Bag).Items[i].Name == "至高权杖" {
			find = true
			bag.(*Bag).Items[i].Count++
			break
		}

		if bag.(*Bag).Items[i] == nil {
			if len(bag.(*Bag).Items) > 1 {
				bag.(*Bag).Items[i] = bag.(*Bag).Items[len(bag.(*Bag).Items)-1]
				bag.(*Bag).Items = bag.(*Bag).Items[:len(bag.(*Bag).Items)-1]
			} else {
				bag.(*Bag).Items = nil
			}
		}
	}

	if !find {
		bag.(*Bag).Items = append(bag.(*Bag).Items, &Item{"至高权杖", 1})
	}
	b.setPersonValue("Bag", fromQQ, bag)
	b.setMoney(fromQQ, -1000)
	return "你精心挑选了一份外表精致的至高权杖，把它买回家。"
}

func (b *Bot) buyItem(fromQQ uint64, item string, price uint64) string {
	if b.getMoney(fromQQ) < price {
		return "你盯着商品报价，攥紧了空荡荡的钱包，穷的掩面而去。"
	}

	bag := b.getPersonValue("Bag", fromQQ, &Bag{})
	find := false
	for i := 0; i < len(bag.(*Bag).Items); i++ {
		if bag.(*Bag).Items[i] != nil && bag.(*Bag).Items[i].Name == item {
			find = true
			bag.(*Bag).Items[i].Count++
			break
		}

		if bag.(*Bag).Items[i] == nil {
			if len(bag.(*Bag).Items) > 1 {
				bag.(*Bag).Items[i] = bag.(*Bag).Items[len(bag.(*Bag).Items)-1]
				bag.(*Bag).Items = bag.(*Bag).Items[:len(bag.(*Bag).Items)-1]
			} else {
				bag.(*Bag).Items = nil
			}
		}
	}

	if !find {
		bag.(*Bag).Items = append(bag.(*Bag).Items, &Item{item, 1})
	}
	b.setPersonValue("Bag", fromQQ, bag)
	b.setMoney(fromQQ, -1*int(price))
	return "你精心挑选了一份外表精致的" + item + "，把它买回家。"
}

func (b *Bot) useItem(fromQQ uint64, name string) bool {
	bag := b.getPersonValue("Bag", fromQQ, &Bag{}).(*Bag)
	if len(bag.Items) == 0 {
		return false
	}

	for i := 0; i < len(bag.Items); i++ {
		if bag.Items[i].Name == name {
			if bag.Items[i].Count > 1 {
				bag.Items[i].Count--
			} else {
				if len(bag.Items) > 1 {
					bag.Items[i] = bag.Items[len(bag.Items)-1]
					bag.Items = bag.Items[:len(bag.Items)-1]
				} else {
					bag.Items = nil
				}
			}
			b.setPersonValue("Bag", fromQQ, bag)
			return true
		}
	}
	return false
}
