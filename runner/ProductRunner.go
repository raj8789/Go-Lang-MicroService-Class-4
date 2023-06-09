package main

import (
	handler "MircoHandler4"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func main() {
	//sm := http.NewServeMux()
	logger := log.New(os.Stdout, "Coffee -API", log.LstdFlags)
	ph := handler.NewProductHandlerList(logger, "Raj")

	sm:=mux.NewRouter()
	
	getsubRouter:=sm.Methods(http.MethodGet).Subrouter()
	getsubRouter.HandleFunc("/",ph.GetProducts)

	putsubRouter:=sm.Methods(http.MethodPut).Subrouter()
	putsubRouter.HandleFunc("/{id:[0-9]+}",ph.UpdateProduct)
	putsubRouter.Use(ph.MiddlewareProductValidation)

	postRouter:=sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/",ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	//sm.Handle("/", ph)
	http.ListenAndServe(":8080", sm)
}
