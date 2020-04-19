package secret

import (
	"testing"

	"github.com/molin0000/secretMaster/qlog"
)

func TestGetProperty(t *testing.T) {
	fromQQ := uint64(121)
	b := NewSecretBot(1234, 4567, "aa", false, &debugInteract{})

	qlog.Println(b.getProperty(fromQQ))
}
