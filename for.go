package main

import "fmt"

func main() {
	// For tradicional
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	j := 0
	// Apenas teste: "while"
	for j < 1 {
		fmt.Println(j)
		if j == 0 { // if tradicional
			j = 1
		}
	} 

	// Sem condição: loop infinito
	for {
		if n, _ := fmt.Println("Loop infinito"); n > 0 { // if com inicialização
			break
		}
		// continue
	}
}