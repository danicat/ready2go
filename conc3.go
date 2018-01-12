package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func fazAlgoImportante(s string) {
	time.Sleep(1 * time.Second)
	fmt.Println(s)
}

// Programando em paralelo utilizando um waitgroups para sincronizar
func fazTudo(tarefas []string) {
	var wg sync.WaitGroup
	for _, tarefa := range tarefas {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			fmt.Println(t)
		}(tarefa)
		//time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()
}

func main() {
	tarefas := []string{"dormir", "comer", "correr", "trabalhar", "comer", "dormir"}
	fazTudo(tarefas)
}
// END OMIT


