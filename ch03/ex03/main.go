// Surface computes an SVG rendering of a 3-D surface function.
//
// Exercise 3.3: Color each polygon based on its height,
// so that the peaks are colored red (#ff0000) and the valleys blue (#0000ff).
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
	red           = "#FF0000"
	blue          = "#0000FF"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, color := corner(i+1, j)
			bx, by, color := corner(i, j)
			cx, cy, color := corner(i, j+1)
			dx, dy, color := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style=\"fill:%s\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, string) {
	// var color string
	// Find point (x,y) at corner of cell (i,j).
	paint := "white"
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if z > .06 {
		paint = red
	} else if z < -0.06 {
		paint = blue
	}
	fmt.Println(z)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, paint
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	height := math.Sin(r) / r
	// fmt.Println(height)
	return height
}

//!-
