package models

import (
	"fmt"
	"net/http"
)

type app struct{}

func (a *app) Run() {
	// start routes
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		phrase := "Ola mundo!\n"
		w.Write([]byte(phrase))
	})

	fmt.Printf("server start :0\n")

	// start server
	http.ListenAndServe(":8080", mux)
}

func NewApp() *app {
	return &app{}
}
