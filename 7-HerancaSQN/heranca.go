package main

import "fmt"

//Em Go não existe o conceito de herança como temos em outras linguagens, então o conceito
//apresentado a seguir, seria o mais próximo do conceito de herança que existe, em Go.

//Declarando uma struct do tipo pessoa
type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

//Nessa struct o primeiro atributo seria toda a struct de pessoa declarada anteriormente, e basta
//passar apenas o nome da struct sem a necessidade de informar o tipo (pois a struct ja é do tipo pessoa)
type estudante struct {
	pessoa
	faculdade string
	curso string
}

func main() {

	//Criando uma variável para receber a struct de pessoa
	p1 := pessoa{"Cidadão", "Fulano", 20, 180}
	fmt.Println(p1)

	//Criando uma variável para receber a struct pessoa já declarada, junto com a struct estudante
	e1 := estudante{p1, "ADS", "USP"}
	fmt.Println(e1)
}