package dto

import (
	"strconv"

	"github.com/LGuilhermeMoreira/bank_api/src/utils/models"
	"github.com/google/uuid"
)

type EntrieInput struct {
	AccountID string `json:"account_id"`
	Amount    string `json:"amount"`
}

func (d EntrieInput) ConvertInputToModel() (*models.EntrieModel, error) {
	amount, err := strconv.ParseFloat(d.Amount, 64)

	if err != nil {
		return nil, err
	}

	return &models.EntrieModel{
		ID:        uuid.New(),
		AccountID: uuid.MustParse(d.AccountID),
		Amount:    amount,
	}, nil
}
