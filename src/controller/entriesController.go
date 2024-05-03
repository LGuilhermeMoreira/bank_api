package controller

import (
	"encoding/json"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/utils/dto"
	"github.com/google/uuid"
)

type entriesController struct {
	conn *database.Connection
}

func NewEntrieController(db *database.Connection) *entriesController {
	return &entriesController{
		conn: db,
	}
}

func (e entriesController) HandleGetAllEntries(w http.ResponseWriter, r *http.Request) {
	rows, err := e.conn.Db.Query("select id,account_id,amount from entries")

	if err != nil {
		msg := "Error querying database" + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var entries []struct {
		id        int
		accountID uuid.UUID
		amount    float64
	}

	for rows.Next() {
		var entrie struct {
			id        int
			accountID uuid.UUID
			amount    float64
		}

		err := rows.Scan(&entrie.id, &entrie.accountID, &entrie.amount)

		if err != nil {
			msg := "error scanning row: " + err.Error()
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		entries = append(entries, entrie)
	}

	if err = json.NewEncoder(w).Encode(entries); err != nil {
		msg := "error encoding: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	// w.WriteHeader(http.StatusOK)
}

func (e entriesController) HandlePostEntrie(w http.ResponseWriter, r *http.Request) {
	requestBody := r.Body

	var dtoEntrie dto.DtoEntrie

	if err := json.NewDecoder(requestBody).Decode(&dtoEntrie); err != nil {
		msg := "error decoding: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	entrie := dtoEntrie.ConvertToModelEntrie()

	stmt, err := e.conn.Db.Prepare("insert into entries(account_id,amount) values (?,?)")

	if err != nil {
		msg := "error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	if _, err = stmt.Exec(&entrie.Account_ID, &entrie.Amount); err != nil {
		msg := "error executing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
