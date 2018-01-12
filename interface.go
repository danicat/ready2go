package main

import (
	"fmt"
)

// START OMIT
type rect struct {
	h, w int
}

type obj interface {
	area() int
}

func (r rect) area() int {
	return r.h * r.w
}

func printArea(o obj) {
	fmt.Printf("Area do %T: %v", o, o.area())
}

func main() {
	r := rect{4,3}
	printArea(r)
}
// END OMIT