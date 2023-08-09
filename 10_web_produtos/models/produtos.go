package models

import (
	"log"

	"web_produtos/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetProdutos() []Produto {
	d := db.DbConnect()
	defer d.Close()

	rows, err := d.Query("SELECT * FROM produtos")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	produtos := []Produto{}

	for rows.Next() {
		tmp := Produto{}

		err := rows.Scan(&tmp.Id, &tmp.Nome, &tmp.Descricao, &tmp.Preco, &tmp.Quantidade)
		if err != nil {
			log.Println(err)
		}

		produtos = append(produtos, tmp)
	}
	return produtos
}
