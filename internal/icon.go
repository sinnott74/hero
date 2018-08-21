package internal

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/DAddYE/vips"
)

// AddIcons add icons to the pic
func AddIcons(pic *image.RGBA, iconPaths []string, iconSize image.Point) error {

	width := pic.Bounds().Max.X
	height := pic.Bounds().Max.Y

	numIcons := len(iconPaths)
	horizonalPaddingNum := numIcons + 1
	paddingWidth := (width - iconSize.Y*numIcons) / horizonalPaddingNum

	for i, iconPath := range iconPaths {
		icon, err := readIcon(&iconPath)
		if err != nil {
			return err
		}
		resizedIcon, err := resizeIcon(&icon, iconSize)
		if err != nil {
			return err
		}

		diff := iconSize.Sub(resizedIcon.Bounds().Max).Div(2)
		offset := image.Pt(paddingWidth+iconSize.X*i+paddingWidth*i, height/2-iconSize.Y/2).Add(diff)
		addIconToPic(pic, resizedIcon, offset)
	}
	return nil
}

// CreatePNG creates the image in the filesystem
func CreatePNG(pic draw.Image, out *string) error {
	// Create output file
	outFile, err := os.Create(*out)
	if err != nil {
		return err
	}

	// Write the file
	err = png.Encode(outFile, pic)
	if err != nil {
		return err
	}
	return nil
}

// readIcon reads in the btyes of a given file
func readIcon(iconPath *string) ([]byte, error) {
	f, err := os.Open(*iconPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	inBuf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return inBuf, nil
}

// resizeIcon resizes the file
func resizeIcon(icon *[]byte, iconSize image.Point) (image.Image, error) {
	options := vips.Options{
		Width:        iconSize.X,
		Height:       iconSize.Y,
		Crop:         false,
		Extend:       vips.EXTEND_WHITE,
		Interpolator: vips.BILINEAR,
		Gravity:      vips.CENTRE,
		Quality:      100,
		Format:       vips.PNG,
	}

	buf, err := vips.Resize(*icon, options)
	if err != nil {
		return nil, err
	}

	resizedIcon, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	return resizedIcon, nil
}

// addIconToPic adds the icon to the picture
func addIconToPic(pic draw.Image, icon image.Image, offset image.Point) {
	draw.Draw(pic, icon.Bounds().Add(offset), icon, image.ZP, draw.Over)
}
