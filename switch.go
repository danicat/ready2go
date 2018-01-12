package main

import "fmt"

func main() {
	// START OMIT
	a, b := 0, 1
	switch a {
	case 0:
		fmt.Println("zero") // break implicito
	case b + 1: 
		fmt.Println("um")
	}

	switch c := "dois"; c {
	case "dois": 
		fmt.Println("acertou")
	default: 
		fmt.Println("erroooou!")
	}

	var x interface{} = "texto"
	switch x.(type) {
	case string: 
		fmt.Println("é uma string")
	case int: 
		fmt.Println("é um inteiro")
	}
	// END OMIT
}