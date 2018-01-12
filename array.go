package main

import "fmt"

func main() {
	// START OMIT
	var inteirosVazio [3]int 	// 1. Declaração de um array vazio
	inteirosVazio2 := [3]int{}

	inteiros := [3]int{1, 2, 3} // 2. Declaração de um array com valores
	frutas := [...]string{"banana", "laranja", "maçã", "abacate"}

	fmt.Println(inteirosVazio, inteirosVazio2, inteiros)

	// Os operadores len e cap mostram o tamanho atual e a capacidade máxima do array...
	fmt.Printf("len: %v\tcap: %v\ttype: %T\n",
		len(frutas),
		cap(frutas),
		frutas)

	for indice, fruta := range frutas {
		fmt.Println(indice, fruta)
	}

	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
	// END OMIT
}
