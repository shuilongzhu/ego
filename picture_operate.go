package ego

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"strings"
)

// AddLabel @description: 操作图片添加水印标签
// @parameter byteList(原图片字节数组)
// @parameter label(要添加的文字)
// @parameter x(图片x轴值)
// @parameter y(图片y轴值)
// @parameter fontColor(文字颜色)
// @parameter size(文字大小)
// @parameter fontPath(指定字体的位置)
// @return err
// @return result(字节数组)
func AddLabel(byteList []byte, label string, x, y int, fontColor color.Color, size float64, fontPath string) (err error, result []byte) {
	img, _, err := image.Decode(bytes.NewReader(byteList))
	if nil != err {
		log.Errorln("png.Decode err:", err)
		return
	}

	bound := img.Bounds()
	// 创建一个新的图片
	rgba := image.NewRGBA(image.Rect(0, 0, bound.Dx(), bound.Dy()))
	// 读取字体
	fontBytes, err := ioutil.ReadFile(fontPath)
	if nil != err {
		log.Errorln("ioutil.ReadFile err:", err)
		return
	}

	myFont, err := freetype.ParseFont(fontBytes)
	if nil != err {
		log.Errorln("freetype.ParseFont err:", err)
		return
	}

	draw.Draw(rgba, rgba.Bounds(), img, bound.Min, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(myFont)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	uni := image.NewUniform(fontColor)
	c.SetSrc(uni)
	c.SetHinting(font.HintingNone)

	// 在指定的位置显示
	pt := freetype.Pt(x, y+int(c.PointToFixed(size)>>6))
	if _, err = c.DrawString(label, pt); err != nil {
		log.Errorln("c.DrawString err:", err)
		return
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, rgba)
	if err != nil {
		log.Errorln("png.Encode err:", err)
		return
	}
	result = buf.Bytes()

	return
}

// SaveThumbnailImg @description:缩略图存储
// @parameter width
// @parameter height
// @parameter filePath(原来图片地址)
// @return int
func SaveThumbnailImg(width, height int, filePath string) (err error) {
	//读取本地原来图片
	imgData, err := ioutil.ReadFile(filePath)
	if nil != err {
		return
	}
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if nil != err {
		return
	}
	//width传了值，height传0则表示宽按传值宽度拉伸，高则等比例拉伸
	//最后缩略图尺寸为100*133 height==0
	image = imaging.Resize(image, width, height, imaging.Lanczos)
	var oldStr = "."
	err = imaging.Save(image, strings.Replace(filePath, oldStr, "-Thumbnails"+oldStr, -1))
	return
}
