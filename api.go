package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"	
)
type APIServer struct {
	addr string // eg port:3000
	store Store // point/repo where we have connection to the database
}



// constructor for the server 
func NewAPIServer(addr string , store Store) * APIServer {
	return &APIServer{addr: addr, store: store}
}

// this will initilize mux_router and register all services
func(s*APIServer) Serve() {
	router := mux.NewRouter()	
	subrouter := router.PathPrefix("/api/v1").Subrouter() 


	// here service are regsitered 

	log.Println("Starting the API Server at ",s.addr)

	log.Fatal(http.ListenAndServe(s.addr,subrouter))
}