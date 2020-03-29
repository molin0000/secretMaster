package secret

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	fromQQ := uint64(67939461000)
	b := NewSecretBot(3334, 333, "aaa", false, &debugInteract{})

	checkResult(b.Run("[CQ:at,qq=3334] 自杀", fromQQ, "mm"), "人物删除成功", t)
	b.Update(fromQQ, "fish")
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;探测", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;埋入", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;埋入;10", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;埋入;1000", fromQQ, "mm"))
	b.setMoney(fromQQ, 2000)
	fmt.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;埋入;1000", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;埋入;1000", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;探测", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;挖出;3000", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;挖出;2000", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 地窖;探测", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ, "mm"))
}

func TestWork(t *testing.T) {
	fromQQ := uint64(67939461000)
	b := NewSecretBot(3334, 333, "aaa", false, &debugInteract{})

	checkResult(b.Run("[CQ:at,qq=3334] 自杀", fromQQ, "mm"), "人物删除成功", t)
	b.Update(fromQQ, "fish")

	fmt.Println(b.Run("[CQ:at,qq=3334] 工作", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 工作;停止", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 工作;女装直播", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 工作;女装直播", fromQQ, "mm"))
	fmt.Println(b.Run("[CQ:at,qq=3334] 工作;停止", fromQQ, "mm"))
}

func TestFishing(t *testing.T) {
	fromQQ := uint64(67939461000)
	b := NewSecretBot(3334, 333, "aaa", false, &debugInteract{})
	checkResult(b.Run("[CQ:at,qq=3334] 自杀", fromQQ, "mm"), "人物删除成功", t)
	b.Update(fromQQ, "fish")

	b.setMagic(fromQQ, 1200)
	for i := 0; i < 300; i++ {
		fmt.Println(b.Run("[CQ:at,qq=3334] 钓鱼", fromQQ, "mm"))
	}

	fmt.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ, "mm"))

}

func TestLottery(t *testing.T) {
	fromQQ := uint64(67939461000)
	b := NewSecretBot(3334, 333, "aaa", false, &debugInteract{})
	checkResult(b.Run("[CQ:at,qq=3334] 自杀", fromQQ, "mm"), "人物删除成功", t)
	b.Update(fromQQ, "fish")

	b.setMoney(fromQQ, 1200)
	for i := 0; i < 1000; i++ {
		fmt.Println(b.Run("[CQ:at,qq=3334] 许愿", fromQQ, "mm"))
	}

	fmt.Println(b.Run("[CQ:at,qq=3334] 属性", fromQQ, "mm"))
}
