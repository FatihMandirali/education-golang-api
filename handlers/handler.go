package handlers

import (
	. "education.api/services"
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	r := mux.NewRouter().StrictSlash(true)
	//Login
	r.HandleFunc("/api/login", PostLogin).Methods("POST")
	//Admins
	r.HandleFunc("/api/admin", GetAdmins).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
