package main

import "fmt"

//Só é possível 1 parâmetro variático por função e deve ser declarado por último obrigatoriamente
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total +=numero
	}
	return total
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalDaSoma)
}