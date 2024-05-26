package dto

import (
	"strconv"

	"github.com/LGuilhermeMoreira/bank_api/internal/models"
	"github.com/google/uuid"
)

type AccountInput struct {
	LoginAccountID string `json:"login_account_id"`
	Owner          string `json:"owner"`
	Balance        string `json:"balance"`
}

func (d AccountInput) ConvertInputToModel() (*models.AccountModel, error) {

	balance, err := strconv.ParseFloat(d.Balance, 64)

	if err != nil {
		return nil, err
	}

	return &models.AccountModel{
		ID:             uuid.New(),
		LoginAccountID: uuid.MustParse(d.LoginAccountID),
		Balance:        balance,
		Owner:          d.Owner,
	}, nil
}
