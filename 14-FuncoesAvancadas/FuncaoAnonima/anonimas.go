package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá Mundo!")
	}()

	//Função anônima com parâmetro
	func(texto string) {
		fmt.Println(texto)
	}("Passando parâmetro")

	//Função anônima com retorno
	retorno := func(texto string) string{
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")
	
	fmt.Println(retorno)

}