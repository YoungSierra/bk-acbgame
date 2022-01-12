package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type StructHome struct {
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type StructUser struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type JsonUsers struct {
	Status int          `json:"status"`
	Data   []StructUser `json:"data"`
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", GetHome).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")

	fmt.Println("Servidor corriendo...")
	http.ListenAndServe(":5000", router)
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	res := StructHome{}
	res.Status = 200
	res.Title = "Endpoint test ACB Game"
	res.Description = "Endpoint for development testing ACB Game, by Emperador"
	json.NewEncoder(w).Encode(res)
}

//funcion para obtener todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {

	arrayUsers := []StructUser{}

	user := StructUser{}
	for i := 1; i <= 10000; i++ {

		user.Id = i
		user.Firstname = "Firstname " + strconv.Itoa(i)
		user.Lastname = "Lastname " + strconv.Itoa(i)
		user.Email = "Firstname" + strconv.Itoa(i) + "@gmail.com"
		arrayUsers = append(arrayUsers, user)

	}

	res := JsonUsers{}
	res.Status = 200
	res.Data = arrayUsers
	json.NewEncoder(w).Encode(res)
}
