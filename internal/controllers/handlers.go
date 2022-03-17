package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"someApp/internal/user"
	"someApp/pkg/postgresql"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cmd/main/sth.html")
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)

	var person user.User
	err := json.Unmarshal(requestBody, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := postgresql.NewConnection("localhost", "5432",
		"postgres", "postgres", "postgres")
	defer conn.CloseConnection()
	if err != nil {
		fmt.Println(errors.Wrap(err, fmt.Sprintf("error in /create: %v\n", err)))
		return
	}
	fmt.Println(person)
	if err := conn.InsertData(person); err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person []user.User

	conn, err := postgresql.NewConnection("localhost", "5432",
		"postgres", "postgres", "postgres")
	defer conn.CloseConnection()
	if err != nil {
		fmt.Println(errors.Wrap(err, fmt.Sprintf("error in /get: %v\n", err)))
		return
	}
	person, _ = conn.Get(key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person user.User

	err := json.Unmarshal(requestBody, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := postgresql.NewConnection("localhost", "5432",
		"postgres", "postgres", "postgres")
	defer conn.CloseConnection()
	if err != nil {
		fmt.Println(errors.Wrap(err, fmt.Sprintf("error in /update: %v\n", err)))
		return
	}
	conn.Update(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func DeletePersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	conn, err := postgresql.NewConnection("localhost", "5432",
		"postgres", "postgres", "postgres")
	defer conn.CloseConnection()
	if err != nil {
		fmt.Println(errors.Wrap(err, fmt.Sprintf("error in /delete: %v\n", err)))
		return
	}
	conn.Delete(key)
	w.WriteHeader(http.StatusNoContent)
}
