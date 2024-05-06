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

	login := controller.NewLoginAccountController(conn)
	account := controller.NewAccountController(conn)
	mux.HandleFunc("POST /login/", login.HandleCreateLoginAccount)
	mux.HandleFunc("POST /login/verify/", login.HandleVerifyLoginAccount)
	mux.HandleFunc("POST /account/", account.HandleCreateAccountController)
	mux.HandleFunc("GET /account/{id}", account.HandleGetAccountByID)
	mux.HandleFunc("GET /account/", account.HandleGetAllAccounts)

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
