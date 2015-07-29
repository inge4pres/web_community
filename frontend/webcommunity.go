package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

func main() {
	router.HandleFunc("/", indexHandler)
	//TODO
	http.Handle("/", router)
}
