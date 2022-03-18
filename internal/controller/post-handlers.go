package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"someApp/domain"
	"someApp/internal/service"
)

type controller struct {
	service service.PostService
}

type PostController interface {
	CreatePerson(w http.ResponseWriter, r *http.Request)
	GetPersonByID(w http.ResponseWriter, r *http.Request)
	UpdatePersonByID(w http.ResponseWriter, r *http.Request)
	DeletePersonByID(w http.ResponseWriter, r *http.Request)
}

func NewController(serv service.PostService) PostController {
	return &controller{service: serv}
}

func (c *controller) CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)

	var person domain.User
	err := json.Unmarshal(requestBody, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := c.service.Create(r.Context(), &person); err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func (c *controller) GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	person, err := c.service.Find(r.Context(), key)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func (c *controller) UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person domain.User

	if err := json.Unmarshal(requestBody, &person); err != nil {
		fmt.Println(err)
		return
	}
	if err := c.service.Update(r.Context(), &person); err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func (c *controller) DeletePersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if err := c.service.Delete(r.Context(), key); err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
