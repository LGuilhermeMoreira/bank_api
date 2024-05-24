package app

import (
	"fmt"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/config"
	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/utils/routes"
)

type app struct{}

func (a *app) Run() {
	config := config.NewConfig()

	conn := database.NewConnection(config.DatabaseName, config.DatabaseUri)

	router := routes.NewRounter(conn)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", config.Port),
		Handler: router.Mux,
	}

	// start server
	fmt.Printf("server start :0\n")

	server.ListenAndServe()
}

func NewApp() *app {
	return &app{}
}
