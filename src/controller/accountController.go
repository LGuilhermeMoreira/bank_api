package controller

import (
	"encoding/json"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/internal/dto"
)

type accountController struct {
	conn *database.Connection
}

func NewAccountController(conn *database.Connection) *accountController {
	return &accountController{
		conn: conn,
	}
}

func (a accountController) HandleCreateAccountController(w http.ResponseWriter, r *http.Request) {
	var account dto.AccountInput

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		msg := "Error decoding: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	stmt, err := a.conn.Db.Prepare("insert into accounts(id,login_account_id,owner,balance) values (?,?,?,?)")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	accountModel, err := account.ConvertInputToModel()

	if err != nil {
		msg := "Error converting input to model: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if _, err = stmt.Exec(&accountModel.ID, &accountModel.LoginAccountID, &accountModel.Owner, &accountModel.Balance); err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
