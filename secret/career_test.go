package secret

import (
	"testing"

	"github.com/molin0000/secretMaster/qlog"
)

func TestSkill(t *testing.T) {
	fromQQ := uint64(111)
	b := NewSecretBot(1234, 4567, "aa", false, &debugInteract{})
	b.clearSkill(fromQQ)

	qlog.Println(b.getSkill(fromQQ))

	b.setSkill(fromQQ, 0, 1)
	b.setSkill(fromQQ, 1, 1)
	b.setSkill(fromQQ, 2, 1)

	qlog.Println(b.getSkill(fromQQ))

	b.skillLevelUp(fromQQ, 0)
	b.skillLevelUp(fromQQ, 1)
	b.skillLevelUp(fromQQ, 2)
	b.skillLevelUp(fromQQ, 0)
	b.skillLevelUp(fromQQ, 0)
	b.skillLevelUp(fromQQ, 0)
	b.skillLevelUp(fromQQ, 0)
	b.skillLevelUp(fromQQ, 0)

	qlog.Println(b.getSkill(fromQQ))
	qlog.Println(b.allSkillLevelUp(fromQQ))
	qlog.Println(b.getSkill(fromQQ))

	b.clearSkill(fromQQ)
	b.skillLevelUp(fromQQ, 1)

	qlog.Println(b.getSkill(fromQQ))
}

func TestPromotion(t *testing.T) {
	fromQQ := uint64(111)
	b := NewSecretBot(1234, 4567, "aa", false, &debugInteract{})
	qlog.Println(b.deletePerson(fromQQ))
	b.Update(fromQQ, "ThinkCat")
	qlog.Println(b.promotion(fromQQ))
	qlog.Println(b.changeSecretList("更换1", fromQQ))
	god := uint64(123)
	b.setGodToDb(0, &god)
	b.Update(fromQQ, "ThinkCat")
	b.setExp(fromQQ, 101)
	qlog.Println(b.promotion(fromQQ))
	b.setMoney(fromQQ, 200)
	b.buyMagicItem(fromQQ)
	qlog.Println(b.promotion(fromQQ))
	b.setExp(fromQQ, 101)
	b.setMoney(fromQQ, 200)
	b.buyMagicItem(fromQQ)
	qlog.Println(b.getMoney(fromQQ), b.getExp(fromQQ))
	qlog.Println(b.promotion(fromQQ))

	for i := 0; i < 500; i++ {
		b.setMoney(fromQQ, 200)
		b.setExp(fromQQ, 101)

		qlog.Println(b.getMoney(fromQQ), b.getExp(fromQQ))
		qlog.Println(b.promotion(fromQQ))
	}
	god = uint64(0)
	b.setGodToDb(0, &god)
	for i := 0; i < 500; i++ {
		b.setMoney(fromQQ, 200)
		b.setExp(fromQQ, 101)

		qlog.Println(b.getMoney(fromQQ), b.getExp(fromQQ))
		qlog.Println(b.promotion(fromQQ))
	}
}
