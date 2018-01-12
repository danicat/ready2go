package main

import "fmt"

// Função que retorna um canal
func geraPares(quantidade int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= quantidade; i++ {
			ch <- i * 2
		}
	}()
	return ch
}

// START OMIT
// Função que recebe um canal e devolve outro
func triplicaValores(entrada chan int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for val := range entrada {
			ch <- val * 3
		}
	}()
	return ch
}

func main() {
	fmt.Println("Exemplo de pipeline:")
	meusDados3 := triplicaValores(geraPares(10))
	for val := range meusDados3 {
		fmt.Printf("Valor: %v\n", val)
	}
}
// END OMIT
