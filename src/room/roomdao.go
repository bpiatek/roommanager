package room

import (
	"encoding/json"
	"fmt"
	"github.com/bpiatek/roommanager/src"
	"log"
)

func addRoomDB(room Room) (b bool, err error) {
	stmtIns, err := src.GetDB().Prepare("INSERT INTO rooms VALUES( ?, ?, ? )")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtIns.Close()

	_, err = stmtIns.Exec(0, room.Name, room.Price)
	if err != nil {
		return false, err
	}
	log.Printf("Added room to database. Login: %s\n", room.Name)

	return  true, err
}


func FindAll() []byte {
	var (
		room  Room
		rooms []Room
	)
	rows, err := src.GetDB().Query("SELECT * FROM rooms")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for rows.Next() {
		err = rows.Scan(&room.Id, &room.Name, &room.Price)
		if err != nil {
			panic(err.Error())
		}

		rooms = append(rooms, room)
	}

	defer rows.Close()

	jsonResponse, jsonError := json.Marshal(rooms)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}

	return jsonResponse
}

func getRoomByIdDB(id int) []byte {
	var room Room
	rows, err := src.GetDB().Prepare("SELECT * FROM rooms where id=?")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = rows.QueryRow(id).Scan(&room.Id, &room.Name, &room.Price)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	if room.Id == 0{
		return nil
	}

	jsonResponse, jsonError := json.Marshal(room)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}

	return jsonResponse
}
