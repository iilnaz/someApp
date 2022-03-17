package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"someApp/internal/controllers"
	"someApp/schema/migrations"
)

func main() {
	mg, err := migrations.NewMigration("localhost", "5432",
		"postgres", "postgres", "postgres")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := mg.Up(); err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletePersonByID).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}
