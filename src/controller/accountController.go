package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/utils/dto"
)

type accountController struct {
	conn *database.Connection
}

func NewAccountController(db *database.Connection) *accountController {
	return &accountController{
		conn: db,
	}
}

func (a accountController) HandlePostAccount(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	// criando o modelo que vai receber o desserializado
	var accountDTO dto.DtoAccount

	// desserializando
	if err := json.NewDecoder(body).Decode(&accountDTO); err != nil {
		msg := "unmarshal not completed"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatalln(msg)
		return
	}

	account := accountDTO.ConvertToModelAccount()

	// fazer o save no banco de dados

	stmt, err := a.conn.Db.Prepare("insert into accounts(id,owner,balance) values (?,?,?)")

	if err != nil {
		msg := "error preparing the query"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatalln(msg)
		return
	}

	defer stmt.Close()

	if _, err := stmt.Exec(account.ID, account.Owner, account.Balance); err != nil {
		msg := "error executing the query"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatalln(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Println("conta cadastrada")

}

func (a accountController) HandleGetAccountByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var queryResult struct {
		ID      string `json:"id"`
		Owner   string `json:"owner"`
		Balance string `json:"balance"`
	}
	stmt, err := a.conn.Db.Prepare("Select id,owner,balance from accounts where id = ?")

	if err != nil {
		msg := "error preparing the query"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatalln(msg)
		return
	}

	defer stmt.Close()

	if err = stmt.QueryRow(id).Scan(&queryResult.ID, &queryResult.Owner, &queryResult.Balance); err != nil {
		msg := "error executing the query"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatalln(msg)
		return
	}

	// sending the json
	if err := json.NewEncoder(w).Encode(&queryResult); err != nil {
		msg := "error marshaling the struct"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Fatalln(msg)
		return
	}
}

func (a accountController) HandleGetAllAccount(w http.ResponseWriter, r *http.Request) {
	var accounts []struct {
		ID      string  `json:"id"`
		Owner   string  `json:"owner"`
		Balance float64 `json:"balance"` // Use float64 for balance
	}

	rows, err := a.conn.Db.Query("SELECT id, owner, balance FROM accounts")
	if err != nil {
		msg := "error querying accounts: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var account struct {
			ID      string  `json:"id"`
			Owner   string  `json:"owner"`
			Balance float64 `json:"balance"`
		}
		err := rows.Scan(&account.ID, &account.Owner, &account.Balance)
		if err != nil {
			msg := "error scanning row: " + err.Error()
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		accounts = append(accounts, account)
	}

	err = json.NewEncoder(w).Encode(accounts)
	if err != nil {
		msg := "error encoding response: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}