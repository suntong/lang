// https://www.zoulei.net/2018/06/06/the_go_programming_lang_usage_answer_3/
// by 邹雷

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	min, max := getMinMax()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, r, g, b := corner(i+1, j, min, max)
			bx, by, r, g, b := corner(i, j, min, max)
			cx, cy, r, g, b := corner(i, j+1, min, max)
			dx, dy, r, g, b := corner(i+1, j+1, min, max)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%02x%02x%02x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Println("</svg>")
}

func getColor(min, max, current float64) (int, int, int) {
	step := (max - min) / 255
	v := int((current - min) / step)
	r := v
	g := 0
	b := 255 - v
	return r, g, b
}

func getMinMax() (float64, float64) {
	var min, max float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)

			z := f(x, y)
			if z < min {
				min = z
			}
			if z > max {
				max = z
			}
		}
	}
	return min, max
}

func corner(i, j int, min, max float64) (float64, float64, int, int, int) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	// 将(x,y,z)等角投射到二维SVG绘图平面上,坐标是(sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	r, g, b := getColor(min, max, z)
	return sx, sy, r, g, b
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
