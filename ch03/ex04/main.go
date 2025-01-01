// Surface computes an SVG rendering of a 3-D surface function.
//
// Exercise 3.4: Following the approach of the Lissajous example in Section 1.7,
// construct a web server that computes surfaces and writes SVG data to the client
package main

import (
	"bytes"
	"fmt"
	// "image"
	// "image/color"
	// "image/gif"
	"io"
	"log"
	"math"
	// "math/rand"
	"net/http"
	"net/url"
	"strconv"
    "strings"
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
	white         = "white"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)
	width, height := width, height
	color := "white"

	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(u)
	for k, v := range u {
		keyValue, exists := u[k]
		if !exists || len(keyValue) == 0 {
			fmt.Fprintf(w, "No 'cycles' parameter provided\n")
		}
		if k == "w" {
			width, _ = strconv.Atoi(v[0])
		} else if k == "h" {
			height, _ = strconv.Atoi(v[0])
		} else if k == "color" {
			color = strings.Trim(v[0], "\"")
		}
	}
	var b bytes.Buffer
	surface(&b, width, height, color)

	_, err = w.Write(b.Bytes())
	if err != nil {
		http.Error(w, "Failed to send GIF", http.StatusInternalServerError)
		fmt.Println("Error writing GIF to response:", err)
	}
	// http.ServeContent(w, r, "example.gif", file.ModTime(), file)
}

func surface(out io.Writer, width int, height int, color string) {
	fmt.Println("color = ", color)
	outStr := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j+1, width, height)
			if math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			if j == cells-1 && i == cells-1 {
				fmt.Println(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style=\"fill:%s\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color))
			}
			outStr += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style=\"fill:%s\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	outStr += fmt.Sprintf("</svg>")
	out.Write([]byte(outStr))
}

func corner(i, j, width, height int) (float64, float64) {
	// var color string
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	height := math.Sin(r) / r
	// fmt.Println(height)
	return height
}

//!-
