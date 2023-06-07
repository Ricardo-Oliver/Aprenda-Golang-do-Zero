package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Programando em Go!")

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			fmt.Sprintln("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}