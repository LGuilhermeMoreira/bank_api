package controller

import (
	"encoding/json"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

type loginAccountController struct {
	conn *database.Connection
}

func NewLoginAccountController(connection *database.Connection) loginAccountController {
	return loginAccountController{
		conn: connection,
	}
}

func (l loginAccountController) HandleCreateLoginAccount(w http.ResponseWriter, r *http.Request) {
	var login dto.LoginAccountInput

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		msg := "Error decoding: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	stmt, err := l.conn.Db.Prepare("INSERT INTO login_accounts(id,user_mail,user_password) values (?,?,?)")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	defer stmt.Close()

	loginModel, err := login.ConvertInputToModel()

	if err != nil {
		msg := "Error converting input to model: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if _, err = stmt.Exec(&loginModel.ID, &loginModel.UserMail, &loginModel.UserPassword); err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (l loginAccountController) HandleVerifyLoginAccount(w http.ResponseWriter, r *http.Request) {
	var login dto.LoginAccountInput

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		msg := "Error decoding: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	stmt, err := l.conn.Db.Prepare("select user_password from login_accounts where user_mail = ?")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer stmt.Close()
	var password string

	if err = stmt.QueryRow(login.UserMail).Scan(&password); err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(login.UserPassword))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
