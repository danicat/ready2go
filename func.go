package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func multiplos(a int) (int, int, int) {
	return a * 2, a * 3, a * 4
}

func main() {
	a, b := 1, 2
	fmt.Printf("%v + %v = %v\n", a, b, soma(a, b))
	fmt.Println(multiplos(2))

	c, d, _ := multiplos(2)
	fmt.Println(c, d)
}
