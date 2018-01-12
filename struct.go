package main

import (
	"fmt"
)
// START OMIT
type str1 struct {
	p1 string
	p2 int
}

func main() {
	v1 := str1{"meaning", 42}

	v2 := struct{
		a string
		b string
		c string
	}{
		"abc",
		"def",
		"ghi",
	}

	fmt.Println(v1, v2)
}
// END OMIT