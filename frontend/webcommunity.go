package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

func main() {
	//STATIC CONTENT
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static/"))))
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/signin", signInHandler)
	router.HandleFunc("/signup", signUpHandler)
	//TODO
	http.Handle("/", router)

	http.ListenAndServe(":80", router)
}
