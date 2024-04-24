package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type app struct{}

// test
type messageJson struct {
	message string `json:"message"`
}

func (a *app) Run() {
	// start routes
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		msg := messageJson{
			message: "ola mundo!",
		}

		result, err := json.Marshal(msg)

		if err != nil {
			panic(err)
		}

		w.Write(result)
	})

	fmt.Printf("server start :0\n")

	// start server
	http.ListenAndServe(":8080", mux)
}

func NewApp() *app {
	return &app{}
}
