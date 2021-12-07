package router

import (
	"github.com/budhirajamadhav/url-shortener/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/shorten", controller.ShortenUrl).Methods("POST")
	return router
}