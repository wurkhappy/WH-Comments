package main

import (
	"github.com/gorilla/mux"
	"github.com/wurkhappy/WH-Comments/handlers"
	"net/http"

)

func main() {
	r := mux.NewRouter()
	r.Handle("/agreement/{agreementID}/comments", dbContextMixIn(handlers.CreateComment)).Methods("POST")
	r.Handle("/agreement/{agreementID}/comments", dbContextMixIn(handlers.GetComments)).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":5050", nil)
}

type dbContextMixIn func(http.ResponseWriter, *http.Request)

func (h dbContextMixIn) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h(w, req)
}
