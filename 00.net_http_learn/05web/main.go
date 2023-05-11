package main

import (
	"net/http"
	"source/a/01_http/05web/controller"
)

func main() {

	controller.RegisterRoutes()
	server := http.Server{Addr: ":9091", Handler: nil}
	server.ListenAndServe()
}
