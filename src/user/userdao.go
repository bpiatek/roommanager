package user

import (
	"fmt"
	"github.com/bpiatek/roommanager/src"
	"log"
)

func addUserDB(user User) (b bool, err error) {
	stmtIns, err := src.GetDB().Prepare("INSERT INTO customer VALUES( ?, ?, ?, ? )")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return false, err
	}

	defer stmtIns.Close()


	_, err = stmtIns.Exec(0, user.Login, user.Email, user.CardId)
	if err != nil {
		panic(err.Error())
		return false, err
	}
	log.Printf("Added user to database. Login: %s\n", user.Login)

	return  true, err
}

func addCardIdToPersonDB(cardId int, userId int) {
	stmtUpd, err := src.GetDB().Prepare("UPDATE customer set cardId=? where id=?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtUpd.Close()

	_, err = stmtUpd.Exec(cardId, userId)
}

func findAllDB() []User {
	var (
		user  User
		users []User
	)

	rows, err := src.GetDB().Query("SELECT * FROM customer")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Login, &user.Email, &user.CardId)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users
}


func GetUserByIdDB(id int) (u User, r error) {
	var user User

	rows, err := src.GetDB().Prepare("SELECT * FROM customer where id=?")
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	defer rows.Close()

	err = rows.QueryRow(id).Scan(&user.Id, &user.Login, &user.Email, &user.CardId)
	if err != nil {
		fmt.Println(err.Error())
	}

	if user.Id == 0 {
		return user, err
	}

	return user, err
}
