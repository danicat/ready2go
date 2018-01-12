package main

import "fmt"

func main() {
	// START OMIT
	ch2 := make(chan int, 2) // Canal `buffered`

	go func() {
		defer close(ch2) // esquecer de fechar vai bloquear o range loop
		ch2 <- 2
		ch2 <- 3
		ch2 <- 4 // Terceiro write - vai esperar liberar um "slot"
	}()

	for val := range ch2 {
		fmt.Printf("ch2: %v\n", val)
	}
	// END OMIT
}
