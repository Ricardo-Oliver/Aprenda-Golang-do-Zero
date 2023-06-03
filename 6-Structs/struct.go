package main

import "fmt"

//Criando uma struct 
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	//Declarando uma variável do tipo usuario que pertence a struct criada
	var u usuario
	u.nome = "Ricardo"
	u.idade = 26
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Alpha", 30}

	//Declarando uma variável, e recebendo a struct usuario aninhada com a struct de endereço
	usuario2 := usuario{"Ricardo", 26, enderecoExemplo}
	fmt.Println(usuario2)

	//Declarando variável e atribuindo um valor específico de um atributo da struct
	usuario3 := usuario{idade: 26}
	fmt.Println(usuario3)
}
