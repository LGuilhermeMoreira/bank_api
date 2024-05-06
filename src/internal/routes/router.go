package routes

import (
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/controller"
	"github.com/LGuilhermeMoreira/bank_api/src/database"
)

type router struct {
	Mux *http.ServeMux
}

func NewRounter(conn *database.Connection) *router {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login/", controller.NewLoginAccountController(conn).HandleCreateLoginAccount)
	mux.HandleFunc("POST /login/verify/", controller.NewLoginAccountController(conn).HandleVerifyLoginAccount)
	mux.HandleFunc("POST /account/", controller.NewAccountController(conn).HandleCreateAccountController)

	return &router{
		Mux: mux,
	}
}

// middleware example
// func alwaysShow(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(r.Method, r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	}

// }
