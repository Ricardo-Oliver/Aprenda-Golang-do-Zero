package main

import "fmt"

func main() {

	//Declarando maps: primeiro o tipo das chaves detro dos parênteses, e depois o tipo dos valores fora
	usuario := map[string]string{
		"nome":      "Fulano",
		"sobrenome": "Ciclano",
	}

	fmt.Println(usuario)

	//Aninhando maps
	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "A",
			"ultimo": "B",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	//Deletando uma chave do map
	delete(usuario2, "nome")

	//Inserindo chave no map
	usuario2["signo"] = map[string] string {
		"nome": "Gêmeos",
	}
}