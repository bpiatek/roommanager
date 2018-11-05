package dupa

import "fmt"

type Person struct {
	CardId string
	Name string
	Surname string
}

func (p Person) Print() {
	fmt.Printf("%+v\n", p)
	//fmt.Println(p.Name + " " + p.Surname)
}

func (p *Person) UpdateName (name string) {
	p.Name = name
}

func (p *Person) SetCard(cardId string) {
	p.CardId = cardId
}
