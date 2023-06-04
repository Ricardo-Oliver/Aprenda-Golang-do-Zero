package main

import "fmt"

//Função criada para chamar a função recover
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucessso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//A função panic interrompe a execução do programa, que só volta caso tenha uma função recover
	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução")
}