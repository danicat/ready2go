package main

import "fmt"

func main() {
	var i int
	var txt string
	i = 1
	txt = "um texto"

	fmt.Println(i, txt)

	txt2 := "outro texto"
	d := 3.14
	var cmp = 4+5i
	fmt.Printf("%v, %v, %v\n", txt2, d, cmp)
	fmt.Printf("%T, %T, %T\n", txt2, d, cmp)
}