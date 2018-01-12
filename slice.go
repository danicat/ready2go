package main

import "fmt"

func main() {
	// START OMIT
	slice := []string{
		"A",
		"B",
	}
	fmt.Println(slice)

	slice = append(slice, "C") // Operação de adição de elementos
	fmt.Println(slice)

	slice2 := []string{
		"d",
		"e",
	}

	slice = append(slice, slice2...)
	fmt.Println(slice)

	fmt.Println(slice[1:2])
	fmt.Println(slice[1:])
	fmt.Println(slice[:3])
	// END OMIT
}
