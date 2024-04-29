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

func (d DtoAccount) ConvertToModelAccount() *models.Account {
	balance, err := strconv.ParseFloat(d.Balance, 64)

	if err != nil {
		balance = 0.0
		log.Fatalf("Incapaz de fazer a conversão de string para float\n")
	}

	return &models.Account{
		ID:      uuid.New(),
		Owner:   d.Owner,
		Balance: balance,
	}
}
