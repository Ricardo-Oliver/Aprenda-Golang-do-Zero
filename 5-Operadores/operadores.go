package main

import "fmt"

func main() {
	//OPERADORES ARITMETICOS ((+)adição, (-)subtração, (*)multiplicação, (/)divisão, (%)resto da divisão)
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 5
	divisao := 10 / 4
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//OPERADORES DE ATRIBUIÇÃO (=, :=)
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS (>, <, >=, <=, ==, !=, sempre retorna um booleano)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS ( (&&)and, (||)ou, (!)negação )
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)

	//OPERADORES UNÁRIOS (++, +=, --, -=, *=, /=, %=)
	numero := 10

	//Incremento
	numero++

	//numero = numero + 2
	numero += 2

	//Decremento
	numero--

	//numero = numero - 2
	numero -= 2

	//numero = numero * 3
	numero *= 3

	//numero = numero /3
	numero /= 3

	//numero = numero % 3
	numero %= 3

}