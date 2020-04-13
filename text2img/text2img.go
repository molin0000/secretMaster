package text2img

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/molin0000/secretMaster/qlog"
	"golang.org/x/image/font"
)

func DrawTextImg(msg string) string {
	//计算图片高度
	strs := strings.Split(msg, "\n")
	i := 0
	for _, v := range strs {
		info := foldLine(v)
		infos := strings.Split(info, "\n")
		for range infos {
			i++
		}
	}

	height := i*(18+5) + 40
	width := 400
	img := imaging.New(width, height, color.NRGBA{0, 0, 0, 255})
	writeOnImage(img, msg)
	fileName := fmt.Sprintf("%d.png", time.Now().UnixNano())
	savePath := path.Join("data", "image", fileName)
	err := imaging.Save(img, savePath)
	if err != nil {
		return "保存文件失败：" + savePath + err.Error()
	}

	autodel := func() {
		d := time.Duration(time.Second * 200)
		t := time.NewTimer(d)
		delFile := savePath
		<-t.C
		err := os.Remove(delFile)
		if err != nil {
			qlog.Println(err.Error())
		} else {
			qlog.Println("成功删除：", delFile)
		}
	}

	go autodel()

	return fileName
}

func foldLine(msg string) string {
	DD := []rune(msg) //需要分割的字符串内容，将它转为字符，然后取长度。
	lineSize := 20
	info := ""

	for i := 0; i < len(DD); i++ {
		if i != 0 && i%lineSize == 0 {
			info += "\n"
		}
		info += string(DD[i])
	}
	return info
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
		qlog.Println("get font family error")
	}
	// 设置用于绘制文本的字体
	c.SetFont(fontFam)

	// 获取字体的尺寸大小
	fixed := c.PointToFixed(FontSize)

	strs := strings.Split(msg, "\n")
	i := 0
	for _, v := range strs {
		info := foldLine(v)
		infos := strings.Split(info, "\n")
		for _, sv := range infos {
			pt := freetype.Pt(20, 40+(fixed.Ceil()+5)*i)
			_, err = c.DrawString(sv, pt)
			if err != nil {
				qlog.Printf("draw error: %v \n", err)
				return
			}
			i++
		}
	}
}

// 获取字符集，仅调用一次
func getFontFamily() (*truetype.Font, error) {
	// 这里需要读取中文字体，否则中文文字会变成方格
	fontBytes, err := ioutil.ReadFile("C:/windows/Fonts/simsun.ttc")
	if err != nil {
		qlog.Println("read file error:", err)
		return &truetype.Font{}, err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		qlog.Println("parse font error:", err)
		return &truetype.Font{}, err
	}

	return f, err
}
