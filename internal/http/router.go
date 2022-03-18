package router

import "net/http"

type Router interface {
	POST(path string, f func(http.ResponseWriter, *http.Request))
	GET(path string, f func(http.ResponseWriter, *http.Request))
	PUT(path string, f func(http.ResponseWriter, *http.Request))
	DELETE(path string, f func(http.ResponseWriter, *http.Request))
	SERVE(port string)
}
