package text2img

import (
	"fmt"
	"strings"
	"testing"
	"time"

	// "github.com/bregydoc/gtranslate"
	"image"
	"image/color"
	"io/ioutil"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func TestImg(t *testing.T) {
	fmt.Println("hello")
	str := `昵称：空想之喵
途径：魔女
序列：序列4：绝望
勋章：3
经验：6085
金镑：892
幸运：1
灵性：91
修炼时间：603小时`
	fmt.Println(DrawTextImg(str))
}

func DrawTextImg(msg string) string {
	strs := strings.Split(msg, "\n")
	height := len(strs)*(18+5) + 40
	width := 400
	img := imaging.New(width, height, color.NRGBA{0, 0, 0, 1})
	writeOnImage(img, msg)
	fileName := fmt.Sprintf("%d.png", time.Now().UnixNano())
	err := imaging.Save(img, fileName)
	if err != nil {
		return ""
	}
	return fileName
}

func writeOnImage(target *image.NRGBA, msg string) {
	const FontSize = 5

	c := freetype.NewContext()

	// 设置屏幕每英寸的分辨率
	c.SetDPI(256)
	// 背景
	c.SetClip(target.Bounds())
	// 设置目标图像
	c.SetDst(target)
	c.SetHinting(font.HintingFull)

	// 设置文字颜色、字体、字大小
	c.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 255, B: 255, A: 255}))
	// 以磅为单位设置字体大小
	c.SetFontSize(FontSize)
	fontFam, err := getFontFamily()
	if err != nil {
		fmt.Println("get font family error")
	}
	// 设置用于绘制文本的字体
	c.SetFont(fontFam)

	// 获取字体的尺寸大小
	fixed := c.PointToFixed(FontSize)
	fmt.Println(fixed.Ceil())

	strs := strings.Split(msg, "\n")

	for i, v := range strs {
		pt := freetype.Pt(20, 40+(fixed.Ceil()+5)*i)
		_, err = c.DrawString(v, pt)
		if err != nil {
			fmt.Printf("draw error: %v \n", err)
			return
		}
	}
}

// 获取字符集，仅调用一次
func getFontFamily() (*truetype.Font, error) {
	// 这里需要读取中文字体，否则中文文字会变成方格
	fontBytes, err := ioutil.ReadFile("/Users/molin/Downloads/simsun.ttc")
	if err != nil {
		fmt.Println("read file error:", err)
		return &truetype.Font{}, err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("parse font error:", err)
		return &truetype.Font{}, err
	}

	return f, err
}
