package main

import "fmt"

func main() {
	// START OMIT
	ch1 := make(chan int) // Canal `unbuffered`

	go func() {
		defer close(ch1) // Se você comentar esta linha...
		ch1 <- 1
	}()

	val, ok := <-ch1
	fmt.Printf("Lendo de ch1 (aberto): %v, %v\n", val, ok)

	val, ok = <-ch1 // ... o código vai travar aqui
	fmt.Printf("Lendo de ch1 (fechado): %v, %v\n", val, ok)
	// END OMIT
}
