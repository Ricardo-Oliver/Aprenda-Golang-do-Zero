package main

import "fmt"

//Declaração de função passando o nome, depois os parâmetros(com os seus tipos), e se houver retorno colocar o que retorna
func somar(n1 int, n2 int) int {
	return n1 + n2
}

//Declaração de função com mais de um retorno
func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

//Declaração da função main que é a principal do código em qualquer projeto, onde tudo é executado
func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	//Como em Go função é um tipo, podemos declarar variáveis recebendo uma função, logo a variável passa a ser daquele tipo
	var f = func() {
		fmt.Println("Função f")
	}

	//Chamando a variável "f" que é do tipo func
	f()

	//Declarando uma variável com o tipo função que possui um parâmetro
	var t = func(txt string) {
		fmt.Println(txt)
	}

	//Chamando a variável "t" que é do tipo func, e passando um parâmetro de acordo com o tipo
	t("Texto da função")

	//Declarando variáveis para chamar a função que tem mais de retorno
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}