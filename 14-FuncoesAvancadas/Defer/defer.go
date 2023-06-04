package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}
//o Defer é executado antes da última instrução do main, ou antes de um return
func main() {
	fmt.Println(alunoEstaAprovado(5, 8))
}
