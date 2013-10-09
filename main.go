package main

import (
	"github.com/gorilla/mux"
	"github.com/wurkhappy/WH-Comments/DB"
	"github.com/wurkhappy/WH-Comments/handlers"
	"labix.org/v2/mgo"
	"net/http"

)

func main() {
	var err error
	DB.Session, err = mgo.Dial(DB.Config["DBURL"])
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.Handle("/agreement/{agreementID}/comments", dbContextMixIn(handlers.CreateComment)).Methods("POST")
	r.Handle("/agreement/{agreementID}/comments", dbContextMixIn(handlers.GetComments)).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":5050", nil)
}

type dbContextMixIn func(http.ResponseWriter, *http.Request, *DB.Context)

func (h dbContextMixIn) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//create the context
	ctx, err := DB.NewContext(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ctx.Close()

	h(w, req, ctx)
}
