package app

import (
	"fmt"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/config"
	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/middleware"
	"github.com/LGuilhermeMoreira/bank_api/src/utils/routes"
)

type app struct{}

func (a *app) Run() {
	config := config.NewConfig()

	conn := database.NewConnection()

	router := routes.NewRounter(conn)

	http.Handle("/", middleware.LogMiddleware(router.Mux))

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
