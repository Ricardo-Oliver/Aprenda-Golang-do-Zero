package main

import "fmt"

func main() {
	var variavel1 int = 10

	//Nessa atribuição o valor da varriavel2 é uma copia do valor da variavel1
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	//Nesse incremento, apenas a variavel1 está sendo modificada
	variavel1++
	fmt.Println(variavel1)


	//PONTEIRO É UMA REFERÊNCIA
	var variavel3 int 
	var ponteiro *int
	
	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro)
}