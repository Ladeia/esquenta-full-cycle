package main

import (
	"fmt"
	"time"
)

func contador(qtd int) {
	for i := 0; i < qtd; i++ { // Idioma do ex não funcionou for i := range qtd
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() { // thread principal
	go contador(10) // Nova thread
	go contador(10) // Nova thread
	contador(10)
}
