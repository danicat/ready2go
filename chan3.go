package main

import "fmt"

// START OMIT
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

func main() {
	fmt.Println("Exemplo de generator:")
	meusDados := geraPares(10)
	for val := range meusDados {
		fmt.Printf("Valor: %v\n", val)
	}
}
// END OMIT
