package main

import "fmt"

// START OMIT
func main() {

	// 1. A palavra chave `make` é usada para criar canais
	canal := make(chan string)
	canal2 := make(chan string)

	// 2. O operador <- é usado para gravar no canal
	go func() {
		canal <- "esta mensagem é do canal"
		canal2 <- "esta mensagem é do canal2"
	}()

	// 3. O operador <-canal é usado para ler do canal
	msg, ok := <-canal
	fmt.Println(msg)
	msg = <-canal2
	fmt.Println(msg)
}

// END OMIT
