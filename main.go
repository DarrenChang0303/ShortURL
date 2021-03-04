package main

import (
	"log"
	"net/http"
	controller "shortURL/controller"

	"github.com/gorilla/mux"
)

func main() {
	//create new router
	router := mux.NewRouter()
	//Show homepage
	router.HandleFunc("/", controller.Original).Methods("PUT")
	//Create short URL
	router.HandleFunc("/create", controller.CreateURL).Methods("GET")
	//Redirect original URL
	router.HandleFunc("/redirect", controller.Redirect).Methods("GET")
	//Server listen port
	log.Fatal(http.ListenAndServe(":9527", router))

}
