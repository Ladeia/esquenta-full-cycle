package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d recebeu %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {

	canal := make(chan int)

	go worker(1, canal)
	go worker(2, canal)

	for i := 0; i < 10; i++ { // este idioma não funcionou for i := range 10 {
		canal <- i
	}

}
