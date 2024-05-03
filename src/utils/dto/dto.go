package dto

import (
	"log"
	"strconv"

	"github.com/LGuilhermeMoreira/bank_api/src/utils/models"
	"github.com/google/uuid"
)

type DtoAccount struct {
	Owner   string `json:"owner"`
	Balance string `json:"balance"`
}

type DtoEntrie struct {
	Account_ID string `json:"account_id"`
	Amount     string `json:"amount"`
}

type DtoTrasnfer struct {
	ToAccountID   string `json:"to_account_id"`
	FromAccountID string `json:"from_account_id"`
	Amount        string `json:"amount"`
}

func (d DtoTrasnfer) ConvertToTransferModel() *models.Transfer {

	return &models.Transfer{
		ID:            uuid.New(),
		ToAccountID:   uuid.MustParse(d.ToAccountID),
		FromAccountID: uuid.MustParse(d.FromAccountID),
	}
}

func (d DtoEntrie) ConvertToModelEntrie() *models.Entrie {
	amount, err := strconv.ParseFloat(d.Amount, 64)

	if err != nil {
		log.Fatalf("Error parsing string to float\n")
		return &models.Entrie{}
	}

	return &models.Entrie{
		Account_ID: uuid.MustParse(d.Account_ID),
		Amount:     amount,
	}
}

func (d DtoAccount) ConvertToModelAccount() *models.Account {
	balance, err := strconv.ParseFloat(d.Balance, 64)

	if err != nil {
		log.Fatalf("Error parsing string to float\n")
		return &models.Account{}
	}

	return &models.Account{
		ID:      uuid.New(),
		Owner:   d.Owner,
		Balance: balance,
	}
}
