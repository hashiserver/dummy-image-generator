package lib

import (
	"image"
	"image/color"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/hashiserver/dummy-image-generator/model"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/math/fixed"
)

// 画像を生成
func GenerateImage(option model.Option) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, option.Width, option.Height))

	// gray
	fillRect(img, color.RGBA{128, 128, 128, 255})

	// 画像中央にwidth と heightを記載
	drawString(img, option.Width, option.Height)

	return img, nil

}

// 画像を単色に染める
func fillRect(img *image.RGBA, col color.Color) {
	rect := img.Rect
	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, col)
		}
	}
}

// 画像に文字を追加
func drawString(img *image.RGBA, width int, height int) error {
	fontstr := strconv.Itoa(width) + " x " + strconv.Itoa(height)
	fontsize := 48.0

	// フォントの読み込み
	ft, err := truetype.Parse(gobold.TTF)
	if err != nil {
		return err
	}

	opt := truetype.Options{Size: fontsize}
	face := truetype.NewFace(ft, &opt)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: face,
		Dot:  fixed.Point26_6{},
	}
	// 中央に配置
	d.Dot.X = (fixed.I(width) - d.MeasureString(fontstr)) / 2
	d.Dot.Y = fixed.I((height / 2) + int(fontsize/2))

	d.DrawString(fontstr)
	return nil
}
