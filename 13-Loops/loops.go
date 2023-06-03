package main

import "fmt"

func main() {

	i := 0

	for i < 10 {
		i++
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j")
	}

	nomes := []string{"João", "Davi", "Lucas"}

	//Iterando com slice
	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

	//Iterando com string
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string {
		"nome": "Fulano",
		"sobrenome": "Ciclano",
	}

	//Iterando com map
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}