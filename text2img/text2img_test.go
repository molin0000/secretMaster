package text2img

import (
	"testing"

	"github.com/molin0000/secretMaster/qlog"
	// "github.com/bregydoc/gtranslate"
)

func TestImg(t *testing.T) {
	str := `昵称：空想之喵
途径：魔女
序列：序列4：绝望
勋章：3🎖🎖🎖
经验：6085
金镑：892
幸运：1
灵力：91
一二三四五六七吧旧时一二三四五六七吧旧时
一二三四五六七吧旧时一二三四五六七吧旧时一二三四五六七吧旧时一二三四五六七吧旧时一二三四五六七吧旧时
修炼时间：603小时`
	qlog.Println(DrawTextImg(str))
}
