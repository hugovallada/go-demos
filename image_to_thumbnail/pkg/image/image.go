package image

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
)

func GenerateThumbnail(path string) error {
	imageFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	img, err := png.Decode(imageFile)
	if err != nil {
		return err
	}
	thumbnail := imaging.Thumbnail(img, 80, 80, imaging.CatmullRom)
	thumbnailImg := imaging.New(80, 80, color.NRGBA{0, 0, 0, 0})
	thumbnailImg = imaging.Paste(thumbnailImg, thumbnail, image.Pt(0, 0))
	tmpFile, err := os.CreateTemp(".", "thumbnail*.png")
	if err != nil {
		return err
	}
	err = imaging.Save(thumbnailImg, tmpFile.Name())
	return err
}
