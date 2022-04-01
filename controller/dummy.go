package controller

import (
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/hashiserver/dummy-image-generator/lib"
	"github.com/hashiserver/dummy-image-generator/model"
)

// CreateImageFile ファイルにimageを書き込み
func CreateImageFile(filePath string, option model.Option) error {

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := lib.GenerateImage(option)
	if err != nil {
		return err
	}

	switch option.Extension {
	case "jpg", "jpeg":
		err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
		if err != nil {
			return err
		}
	case "gif":
		err := gif.Encode(f, img, nil)
		if err != nil {
			return err
		}
	default:
		err := png.Encode(f, img)
		if err != nil {
			return err
		}
	}

	return nil
}
