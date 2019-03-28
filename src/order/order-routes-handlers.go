package order

import (
	"encoding/json"
	"fmt"
	"github.com/bpiatek/roommanager/src/orderdetails"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func AddOrder(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var orderDetails orderdetails.OrderDetailsDTO
	json.Unmarshal(body, &orderDetails)

	log.Println("ROOM ID: ", orderDetails.RoomId)

	if orderDetails.RoomId == 0 || orderDetails.Minutes == 0{
		webError(response, request, "Missing parameters or wrong format.", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(request)
	userID := vars["id"]

	fmt.Println("ID: " + userID)

	i, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err)
		webError(response, request, "Invalid form of id. It should be numeric", http.StatusBadRequest)
		return
	}

	_, erro := makeOrderDB(i, orderDetails)
	if erro != nil {
		webError(response, request, "Error when adding order to DB", http.StatusInternalServerError)
		fmt.Println(erro)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func GetOrdersByCustomerId(response http.ResponseWriter, request *http.Request) {
	//userID := mux.Vars(request)["id"]

	vars := mux.Vars(request)
	userID := vars["id"]

	fmt.Println("ID: " + userID)

	i, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err)
		webError(response, request, "Invalid form of id. It should be numeric", http.StatusBadRequest)
		return
	}

	jsonResponse := getOrdersByCustomerId(i)
	if string(jsonResponse) == "null" {
		webError(response, request,"Orders for user: " + userID + " not found", http.StatusNotFound)
		return
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}


func webError(response http.ResponseWriter, request *http.Request , err string, status int) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	fmt.Fprintf(response, `{"result":"","error":%q}`, err)
}
