package room

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*
	ROOMS
*/
func GetRooms(response http.ResponseWriter, request *http.Request) {
	jsonResponse := FindAll()
	str := string(jsonResponse)

	if str == "null" {
		webError(response, request,"There are no rooms in DB", http.StatusNotFound)
		return
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

func GetRoomById(response http.ResponseWriter, request *http.Request) {
	roomID := mux.Vars(request)["id"]

	i, err := strconv.Atoi(roomID)
	if err != nil {
		fmt.Println(err)
		webError(response, request, "Invalid form of id. It should be numeric", http.StatusBadRequest)
		return
	}

	jsonResponse := getRoomByIdDB(i)
	if jsonResponse == nil {
		webError(response, request,"Room with id: " + roomID + " not found", http.StatusNotFound)
		return
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

func AddRoom(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var roomv Room
	json.Unmarshal(body, &roomv)

	if roomv.Name == "" || roomv.Price == 0 {
		webError(response, request, "Missing parameters.", http.StatusBadRequest)
		return
	}
	_, erro := addRoomDB(roomv)
	if erro != nil {
		if strings.Contains(erro.Error(), ": Duplicate entry") {
			log.Printf("Room already exists in database. Name: %s\n", roomv.Name)
			webError(response, request, "Room '" + roomv.Name + "' already exist.", http.StatusConflict)
			return
		}
	}

	response.WriteHeader(http.StatusCreated)
}

func webError(response http.ResponseWriter, request *http.Request , err string, status int) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	fmt.Fprintf(response, `{"result":"","error":%q}`, err)
}
