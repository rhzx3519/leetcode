package main

import "github.com/fogleman/gg"

const BASE_PATH = "src/leetcode/二叉树/visualize_tree/examples/"

func main() {
	const NX = 4
	const NY = 3
	im, err := gg.LoadPNG(BASE_PATH + "gopher.png")
	if err != nil {
		panic(err)
	}
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	dc := gg.NewContext(w*NX, h*NY)
	for y := 0; y < NY; y++ {
		for x := 0; x < NX; x++ {
			dc.DrawImage(im, x*w, y*h)
		}
	}
	dc.SavePNG(BASE_PATH + "tiling.png")
}
