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
	entrie := controller.NewEntrieController(conn)
	transfer := controller.NewTransferController(conn)
	// login
	mux.HandleFunc("POST /login/", login.HandleCreateLoginAccount)
	mux.HandleFunc("POST /login/verify/", login.HandleVerifyLoginAccount)
	// account
	mux.HandleFunc("POST /account/", account.HandleCreateAccountController)
	mux.HandleFunc("GET /account/{id}/", account.HandleGetAccountByID)
	mux.HandleFunc("GET /account/", account.HandleGetAllAccounts)
	mux.HandleFunc("DELETE /account/{id}/", account.HandleDeleteAccount)
	mux.HandleFunc("PUT /account/", account.HandleUpdateAccount)
	// entrie
	mux.HandleFunc("POST /entrie/", entrie.HandleCreateEntrie)
	mux.HandleFunc("GET /entrie/", entrie.HandleGetAllEntries)
	mux.HandleFunc("GET /entrie/{id}/", entrie.HandleGetAllAccontEntries)
	// transfer
	mux.HandleFunc("POST /transfer/", transfer.HandleCreateTransfer)
	mux.HandleFunc("GET /transfer/", transfer.HandleGetAllTransfer)
	mux.HandleFunc("GET /transfer/{id}/", transfer.HandleGetAllTransferAccount)

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
