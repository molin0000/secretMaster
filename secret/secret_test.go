package secret

import (
	"testing"

	"github.com/molin0000/secretMaster/qlog"
)

func TestRun(t *testing.T) {
	b := NewSecretBot(3334, 333, "cat", false, &debugInteract{})
	qlog.Println(b.Update(112, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 自杀", 112, "mm"))
	qlog.Println(b.Update(112, "mm"))

	qlog.Println(b.Run("[CQ:at,qq=3334] 属性", 67939461, "mm"))
	qlog.Println(b.getProperty(112))

	for i := 0; i < 2000; i++ {
		qlog.Println(b.Update(112, "mm"))
		qlog.Println(b.Run("[CQ:at,qq=3334] 水水水水", 67939461, "mm"))
	}
}

func TestCmds(t *testing.T) {
	b := NewSecretBot(3334, 333, "cat0101", false, &debugInteract{})
	qlog.Println(b.Run("@cat 属性", 112, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 帮助", 67939461, "mm"))

	for i := 0; i < len(menus.SubMenu); i++ {
		qlog.Println(b.Run("[CQ:at,qq=3334] "+menus.SubMenu[i].Title, 67939461, "mm"))
		for m := 0; m < len(menus.SubMenu[i].SubMenu); m++ {
			qlog.Println(b.Run("[CQ:at,qq=3334] "+menus.SubMenu[i].SubMenu[m].Title, 67939461, "mm"))
		}
	}
}
