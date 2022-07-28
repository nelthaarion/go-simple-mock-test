package main

import (
	"net/http"
	"testProject/controllers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", controllers.PingHandler)

	http.ListenAndServe(":3000", mux)
}
