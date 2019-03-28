package user

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
	USERS
*/
func GetUsers(response http.ResponseWriter, request *http.Request) {
	users := findAllDB()

	if len(users) < 1 {
		webError(response, request,"There are no users in DB", http.StatusNotFound)
		return
	}

	jsonUsers, jsonError := json.Marshal(users)

	if jsonError != nil {
		fmt.Println(jsonError)
		webError(response, request, "Internal Error parsing to json", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonUsers)

}

func AddCardToUser(response http.ResponseWriter, request *http.Request) {
	userID := mux.Vars(request)["id"]

	usrID, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err)
		webError(response, request, "Invalid form of id. It should be numeric", http.StatusBadRequest)
		return
	}

	fmt.Println(usrID)

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var cardId CardId
	json.Unmarshal(body, &cardId)

	fmt.Println(cardId)
	addCardIdToPersonDB(cardId.Id, usrID)
}

func GetUserById(response http.ResponseWriter, request *http.Request) {
	userID := mux.Vars(request)["id"]

	id, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err)
		webError(response, request, "Invalid form of id. It should be numeric", http.StatusBadRequest)
		return
	}

	user, erro := GetUserByIdDB(id)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(user)
	jsonResponse, jsonError := json.Marshal(user)
	if jsonError != nil {
		fmt.Println(jsonError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonResponse)
}

func AddUser(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var usr User
	json.Unmarshal(body, &usr)

	if usr.Login == "" || usr.Email == "" {
		webError(response, request, "Missing parameters.", http.StatusBadRequest)
		return
	}
	_, erro := addUserDB(usr)
	if erro != nil {
		if strings.Contains(erro.Error(), ": Duplicate entry") {
			log.Printf("User or email already exists in database. Login: %s\n", usr.Login)
			webError(response, request, "User '" + usr.Login + "' already exist.", http.StatusConflict)
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

