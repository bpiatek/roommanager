package main

import (
	"fmt"
	"github.com/bpiatek/roommanager/src/order"
	"github.com/bpiatek/roommanager/src/room"
	"github.com/bpiatek/roommanager/src/user"
	"github.com/gorilla/mux"
	"net/http"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public./"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public", fs))
}

func AddApproutes(route *mux.Router) {
	setStaticFolder(route)

	fmt.Println("\nUSER ENDPOINTS:")
	fmt.Println("GET /users => get all users")
	route.HandleFunc("/users", user.GetUsers).Methods("GET")
	fmt.Println("GET /users/{id} => get user by id")
	route.HandleFunc("/users/{id}", user.GetUserById).Methods("GET")
	fmt.Println("POST /users => add user")
	route.HandleFunc("/users", user.AddUser).Methods("POST")
	fmt.Println("POST /users/{id}/card => add card to the user")
	route.HandleFunc("/users/{id}/card", user.AddCardToUser).Methods("POST")

	fmt.Println("\nROOM ENDPOINTS:")
	fmt.Println("GET /rooms => get all rooms")
	route.HandleFunc("/rooms", room.GetRooms).Methods("GET")
	fmt.Println("GET /rooms/{id} => get room by id")
	route.HandleFunc("/rooms/{id}", room.GetRoomById).Methods("GET")
	fmt.Println("POST /rooms => add room")
	route.HandleFunc("/rooms", room.AddRoom).Methods("POST")

	fmt.Println("\nORDERS ENDPOINTS:")
	fmt.Println("POST /users/{id}/orders => reserve room for given user")
	route.HandleFunc("/users/{id}/orders", order.AddOrder).Methods("POST")
	fmt.Println("GET users/{id}/orders => get orders by user id")
	route.HandleFunc("/users/{id}/orders", order.GetOrdersByCustomerId).Methods("GET")
}
