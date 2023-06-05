package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Fulano", 16}
	usuario1.salvar()

	usuario2 := usuario{"Ciclano", 30}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Printf("O usuário %s é maior de idade ? %t", usuario1.nome, maiorDeIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
