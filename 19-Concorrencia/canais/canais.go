package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {

	for i := 0; i < 5; i++ {
		canal <- texto // Mandando um valor pra dentro do canal
		time.Sleep(time.Second)
	}

	close(canal)
}
