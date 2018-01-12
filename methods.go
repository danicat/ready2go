package main

import "fmt"

type Rectangle struct {
    length, width int
}

func (r Rectangle) Area() int {
    return r.length * r.width
}

func main() {
    r1 := Rectangle{4, 3}
    fmt.Println("Rectangle is: ", r1)
    fmt.Println("Rectangle area is: ", r1.Area())
}