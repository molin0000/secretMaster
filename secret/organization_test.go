package secret

import (
	"strings"
	"testing"

	"github.com/molin0000/secretMaster/qlog"
)

func checkResult(input string, compare string, t *testing.T) {
	qlog.Println(input)
	if strings.Trim(input, "\n") != strings.Trim(compare, "\n") {
		t.Errorf("%s\nexpected: %s", input, compare)
	}
}
func TestChurchs(t *testing.T) {
	fromQQ := uint64(67939461000)
	b := NewSecretBot(3334, 333, "aaa", false, &debugInteract{})

	checkResult(b.Run("[CQ:at,qq=3334] 自杀", fromQQ, "mm"), "人物删除成功", t)
	checkResult(b.Run("[CQ:at,qq=3334] 自杀", fromQQ+1, "mm"), "人物删除成功", t)
	checkResult(b.Run("[CQ:at,qq=3334] 自杀", fromQQ+2, "mm"), "人物删除成功", t)
	checkResult(b.Run("[CQ:at,qq=3334] 自杀", fromQQ+3, "mm"), "人物删除成功", t)
	checkResult(b.Run("[CQ:at,qq=3334] 自杀", fromQQ+4, "mm"), "人物删除成功", t)

	checkResult(b.Run("[CQ:at,qq=3334] 创建", fromQQ, "mm"), "指令格式不正确！别想欺骗机器人！[[CQ:at,qq=3334] 创建]", t)

	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa;bbb;ccc;ddd", fromQQ, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa;bbb;ccc;200", fromQQ, "mm"))
	p := b.getPersonFromDb(fromQQ)
	p.SecretLevel = 5
	b.setPersonToDb(fromQQ, p)
	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa;bbb;ccc;200", fromQQ, "mm"))

	b.setMoney(fromQQ, 3000)
	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa;bbb;ccc;200", fromQQ, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa;bbb;幸运光环;200", fromQQ, "mm"))
	b.buyMace(fromQQ)
	qlog.Println(b.getItems(fromQQ))
	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa;bbb;幸运光环;100", fromQQ, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa;bbb;幸运光环;100", fromQQ, "mm"))

	qlog.Println(b.deleteChurch(fromQQ, "解散;aaa"))
	b.setMoney(fromQQ+1, 3000)
	b.setMoney(fromQQ+2, 3000)
	b.setMoney(fromQQ+3, 3000)
	b.buyMace(fromQQ + 1)
	b.buyMace(fromQQ + 2)
	b.buyMace(fromQQ + 3)
	b.setPersonToDb(fromQQ+1, p)
	b.setPersonToDb(fromQQ+2, p)
	b.setPersonToDb(fromQQ+3, p)

	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa1;bbb;幸运光环;100", fromQQ+1, "mm1"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa2;bbb;幸运光环;100", fromQQ+2, "mm2"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 创建;aaa3;bbb;幸运光环;100", fromQQ+3, "mm3"))

	qlog.Println(b.Run("[CQ:at,qq=3334] 寻访", fromQQ, "mm"))
	b.setMoney(fromQQ+1, 13000)
	b.setMoney(fromQQ+2, 33000)
	b.setMoney(fromQQ+3, 53000)
	p.SecretLevel = 9
	b.setPersonToDb(fromQQ+3, p)

	qlog.Println(b.Run("[CQ:at,qq=3334] 寻访", fromQQ, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 寻访", fromQQ, "mm"))

	qlog.Println(b.Run("[CQ:at,qq=3334] 加入;aaa1", fromQQ, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 加入;aaa2", fromQQ, "mm"))
	b.Update(fromQQ+4, "mouse")
	qlog.Println(b.Run("[CQ:at,qq=3334] 加入;aaa3", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ+4, "mm"))

	b.setMoney(fromQQ+4, 1000)
	qlog.Println(b.Run("[CQ:at,qq=3334] 加入;aaa3", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ+4, "mm"))

	qlog.Println(b.Run("[CQ:at,qq=3334] 寻访", fromQQ, "mm"))
	qlog.Println("money", b.getMoney(fromQQ+4))

	qlog.Println(b.Run("[CQ:at,qq=3334] 祈祷", fromQQ, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 祈祷", fromQQ+1, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 祈祷", fromQQ+2, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 祈祷", fromQQ+4, "mm"))
	b.buyMagicItem(fromQQ + 4)
	b.buyMagicItem(fromQQ + 4)

	qlog.Println(b.Run("[CQ:at,qq=3334] 祈祷", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 祈祷", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ+4, "mm"))

	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 探险", fromQQ+4, "mm"))

	qlog.Println(b.Run("[CQ:at,qq=3334] 退出", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 退出", fromQQ+4, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 退出", fromQQ+1, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 退出", fromQQ+3, "mm"))

	qlog.Println(b.Run("[CQ:at,qq=3334] 寻访", fromQQ, "mm"))
	qlog.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ+4, "mm"))

	qlog.Println(b.Run("[CQ:at,qq=3334] 排行", fromQQ+4, "mm"))

	qlog.Println(b.deleteChurch(fromQQ+1, "解散;aaa1"))
	qlog.Println(b.deleteChurch(fromQQ+2, "解散;aaa2"))
	qlog.Println(b.deleteChurch(fromQQ+3, "解散;aaa3"))

}
