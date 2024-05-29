package app

import (
	"fmt"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/config"
	"github.com/LGuilhermeMoreira/bank_api/internal/infra/database"
	"github.com/LGuilhermeMoreira/bank_api/routes"
)

type app struct{}

func (a *app) Run() {
	config := config.NewConfig()

	conn := database.NewConnection(config.DatabaseName, config.DatabaseUri)

	router := routes.NewRounter(conn, config)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", config.Port),
		Handler: router.Mux,
	}

	// start server
	fmt.Printf("server start :0\n\n")

	server.ListenAndServe()
}

func NewApp() *app {
	return &app{}
}
