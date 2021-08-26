package utility

import (
	"fmt"
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

//"golang.org/x/image/font/inconsolata"
//// To use regular Inconsolata font family:
//Face: inconsolata.Regular8x16,
//
//// To use bold Inconsolata font family:
//Face: inconsolata.Bold8x16,

func AddLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

type Request struct {
	BgImgPath string
	FontPath  string
	FontSize  float64
	Text      string
	PathSave  string
}

func TextOnImg(request Request) (Error error) {
	bgImage, err := gg.LoadImage(request.BgImgPath)
	if err != nil {
		Error = err
		return
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	//if err := dc.LoadFontFace(request.FontPath, request.FontSize); err != nil {
	//	fmt.Println("2 : ",err)
	//}

	x := float64(imgWidth / 2)
	y := float64((imgHeight / 2) - 80)
	maxWidth := float64(imgWidth) - 60.0
	color := color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}
	dc.SetColor(color)
	dc.SetFillRuleWinding()
	dc.DrawStringWrapped(request.Text, x, y, 0.5, -38.5, maxWidth, 5, gg.AlignCenter)
	dc.Image()

	out, err := os.Create(request.PathSave)
	if err != nil {
		fmt.Println(err)
	}
	var opt jpeg.Options
	opt.Quality = 80
	jpeg.Encode(out, dc.Image(), &opt)
	return
}
