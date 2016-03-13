package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

const (
	width, height = 1024, 768
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Println("listening on http://localhost:8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", RenderSurface)
	mux.HandleFunc("/animated", RenderAnimatedSurface)
	http.ListenAndServe(":8080", mux)
}

func RenderSurface(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(rw, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: rgb(40, 40, 40); stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, 0)
			bx, by, bz := corner(i, j, 0)
			cx, cy, cz := corner(i, j+1, 0)
			dx, dy, dz := corner(i+1, j+1, 0)
			if valid(ax, ay, bx, by, cx, cy, dx, dy) {
				r, g, b := depthColor(az, bz, cz, dz)
				fmt.Fprintf(rw, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
					" style='fill: rgb(%d, %d, %d)' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
			}
		}
	}
	fmt.Fprintln(rw, "</svg>\n")
}

func RenderAnimatedSurface(rw http.ResponseWriter, r *http.Request) {
	const boundary = "NEWSVGDATA"
	rw.Header().Set("Content-Type", "multipart/x-mixed-replace;boundary="+boundary)
	rw.WriteHeader(http.StatusOK)

	for frame := 0; ; frame++ {
		fmt.Fprintf(rw, "\r\n--"+boundary+"\r\n")
		fmt.Fprintf(rw, "Content-type: text/html\r\n\r\n")

		fmt.Fprintf(rw, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: rgb(40, 40, 40); stroke-width: 0.7' "+
			"width='%d' height='%d'>\n", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, az := corner(i+1, j, frame)
				bx, by, bz := corner(i, j, frame)
				cx, cy, cz := corner(i, j+1, frame)
				dx, dy, dz := corner(i+1, j+1, frame)
				if valid(ax, ay, bx, by, cx, cy, dx, dy) {
					r, g, b := depthColor(az, bz, cz, dz)
					fmt.Fprintf(rw, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
						" style='fill: rgb(%d, %d, %d)' />\n",
						ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
				}
			}
		}
		fmt.Fprintf(rw, "</svg>")
		rw.(http.Flusher).Flush()
		time.Sleep(time.Millisecond)
	}
}

func valid(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func corner(i, j, frame int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := depth(x, y, float64(frame)/24.0*math.Pi)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func depth(x, y, offset float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r+offset) / r
}

func depth2(x, y float64) float64 {
	return math.Cos(x/xyrange*2) + math.Sin(y/xyrange*2)
}

func depthColor(points ...float64) (int, int, int) {
	var max float64
	for _, v := range points {
		max = math.Max(max, v)
	}
	max *= 2
	return int(255 * max), int(255 * (1 - math.Abs(0.5-max))), int(255 * (1 - max))
}
