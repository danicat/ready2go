package main

import "fmt"

func main() {
	something := 10
	addrOfSomething := &something
	fmt.Printf("valor: %v, endereço: %v\n", *addrOfSomething, addrOfSomething)
	// addrOfSomething++ // Erro! Não é um inteiro!
}