package secret

import (
	"testing"

	"github.com/molin0000/secretMaster/qlog"
)

func TestFight(t *testing.T) {
	b := NewSecretBot(1234, 4567, "aa", false, &debugInteract{})
	qlog.Println(b.getBattleField())
}
