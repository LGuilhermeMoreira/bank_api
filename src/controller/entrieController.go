package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/internal/dto"
	"github.com/google/uuid"
)

type entrieController struct {
	conn *database.Connection
}

func NewEntrieController(conn *database.Connection) *entrieController {

	return &entrieController{
		conn: conn,
	}
}
func (e entrieController) HandleCreateEntrie(w http.ResponseWriter, r *http.Request) {
	var entrie dto.EntrieInput

	if err := json.NewDecoder(r.Body).Decode(&entrie); err != nil {
		msg := "Error decoding the request body: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Convert amount to float
	amount, err := strconv.ParseFloat(entrie.Amount, 64)
	if err != nil {
		msg := "Error parsing amount to float: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Get current balance of the account
	var currentBalance float64
	err = e.conn.Db.QueryRow("SELECT balance FROM accounts WHERE id = ?", entrie.AccountID).Scan(&currentBalance)
	if err != nil {
		msg := "Error getting current balance from the database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// Calculate new balance
	newBalance := currentBalance + amount

	// Update balance in the accounts table
	_, err = e.conn.Db.Exec("UPDATE accounts SET balance = ? WHERE id = ?", newBalance, entrie.AccountID)
	if err != nil {
		msg := "Error updating account balance in the database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// Insert new entry into entries table
	stmt, err := e.conn.Db.Prepare("INSERT INTO entries (id, account_id, amount) VALUES (?, ?, ?)")
	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	entrieModel, err := entrie.ConvertInputToModel()
	if err != nil {
		msg := "Error converting to a model: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if _, err := stmt.Exec(&entrieModel.ID, &entrieModel.AccountID, &entrieModel.Amount); err != nil {
		msg := "Error running database query: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (e entrieController) HandleGetAllEntries(w http.ResponseWriter, r *http.Request) {
	var entries []struct {
		ID        uuid.UUID `json:"id"`
		AccountID uuid.UUID `json:"account_id"`
		Amount    float64   `json:"amount"`
	}

	stmt, err := e.conn.Db.Prepare("select id,account_id,amount from entries")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query()

	if err != nil {
		msg := "Error querying database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var entrie struct {
			ID        uuid.UUID `json:"id"`
			AccountID uuid.UUID `json:"account_id"`
			Amount    float64   `json:"amount"`
		}

		if err = rows.Scan(&entrie.ID, &entrie.AccountID, &entrie.Amount); err != nil {
			msg := "Error scanning database: " + err.Error()
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		entries = append(entries, entrie)
	}

	response, err := json.Marshal(&entries)

	if err != nil {
		msg := "Error marshalling response: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// pegar todas as entradas de um usuario
func (e entrieController) HandleGetAllAccontEntries(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	stmt, err := e.conn.Db.Prepare("select id,amount,created_at from entries where account_id = ?")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var entries []struct {
		ID       uuid.UUID `json:"id"`
		Amount   float64   `json:"amount"`
		CreateAt time.Time `json:"created_at"`
	}

	for rows.Next() {
		var entrie struct {
			ID       uuid.UUID `json:"id"`
			Amount   float64   `json:"amount"`
			CreateAt time.Time `json:"created_at"`
		}

		if err = rows.Scan(&entrie.ID, &entrie.Amount, &entrie.CreateAt); err != nil {
			msg := "Error scanning database: " + err.Error()
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		entries = append(entries, entrie)
	}

	response, err := json.Marshal(&entries)

	if err != nil {
		msg := "Error marshalling response: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// mudar o post para somar ao valor atual
