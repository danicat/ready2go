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

func fazTudo(tarefas []string) {
	for _, tarefa := range tarefas {
		fazAlgoImportante(tarefa)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	tarefas := []string{"dormir", "comer", "correr", "trabalhar", "comer", "dormir"}
	fazTudo(tarefas)
}
// END OMIT
