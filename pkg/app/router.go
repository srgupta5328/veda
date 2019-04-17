package app

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", Health).Methods("Get")
	return router
}
