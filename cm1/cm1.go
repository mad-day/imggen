/*
Copyright (c) 2018 Simon Schmidt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/


package cm1

import "image"
import "image/color"
import "github.com/mad-day/imggen/conv"

func in(r image.Rectangle, p image.Point) bool {
	return (r.Min.X <= p.X) && (r.Min.Y <= p.Y) && (p.X < r.Max.X) && (p.Y < r.Max.Y)
}
func to(I color.Color) (c [3]uint32) {
	c[0],c[1],c[2],_ = I.RGBA()
	return
}

var apts = [2]image.Point{
	{0,-1},{-1,0},
}
const div float64 = 0xFFFF

type ColorSet struct{
	point [6]float64
	val color.Color
}
func(cs *ColorSet) Decode(img image.Image,r image.Rectangle, p image.Point)  {
	var col [2][3]uint32
	for i := range col {
		x := p.Add(apts[i])
		if in(r,p) { col[i] = to(img.At(x.X,x.Y)) }
	}
	cs.point = [6]float64{
		/* ************************ */
		float64(col[0][0])/div,
		float64(col[0][1])/div,
		float64(col[0][2])/div,
		/* ************************ */
		float64(col[1][0])/div,
		float64(col[1][1])/div,
		float64(col[1][2])/div,
		/* ************************ */
	}
}
func (cs *ColorSet) Rect(ctx interface{}) (min []float64, max []float64) {
	min = cs.point[:]
	max = min
	return
}
func (cs *ColorSet) Value() color.Color {
	return cs.val
}
func (cs *ColorSet) SetValue(cval color.Color) {
	cs.val = cval
}

func MakeItem() conv.ImageItem {
	return new(ColorSet)
}

