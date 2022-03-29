package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/math/fixed"
)

func CreateImageFile(width int, height int, extension string) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// gray
	fillRect(img, color.RGBA{128, 128, 128, 255})

	f, _ := os.Create("./dummy_" + strconv.Itoa(width) + "_" + strconv.Itoa(height) + "." + extension)
	defer f.Close()

	// 画像中央にwidth と heightを記載
	drawString(img, width, height)

	switch extension {
	case "jpg", "jpeg":
		err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	case "gif":
		err := gif.Encode(f, img, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	default:
		err := png.Encode(f, img)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
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

func drawString(img *image.RGBA, width int, height int) {
	fontstr := strconv.Itoa(width) + " x " + strconv.Itoa(height)
	fontsize := 48.0

	// フォントの読み込み
	ft, err := truetype.Parse(gobold.TTF)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	opt := truetype.Options{Size: fontsize}
	face := truetype.NewFace(ft, &opt)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: face,
		Dot:  fixed.Point26_6{},
	}
	d.Dot.X = (fixed.I(width) - d.MeasureString(fontstr)) / 2
	d.Dot.Y = fixed.I((height / 2) + int(fontsize/2))

	d.DrawString(fontstr)
}

func IsExtension(extension string) bool {
	switch extension {
	case "jpg", "jpeg", "gif", "png":
		return true
	default:
		return false
	}
}
