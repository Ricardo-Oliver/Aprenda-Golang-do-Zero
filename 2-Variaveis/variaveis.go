package main

import "fmt"

func main() {
	//Declaração de variável com cláusula "var" (Visível em todo o pacote)
	var variavel1 string = "Variável 1"

	//Declaração de variável com operador curto(gopher), onde o Go irá inferir o tipo de acordo com o valor da variável
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Declarando várias variáveis com apenas uma cláusula "var"
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3, variavel4)

	//Declarando mais de uma variável com o operador curto de declaração
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	//Declarando constante
	const constante1 string = "Constante 1"

	fmt.Println(constante1)

	//Declarando mais de uma constante
	const (
		PI float64 = 3.1415
		semana string = "7 dias"
	)

	fmt.Println(PI, semana)

}