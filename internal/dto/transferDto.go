package dto

import (
	"strconv"

	"github.com/LGuilhermeMoreira/bank_api/internal/models"
	"github.com/google/uuid"
)

type TransferInput struct {
	ToAccountID   string `json:"to_account_id"`
	FromAccountID string `json:"from_account_id"`
	Value         string `json:"value"`
}

func (d TransferInput) ConvertInputToModel() (*models.TransferModel, error) {
	value, err := strconv.ParseFloat(d.Value, 64)

	if err != nil {
		return nil, err
	}

	return &models.TransferModel{
		ID:            uuid.New(),
		ToAccountID:   uuid.MustParse(d.ToAccountID),
		FromAccountID: uuid.MustParse(d.FromAccountID),
		Value:         value,
	}, nil
}
