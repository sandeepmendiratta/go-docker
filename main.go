package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func welcome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello and Welcome !"))
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/hi", welcome).Methods("GET")

	http.ListenAndServe(":8081", router)

}
