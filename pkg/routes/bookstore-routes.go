package routes

import (
	"fmt"	
	"github.com/gorilla/mux"
	"mods/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	fmt.Println("Running in Port 3000")
	router.HandleFunc("/", controllers.HomePage)
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
