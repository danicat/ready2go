package main

import (
	"fmt"
	"time"
)

// START OMIT
// Programando sequencialmente
func fazAlgoImportante(s string) {
	time.Sleep(1 * time.Second)
	fmt.Println(s)
}

// Programando em paralelo utilizando um canal para sincronizar
func fazTudo(tarefas []string, done chan bool) {
	for _, tarefa := range tarefas {
		fazAlgoImportante(tarefa)
		time.Sleep(10 * time.Millisecond)
	}
	done <- true
}

func main() {
	tarefas := []string{"dormir", "comer", "correr", "trabalhar", "comer", "dormir"}
	done := make(chan bool)
	go fazTudo(tarefas, done)
	<-done
}
// END OMIT