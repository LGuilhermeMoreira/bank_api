package controller

import (
	"encoding/json"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/config"
	"github.com/LGuilhermeMoreira/bank_api/internal/dto"
	"github.com/LGuilhermeMoreira/bank_api/internal/infra/database"
	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginAccountController struct {
	conn *database.Connection
	conf *config.Config
}

func NewLoginAccountController(connection *database.Connection, configuration *config.Config) loginAccountController {
	return loginAccountController{
		conn: connection,
		conf: configuration,
	}
}

func (l loginAccountController) HandleCreateLoginAccount(w http.ResponseWriter, r *http.Request) {
	var login dto.LoginAccountInput

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		msg := "Error decoding: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// refact this.
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

	claims := jwt.MapClaims{
		"AccountID": loginModel.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(l.conf.JwtPassword))

	if err != nil {
		http.Error(w, "Erro ao gerar token JWT", http.StatusInternalServerError)
		return
	}

	loginOutput := map[string]interface{}{
		"id":    loginModel.ID,
		"token": signedToken,
	}

	response, err := json.Marshal(&loginOutput)

	if err != nil {
		msg := "Error marshalling response: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (l loginAccountController) HandleVerifyLoginAccount(w http.ResponseWriter, r *http.Request) {
	var loginInput dto.LoginAccountInput

	if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
		msg := "Error decoding body request: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	var verifyData struct {
		id       string
		password string
	}

	stmt, err := l.conn.Db.Prepare("select id,user_password from login_accounts where user_mail = ?")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if err := stmt.QueryRow(loginInput.UserMail).Scan(&verifyData.id, &verifyData.password); err != nil {
		msg := "Error querying database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if !helpVerifyPassword([]byte(loginInput.UserPassword), []byte(verifyData.password)) {
		msg := "Error when loggin in"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	var accountData struct {
		ID      string
		Owner   string
		Balance string
	}

	stmt, err = l.conn.Db.Prepare("select id,owner,balance from accounts where login_account_id = ?")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if err = stmt.QueryRow(verifyData.id).Scan(&accountData.ID, &accountData.Owner, &accountData.Balance); err != nil {
		msg := "Error querying database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	token, err := helpCreateJWT(accountData.ID, accountData.Owner, accountData.Balance, l.conf.JwtPassword)

	if err != nil {
		msg := "Error generating token: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	response := dto.NewLoginAccountVerify(verifyData.id, accountData.ID, accountData.Owner, accountData.Balance, token)

	bytes, err := json.Marshal(response)

	if err != nil {
		msg := "Error marshalling json: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(bytes)
}

func helpCreateJWT(id, owner, balance, key string) (string, error) {
	claims := jwt.MapClaims{
		"id":      id,
		"owner":   owner,
		"balance": balance,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(key))
}

func helpVerifyPassword(password, hashedPassword []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password); err == nil {
		return true
	}
	return false
}
