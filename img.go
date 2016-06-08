package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)


type Image struct{
    width int
    height int
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
    v := uint8(x^y)
    return color.RGBA{v, v, 254, 98}
}


func main() {
    m := Image{256, 64}
    pic.ShowImage(m)
}