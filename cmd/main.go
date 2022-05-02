package main

import (
	"net/http"
	"origin-api/controllers"
	"origin-api/getconf"

	"github.com/gorilla/mux"
)

func Server(router *mux.Router) *http.Server {
	return &http.Server{
		Addr:    getconf.Server.Addr,
		Handler: router,
	}
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/test", controllers.CheckToken(http.HandlerFunc(controllers.Test)))
	return router
}

func main() {
	Server(Router()).ListenAndServe()
}
