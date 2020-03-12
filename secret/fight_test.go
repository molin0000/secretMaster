package secret

import (
	"fmt"
	"testing"
)

func TestFight(t *testing.T) {
	b := NewSecretBot(1234, 4567, "aa", false, &debugInteract{})
	fmt.Println(b.getBattleField())
}
