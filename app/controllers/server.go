package controllers

import (
	"net/http"
	"todo_app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top) //top

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
