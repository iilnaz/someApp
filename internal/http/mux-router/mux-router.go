package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"someApp/internal/http"
)

type muxRouter struct {
	r *mux.Router
}

func NewMuxRouter() router.Router {
	return &muxRouter{r: mux.NewRouter()}
}

func (m muxRouter) POST(path string, f func(http.ResponseWriter, *http.Request)) {
	m.r.HandleFunc(path, f).Methods("POST")
}

func (m muxRouter) GET(path string, f func(http.ResponseWriter, *http.Request)) {
	m.r.HandleFunc(path, f).Methods("GET")
}

func (m muxRouter) PUT(path string, f func(http.ResponseWriter, *http.Request)) {
	m.r.HandleFunc(path, f).Methods("PUT")
}

func (m muxRouter) DELETE(path string, f func(http.ResponseWriter, *http.Request)) {
	m.r.HandleFunc(path, f).Methods("DELETE")
}

func (m muxRouter) SERVE(port string) {
	fmt.Println("mux http server is listening")
	http.ListenAndServe(port, m.r)
}
