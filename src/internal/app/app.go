package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/config"
)

type app struct{}

// test
type messageJson struct {
	Message string `json:"message"`
}

func (a *app) Run() {

	// start a server mux
	mux := http.NewServeMux()

	config := config.NewConfig()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")

		msg := messageJson{
			Message: "ola mundo!",
		}

		result, err := json.Marshal(msg)

		if err != nil {
			panic(err)
		}

		w.Write(result)
	})

	// account routes

	fmt.Printf("server start :0\n")

	// start server
	http.ListenAndServe(fmt.Sprintf(":%v", config.Port), mux)
}

func NewApp() *app {
	return &app{}
}
