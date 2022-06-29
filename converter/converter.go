package converter

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const HHMMSS24 = "10:34:52"

var (
	backgroundWidth  = 100
	backgroundHeight = 50
	utf8FontFile     = "fonts/Raleway-Black.ttf"
	utf8FontSize     = float64(20.0)
	dpi              = float64(72)
	utf8Font         = new(truetype.Font)
	ctx              = new(freetype.Context)
	red              = color.RGBA{255, 0, 0, 255}
	blue             = color.RGBA{0, 0, 255, 255}
	white            = color.RGBA{255, 255, 255, 255}
	black            = color.RGBA{0, 0, 0, 255}
	background       *image.RGBA
)

func BuildTimeInPng(k int) {

	tmptime := time.Now().Format(time.Stamp)
	time := strings.Split(tmptime, " ")[2]

	fontBytes, err := ioutil.ReadFile(utf8FontFile)
 	if err != nil {
 		log.Println(err)
 		return 
 	}

 	utf8Font, err = freetype.ParseFont(fontBytes)
 	if err != nil {
 		log.Println(err)
		return 

 	}

	fontForeGroundColor, fontBackGroundColor := image.NewUniform(black), image.NewUniform(white)

	background = image.NewRGBA(image.Rect(0, 0, backgroundWidth*k, backgroundHeight*k))

	draw.Draw(background, background.Bounds(), fontBackGroundColor, image.Point{}, draw.Src)

	ctx = freetype.NewContext()

	ctx.SetDPI(dpi)
	ctx.SetFont(utf8Font)
	ctx.SetFontSize(utf8FontSize * float64(k))
	ctx.SetClip(background.Bounds())
	ctx.SetDst(background)
	ctx.SetSrc(fontForeGroundColor)

	pt := freetype.Pt(0, int(ctx.PointToFixed(utf8FontSize)>>4))

	_, err = ctx.DrawString(time, pt)
	if err != nil {
		log.Fatal(err)
		return 
	}

	f, err := os.Create("tmp/img.png")
	if err != nil {
		log.Fatal(err)
		return 
	}

	defer func() { _ = f.Close() }()

	buff := bufio.NewWriter(f)

	if err = png.Encode(buff, background); err != nil {
		log.Fatal(err)
		return 
	}

	err = buff.Flush()
	if err != nil {
		log.Fatal(err)
		return 
	}

}
