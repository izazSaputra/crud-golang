package main

import (
	"log"
	"net/http"
	"src/config"
	homeController "src/controllers"
	"src/controllers/categorycontroller"
)

func main() {
	config.ConnectionDB()

	// HomePage
	http.HandleFunc("/", homeController.Welcome)

	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("server running on port: 8080 ")
	http.ListenAndServe(":8080", nil)
}
