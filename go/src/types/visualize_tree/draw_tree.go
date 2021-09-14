package visualize_tree

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	types2 "github.com/rhzx3519/leetcode/go/src/types"
	"golang.org/x/image/font/gofont/goregular"
	"strconv"
)

var	CIRCLE_COLOR = [3]int{0, 0, 0}


type Drawer struct {
	dc *gg.Context
}

func NewDrawer(w, h int) *Drawer {
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 20,
	})
	dc.SetFontFace(face)

	return &Drawer{dc}
}

func (d *Drawer) Save(path string) {
	d.dc.SavePNG(path)
}

func (d *Drawer) DrawTreeNode(x, y float64, node *types2.TreeNode)  {
	if node == nil {
		return
	}
	val := node.Value.(int)
	d.DrawCirle(x, y, 30)
	d.DrawText(x, y, strconv.Itoa(val))
}

func (d *Drawer) DrawCirle(x, y, radius float64) {
	d.dc.DrawCircle(x, y, radius)
	d.dc.SetRGB255(CIRCLE_COLOR[0], CIRCLE_COLOR[1], CIRCLE_COLOR[2])
	d.dc.Stroke()
}

func (d *Drawer) DrawText(x, y float64, text string)  {
	d.dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
}

func (d *Drawer) DrawLine(x1, y1, x2, y2 float64) {
	d.dc.SetRGB(0, 0, 0)
	d.dc.SetLineWidth(2)
	d.dc.DrawLine(x1, y1, x2, y2)
	d.dc.Stroke()
}


