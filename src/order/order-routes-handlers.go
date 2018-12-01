package order

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddOrder(response http.ResponseWriter, request *http.Request) {
	//body, err := ioutil.ReadAll(request.Body)
	//if err != nil {
	//	panic(err)
	//}
	//var orderv Order
	//json.Unmarshal(body, &orderv)


	vars := mux.Vars(request)
	userID := vars["id"]

	fmt.Println("ID: " + userID)

	i, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err)
		webError(response, request, "Invalid form of id. It should be numeric", http.StatusBadRequest)
		return
	}
	//
	//jsonResponse := user.GetUserByIdDB(i)
	//
	//fmt.Println(string(jsonResponse))
	//
	//if jsonResponse == nil {
	//	webError(response, request,"User with id: " + userID + " not found", http.StatusNotFound)
	//	return
	//}

	//if orderv.CustomerId == 0 {
	//	webError(response, request, "Missing parameters.", http.StatusBadRequest)
	//	return
	//}
	_, erro := makeOrderDB(i)
	if erro != nil {
		//if strings.Contains(erro.Error(), ": Duplicate entry") {
		//	log.Printf("Room already exists in database. Name: %s\n", i)
		//	webError(response, request, "Room '" + roomv.Name + "' already exist.", http.StatusConflict)
		//	return
		//}
		fmt.Println(erro)
		return
	}
	//
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
