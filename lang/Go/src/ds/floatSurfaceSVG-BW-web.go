// 练习 3.4： 构造一个web服务器，用于计算函数曲面然后返回SVG数据给客户端。

package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var height, width = 300, 600 // canvas size in pixels

var xyscale = width / 2 / xyrange  // pixels per x or y unit
var zscale = float64(height) * 0.4 // pixels per z unit

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	addr := ":8000"
	fmt.Printf("Visit\n  http://localhost%s/\n  http://localhost%[1]s/?height=300&width=600\n", addr)

	//http服务
	http.HandleFunc("/", handle)
	http.ListenAndServe(addr, nil)
}
func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	if err := r.ParseForm(); err != nil {
		return
	}
	//var height int
	//var width  int
	for k, v := range r.Form {
		if k == "height" {
			h, _ := strconv.Atoi(v[0])
			if h > 0 {
				height = h
			}
		}
		if k == "width" {
			w, _ := strconv.Atoi(v[0])
			if w > 0 {
				width = w
			}
		}
	}
	xyscale = width / 2 / xyrange
	zscale = float64(height) * 0.4
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: #ff0000; fill: #0000ff; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")

}
func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width/2) + (x-y)*cos30*float64(xyscale)
	sy := float64(height/2) + (x+y)*sin30*float64(xyscale) - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
