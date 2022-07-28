package controllers

import (
	"net/http"
	"testProject/services"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {

	result, err := services.RouteServiceInstance.HandlePing()
	if err != nil {
		http.Error(w, "something went wrong", 500)
	}
	w.Write(result)
}
