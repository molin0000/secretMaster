package secret

import (
	"testing"
	"time"

	"github.com/molin0000/secretMaster/qlog"
)

func TestTime(t *testing.T) {
	tm := time.Now()
	qlog.Println(tm.Hour(), tm.Minute())
}
