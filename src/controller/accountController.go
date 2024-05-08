package controller

import (
	"encoding/json"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/internal/dto"
	"github.com/google/uuid"
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

func (a accountController) HandleGetAccountByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var account struct {
		ID             uuid.UUID `json:"id"`
		LoginAccountID uuid.UUID `json:"login_account_id"`
		Owner          string    `json:"owner"`
		Balance        float64   `json:"balance"`
	}

	stmt, err := a.conn.Db.Prepare("select id,login_account_id,owner,balance from accounts where id = ?")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	if err = stmt.QueryRow(id).Scan(&account.ID, &account.LoginAccountID, &account.Owner, &account.Balance); err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	response, err := json.Marshal(&account)

	if err != nil {
		msg := "Error marshalling the json response: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (a accountController) HandleGetAllAccounts(w http.ResponseWriter, r *http.Request) {
	stmt, err := a.conn.Db.Prepare("select id,login_account_id,owner,balance from accounts")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		msg := "Error querying database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var accounts []struct {
		ID             uuid.UUID
		LoginAccountID uuid.UUID
		Owner          string
		Balance        float64
	}

	for rows.Next() {
		var account struct {
			ID             uuid.UUID
			LoginAccountID uuid.UUID
			Owner          string
			Balance        float64
		}

		if err = rows.Scan(&account.ID, &account.LoginAccountID, &account.Owner, &account.Balance); err != nil {
			msg := "Error querying database: " + err.Error()
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		accounts = append(accounts, account)
	}

	response, err := json.Marshal(&accounts)

	if err != nil {
		msg := "Error querying database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (a accountController) HandleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	stmt, err := a.conn.Db.Prepare("delete from accounts where id = ?")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	if _, err = stmt.Query(id); err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a accountController) HandleUpdateAccount(w http.ResponseWriter, r *http.Request) {
	bodyRequest := r.Body

	var Owner struct {
		ID    uuid.UUID `json:"id"`
		Owner string    `json:"owner"`
	}

	if err := json.NewDecoder(bodyRequest).Decode(&Owner); err != nil {
		msg := "Error decoding body request: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	stmt, err := a.conn.Db.Prepare("update accounts set owner = ? where id = ?")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	if _, err = stmt.Query(Owner.Owner, Owner.ID); err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
