package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello world")
	}()

	f := highOrder()
	f()
}

func highOrder() func() {
	return func() {
		fmt.Println("high order function")
	}
}