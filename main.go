package main

import (
	"github.com/gorilla/mux"
	"log"
	"mods/pkg/routes"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
