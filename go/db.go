package main

import (
	"database/sql"
)

type Carro struct {
	modelo string
	ano    int
}

func insertCarro(db *sql.DB, carro Carro) {
	_, err := db.Exec("Insert Into carro (ano, modelo) values (?, ?)", carro.modelo, carro.ano)

	if err != nil {
		panic(err)
	}
}

func main() {
	// sqlite3
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	carro1 := Carro{modelo: "March", ano: 2014}

	insertCarro(db, carro1)
}
