package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/golang/geo/r2"
)

func drawMapImg(mapName string) (*image.RGBA, image.Rectangle) {
	mapFilePath := fmt.Sprintf("assets/%s.jpeg", mapName)
	mapFile, merr := os.Open(mapFilePath)
	if merr != nil {
		panic(merr)
	}
	defer mapFile.Close()

	mapImg, _, er := image.Decode(mapFile)
	if er != nil {
		panic(er)
	}
	bounds := mapImg.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, mapImg, image.Point{}, draw.Over)
	return img, bounds
}

func drawPointsImage(points []r2.Point, bounds image.Rectangle, ptsColor color.RGBA) *image.RGBA {
	ptsImg := image.NewRGBA(bounds)

	for _, pt := range points {
		for inc := 0; inc < 20; inc++ {
			ptsImg.SetRGBA(int(pt.X), int(pt.Y)+inc, ptsColor)
			ptsImg.SetRGBA(int(pt.X), int(pt.Y)-inc, ptsColor)
		}
		for inc := 0; inc < 20; inc++ {
			ptsImg.SetRGBA(int(pt.X)+inc, int(pt.Y), ptsColor)
			ptsImg.SetRGBA(int(pt.X)-inc, int(pt.Y), ptsColor)
		}
	}
	return ptsImg
}

func PlotPointsOnMap(mapName string, points []r2.Point, plotColor color.RGBA) *image.RGBA {
	mapImg, bounds := drawMapImg(mapName)
	ptsImg := drawPointsImage(points, bounds, plotColor)
	draw.Draw(mapImg, bounds, ptsImg, image.Point{}, draw.Over)
	return mapImg
}

func ExportImageToFile(img *image.RGBA, fileName string) {
	destfile, derr := os.Create(fileName)
	if derr != nil {
		panic(derr)
	}
	jerr := jpeg.Encode(destfile, img, &jpeg.Options{Quality: 90})
	if jerr != nil {
		panic(jerr)
	}

}

func ExportImageToBase64(img *image.RGBA) string {
	b := new(bytes.Buffer)
	rr := jpeg.Encode(b, img, &jpeg.Options{Quality: 90})
	base64Encoding := "data:image/jpeg;base64,"
	if rr != nil {
		panic(rr)
	}
	by := b.Bytes()
	base64Encoding += base64.StdEncoding.EncodeToString(by)
	return base64Encoding
}
