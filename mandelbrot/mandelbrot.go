package main

import (
	"flag"
	"fmt"
)

func mandelbrotControl(debug bool) {
	fmt.Println("mandelbrotControl")
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
