package main

import (
	"fmt"
	"github.com/gorilla/mux"
<<<<<<< HEAD
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
	router.HandleFunc("/", controllers.HomeHandler)
	//router.HandleFunc("/postform", controllers.PostFormHandler)
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletePersonByID).Methods("DELETE")
=======
	"github.com/pkg/errors"
	"net/http"
	"someApp/pkg/postgresql"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Home sweet home")
	fmt.Fprint(w, response)
}

func postFormHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	age := r.FormValue("userage")

	num, _ := strconv.Atoi(age)
	conn, err := postgresql.NewConnection("0.0.0.0", "5432",
		"postgres", "postgres", "postgres")
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("error in postform: %v\n", err))
	}
	defer conn.CloseConnection()
	postgresql.InsertData(conn, name, num)
	fmt.Fprintf(w, "Name: %s Age: %s", name, age)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "C:\\Users\\iilna\\GolandProjects\\someApp\\cmd\\main\\sth.html")
	})
	router.HandleFunc("/postform", postFormHandler)
	router.HandleFunc("/home", homeHandler)
>>>>>>> a2c8e2f (lets try)
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
