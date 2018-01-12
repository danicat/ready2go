package main

import "fmt"

func main() {
	// START OMIT
	m := map[string]string{}

	m["hello"] = "world"
	fmt.Println(m["hello"])
	
	m2 := map[int]string {
		1: "um",
		2: "dois",
		3: "três",
	}

	if v, ok := m2[2]; ok {
		fmt.Println("achou: ", v)
	}

	if v, ok := m2[4]; ok {
		fmt.Println("achou: ", v)
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}
	// END OMIT
}
