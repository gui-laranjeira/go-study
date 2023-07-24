package main

import (
	"fmt"
)

type person struct {
	name        string `json:"name"`
	lastName    string `json:"lastName"`
	age         int    `json:"age"`
	contactInfo `json:"contact"`
}

type contactInfo struct {
	email string `json:"email"`
	phone string `json:"phone"`
}

func main() {
	gui := person{
		name:     "Guilherme",
		lastName: "Laranjeira",
		age:      23,
		contactInfo: contactInfo{
			email: "teste@teste.com",
			phone: "123456789",
		},
	}
	p := &gui
	gui.updateName("Bia")
	gui.print()

	fmt.Println(p)
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(name string) {
	(*p).name = name
}
