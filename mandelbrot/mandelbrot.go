package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

// Mandelbrot pseudocode
//
/*
For each pixel (Px, Py) on the screen, do:
{
  x0 = scaled x coordinate of pixel (scaled to lie in the Mandelbrot X scale (-2.5, 1))
  y0 = scaled y coordinate of pixel (scaled to lie in the Mandelbrot Y scale (-1, 1))
  x = 0.0
  y = 0.0
  iteration = 0
  max_iteration = 1000
  while (x*x + y*y <= 2*2  AND  iteration < max_iteration) {
    xtemp = x*x - y*y + x0
    y = 2*x*y + y0
    x = xtemp
    iteration = iteration + 1
  }
  color = palette[iteration]
  plot(Px, Py, color)
}
*/

func mandelbrotControl(debug bool) {
	//	p := image.Point{2, 1}

	r := image.Rect(2, 1, 5, 5)
	fmt.Println(r.Dx(), r.Dy(), image.Pt(0, 0).In(r)) // prints 3 4 false

	fmt.Println("mandelbrotControl")

	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	m.Set(5, 5, color.RGBA{255, 0, 0, 255})

	// new code
	rectangle := "rectangle.png"

	rectImage := image.NewRGBA(image.Rect(0, 0, 200, 200))
	green := color.RGBA{0, 100, 0, 255}

	draw.Draw(rectImage, rectImage.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	file, err := os.Create(rectangle)
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, rectImage)
}

// Main routine
func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "turns print debugging on")
	execPartPtr := flag.String("part", "a", "Which part of day06 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		mandelbrotControl(debug)
	case "b":
		fmt.Println("Not here yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
