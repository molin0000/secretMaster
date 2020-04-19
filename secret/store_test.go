package secret

import (
	"testing"

	"github.com/molin0000/secretMaster/qlog"
)

type debugInteract struct{}

func (i *debugInteract) IsGroupAdmin(group, qq uint64) bool {
	return true
}

func (i *debugInteract) IsGroupMaster(group, qq uint64) bool {
	return true
}

func TestBuyItem(t *testing.T) {
	fromQQ := uint64(111)
	b := NewSecretBot(1234, 4567, "aa", false, &debugInteract{})
	b.clearItem(fromQQ)

	for i := 0; i < 10; i++ {
		b.setMoney(fromQQ, 1000)
		b.buyMagicItem(fromQQ)
		qlog.Println(b.getItems(fromQQ))
	}

	for i := 0; i < 10; i++ {
		b.setMoney(fromQQ, 1000)
		b.buyMace(fromQQ)
		qlog.Println(b.getItems(fromQQ))
	}
}
