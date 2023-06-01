package main

import (
	"errors"
	"fmt"
)

func main() {

	//Declarando números inteiros podendo ser int8, int16, int32, int64, e quando informado apenas int o compilador usa a arquitetura da máquina pra executar
	var numero int64 = 100
	fmt.Println(numero)

	//Declarando apenas números positivos, sendo possível usar uint8, uint16, uint32, uint64
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//Alias -> INT32 = RUNE
	var numero3 rune = 1234
	fmt.Println(numero3)

	//Alias -> UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//Declaração de números reais 
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	// Declaração de números reais sendo possível usar apenas float32 e float64
	var numeroReal2 float64 = 12344444.32
	fmt.Println(numeroReal2)

	//Declarando variável do tipo string
	var str string = "Texto"
	fmt.Println(str)

	//Declarando variável do tipo string por inferência
	str2 := "Texto2"
	fmt.Println(str2)

	//Não existe "char" em Go, mas isso é o mais próximo, sendo o resultado da variável "char", o número do caracter na tabela ASCII 
	char := 'A'
	fmt.Println(char)

	//Declarando variável do tipo booleano
	var booleano bool
	fmt.Println(booleano)

	//Em Go erro é um tipo, sendo do tipo "error", e para criar a mensagem de erro, é necessário usar o pacote errors e a função New
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}