package main

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const S = 1024
	//im,_:=gg.LoadImage("")
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("C:\\Windows\\Fonts\\Arial.ttf", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	dc.SavePNG("out.png")
}

func BlankImage(width,height int, c color.Color){
	upLeft:= image.Point{}
	lowRight:= image.Point{X: width, Y: height}
	img:=image.NewRGBA(image.Rectangle{
		Min: upLeft,
		Max: lowRight,
	})
	for y := 0; y < lowRight.Y; y++ {
		for x := 0; x < lowRight.X; x++ {
			img.Set(x,y,c)
		}
	}
	f,_:=os.Create("gen.png")
	err := png.Encode(f, img)
	if err != nil {
		return
	}
}