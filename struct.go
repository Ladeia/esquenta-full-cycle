package main

import "fmt"

type Carro struct {
	modelo string
	ano    int
}

func (c Carro) andar() {
	fmt.Println("O carro", c.modelo, "está andando")
}

func main() {
	carro1 := Carro{modelo: "March", ano: 2014}
	//carro2 := Carro{modelo: "March", ano: 2014}

	fmt.Println(carro1.modelo, carro1.ano)
	carro1.andar()
}
