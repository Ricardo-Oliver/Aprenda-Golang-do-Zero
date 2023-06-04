package main 

import "fmt"

//A função init é executada antes da função main, independente de onde ela esteja no código, do mesmo arquivo.
//E ao contrário da função main, podemos ter uma função init para cada arquivo
//Geralmente é usa pra fazer algum setup, e inicializar variáveis de ambiente
func init() {
	fmt.Println("Executando a função init")
}
func main() {
	fmt.Println("Função main sendo executada")
}