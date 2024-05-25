package routes

import (
	"encoding/json"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/config"
	"github.com/LGuilhermeMoreira/bank_api/src/controller"
	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/middleware"
)

type router struct {
	Mux *http.ServeMux
}

func NewRounter(conn *database.Connection, conf *config.Config) *router {
	mux := http.NewServeMux()

	login := controller.NewLoginAccountController(conn, conf)
	account := controller.NewAccountController(conn)
	entrie := controller.NewEntrieController(conn)
	transfer := controller.NewTransferController(conn)
	jwt := middleware.NewJWT(*conf)

	// jwt := middleware.NewJWT()

	//pint - pong
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		msg := map[string]string{
			"ping": "pong",
		}

		bytes, err := json.Marshal(msg)

		if err != nil {
			erro := "Error marshalling json"
			http.Error(w, erro, http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})

	// login
	mux.HandleFunc("POST /login/", middleware.LogMiddleware(login.HandleCreateLoginAccount))
	mux.HandleFunc("POST /login/verify/", middleware.LogMiddleware(login.HandleVerifyLoginAccount))

	// account
	mux.HandleFunc("POST /account/", middleware.LogMiddleware(account.HandleCreateAccountController))
	mux.HandleFunc("GET /account/{id}/", middleware.LogMiddleware(jwt.VerifyAuthToken(account.HandleGetAccountByID)))
	mux.HandleFunc("GET /account/", middleware.LogMiddleware(jwt.VerifyAuthToken(account.HandleGetAllAccounts)))
	mux.HandleFunc("DELETE /account/{id}/", middleware.LogMiddleware(jwt.VerifyAuthToken(account.HandleDeleteAccount)))
	mux.HandleFunc("PUT /account/", middleware.LogMiddleware(jwt.VerifyAuthToken(account.HandleDeleteAccount)))
	// entrie
	mux.HandleFunc("POST /entrie/", middleware.LogMiddleware(jwt.VerifyAuthToken(entrie.HandleCreateEntrie)))
	mux.HandleFunc("GET /entrie/", middleware.LogMiddleware(jwt.VerifyAuthToken(entrie.HandleGetAllEntries)))
	mux.HandleFunc("GET /entrie/{id}/", middleware.LogMiddleware(jwt.VerifyAuthToken(entrie.HandleGetAllAccontEntries)))

	// transfer
	mux.HandleFunc("POST /transfer/", middleware.LogMiddleware(jwt.VerifyAuthToken(transfer.HandleCreateTransfer)))
	mux.HandleFunc("GET /transfer/", middleware.LogMiddleware(jwt.VerifyAuthToken(transfer.HandleGetAllTransfer)))
	mux.HandleFunc("GET /transfer/{id}/", middleware.LogMiddleware(jwt.VerifyAuthToken(transfer.HandleGetAllTransferAccount)))

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
