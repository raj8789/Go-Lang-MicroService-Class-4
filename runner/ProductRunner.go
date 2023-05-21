package main

import (
	handler "MircoHandler4"
	"log"
	"net/http"
	"os"
)

func main() {
	sm := http.NewServeMux()
	logger := log.New(os.Stdout, "Coffee -API", log.LstdFlags)
	ph := handler.NewProductHandlerList(logger, "Raj")
	sm.Handle("/", ph)
	http.ListenAndServe(":8080", sm)
}
