package controller

import (
	"encoding/json"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/database"
	"github.com/LGuilhermeMoreira/bank_api/src/internal/dto"
	"github.com/google/uuid"
)

type TransferController struct {
	conn *database.Connection
}

func NewTransferController(conn *database.Connection) *TransferController {
	return &TransferController{
		conn: conn,
	}
}

func (t TransferController) HandleCreateTransfer(w http.ResponseWriter, r *http.Request) {
	var transfer dto.TransferInput

	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		msg := "Error decoding body request: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	transferModel, err := transfer.ConvertInputToModel()

	if err != nil {
		msg := "Error converting to model: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	var fromAccountBalance float64
	var toAccountBalance float64

	if err = t.conn.Db.QueryRow("select balance from accounts where id = ?", transferModel.FromAccountID).Scan(&fromAccountBalance); err != nil {
		msg := "Erro scanning money from FromAccount"
		http.Error(w, msg, http.StatusNotAcceptable)
		return
	}

	if err = t.conn.Db.QueryRow("select balance from accounts where id = ?", transferModel.ToAccountID).Scan(&toAccountBalance); err != nil {
		msg := "Erro scanning money from toAccount"
		http.Error(w, msg, http.StatusNotAcceptable)
		return
	}

	if transferModel.Value > fromAccountBalance {
		msg := "Erro scanning money from account"
		http.Error(w, msg, http.StatusNotAcceptable)
		return
	}

	if _, err = t.conn.Db.Exec("update accounts set balance = ? where id = ?", fromAccountBalance-transferModel.Value, transferModel.FromAccountID); err != nil {
		msg := "Erro scanning money from account"
		http.Error(w, msg, http.StatusNotAcceptable)
		return
	}

	if _, err = t.conn.Db.Exec("update accounts set balance = ? where id = ?", toAccountBalance+transferModel.Value, transferModel.ToAccountID); err != nil {
		msg := "Erro scanning money from account"
		http.Error(w, msg, http.StatusNotAcceptable)
		return
	}

	stmt, err := t.conn.Db.Prepare("insert into transfers(id,to_account_id,from_account_id,value) values (?,?,?,?)")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if _, err = stmt.Query(transferModel.ID, transfer.ToAccountID, transferModel.FromAccountID, transferModel.Value); err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (t TransferController) HandleGetAllTransfer(w http.ResponseWriter, r *http.Request) {
	var transfers []struct {
		ID            uuid.UUID `json:"id"`
		ToAccountID   uuid.UUID `json:"to_account_id"`
		FromAccountID uuid.UUID `json:"from_account_id"`
		Value         float64   `json:"value"`
	}

	stmt, err := t.conn.Db.Prepare("select id,to_account_id,from_account_id,value from transfers")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query()

	if err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var transfer struct {
			ID            uuid.UUID `json:"id"`
			ToAccountID   uuid.UUID `json:"to_account_id"`
			FromAccountID uuid.UUID `json:"from_account_id"`
			Value         float64   `json:"value"`
		}

		if err = rows.Scan(&transfer.ID, &transfer.ToAccountID, &transfer.FromAccountID, &transfer.Value); err != nil {
			msg := "Error scanning database: " + err.Error()
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		transfers = append(transfers, transfer)
	}

	response, err := json.Marshal(&transfers)

	if err != nil {
		msg := "Error marshalling response: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (t TransferController) HandleGetAllTransferAccount(w http.ResponseWriter, r *http.Request) {

	var transfers []struct {
		ID          uuid.UUID `json:"id"`
		ToAccountID uuid.UUID `json:"to_account_id"`
		Value       float64   `json:"value"`
	}

	stmt, err := t.conn.Db.Prepare("select id,to_account_id,value from transfers where from_account_id = ?")

	if err != nil {
		msg := "Error preparing database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query(r.PathValue("id"))

	if err != nil {
		msg := "Error running database: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var transfer struct {
			ID          uuid.UUID `json:"id"`
			ToAccountID uuid.UUID `json:"to_account_id"`
			Value       float64   `json:"value"`
		}

		if err = rows.Scan(&transfer.ID, &transfer.ToAccountID, &transfer.Value); err != nil {
			msg := "Error scanning database: " + err.Error()
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		transfers = append(transfers, transfer)
	}

	response, err := json.Marshal(&transfers)

	if err != nil {
		msg := "Error marshalling response: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
