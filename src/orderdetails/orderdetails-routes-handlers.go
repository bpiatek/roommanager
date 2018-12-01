package orderdetails

import (
	"fmt"
	"net/http"
)

func AddOrderDetails(response http.ResponseWriter, request *http.Request) {
	//vars := mux.Vars(request)
	//userID := vars["userId"]
	//orderID := vars["orderId"]
	//
	//usrID, err := strconv.Atoi(userID)
	//fmt.Println(usrID)
	//if err != nil {
	//	fmt.Println(err)
	//	webError(response, request, "Invalid form of userID. It should be numeric", http.StatusBadRequest)
	//	return
	//}
	//
	//ordrID, erro := strconv.Atoi(orderID)
	//if erro != nil {
	//	fmt.Println(erro)
	//	webError(response, request, "Invalid form orderID. It should be numeric", http.StatusBadRequest)
	//	return
	//}
	//
	//
	//body, err := ioutil.ReadAll(request.Body)
	//if err != nil {
	//	panic(err)
	//}
	//var ordrDet OrderDetailsDTO
	//json.Unmarshal(body, &ordrDet)
	//
	//fmt.Println(ordrDet.RoomId)
	//fmt.Println(ordrDet.Minutes)

	//_, errore := addOrderDetailsDB(ordrID, ordrDet.RoomId, ordrDet.Minutes)
	//if errore != nil {
	//	fmt.Println(errore)
	//}
}

func webError(response http.ResponseWriter, request *http.Request , err string, status int) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	fmt.Fprintf(response, `{"result":"","error":%q}`, err)
}