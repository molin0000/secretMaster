package backend

import (
	"fmt"
	"testing"
	"time"

	"github.com/molin0000/secretMaster/qlog"
	"github.com/molin0000/secretMaster/secret"
)

func testGetGroupInfoList() []*GroupInfo {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			qlog.Println("getGroupList panic")
		}
	}()

	gps := make([]*GroupInfo, 0)
	key := uint64(0)
	for i := 0; i < 10; i++ {
		gi := &GroupInfo{}
		gi.Key = key
		key++
		gi.Group = uint64(10000000 + i)
		gi.Member = fmt.Sprintf("%d/%d", i*10, i*200)
		b := &secret.Bot{}
		b.Group = uint64(10000000 + i)
		gi.Master = b.GetMaster()
		gi.Silence = b.IsSilent()
		gi.Switch = b.GetSwitch()
		gps = append(gps, gi)
	}

	return gps
}

func TestStartServer(t *testing.T) {
	StartServer(testGetGroupInfoList)

	d := time.Duration(time.Second * 20)

	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		<-ticker.C
		qlog.Println("20s passed")
	}
}
