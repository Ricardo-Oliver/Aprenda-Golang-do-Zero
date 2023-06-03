package main

import "fmt"

func main() {

	//Array - Lista de valores do mesmo tipo, com quantidade estático de elementos
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	//Declarando array de forma que seu tamanho vai ser de acordo com a quantidade de elementos passados para ele
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3)

	//Slice - Lista de valores do mesmo tipo, com quantidade dinâmica de elementos
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	//Adicionando um valor no slice
	slice = append(slice, 16)
	fmt.Println(slice)


	//Arrays Internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length 
	fmt.Println(cap(slice3)) // capacidade

}