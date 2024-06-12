package main

import "fmt"

func main() { // Thread principal
	canal := make(chan string)

	go func() { // Instancia nova thread
		canal <- "olá canal"
	}()

	fmt.Println(<-canal) // na 1ª thread imprime o conteúdo de canal
}
