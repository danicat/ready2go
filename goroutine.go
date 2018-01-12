package main

import (
	"fmt"
	"time"
)

// START OMIT
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("chamada direta")

	go f("chamada com goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("goroutine com função anônima")

	time.Sleep(time.Second)

	fmt.Println("fim!")
}
// END OMIT