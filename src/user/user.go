package user

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id int					`json:"id"`
	CardId sql.NullInt64	`json:"card_id"`
	Login string			`json:"login"`
	Email string			`json:"email"`
}

type CardId struct {
	Id int `json:"cardId"`
}

func (p User) Print() {
	fmt.Printf("%+v\n", p)
	//fmt.Println(p.Name + " " + p.Surname)
}

func (p *User) UpdateLogin (name string) {
	p.Login = name
}

//func (p *User) SetCard(cardId int) {
//	p.CardId = cardId
//}
