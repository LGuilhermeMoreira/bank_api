package models

import (
	"fmt"

	"github.com/google/uuid"
)

type TransferModel struct {
	ID            uuid.UUID
	ToAccountID   uuid.UUID
	FromAccountID uuid.UUID
	Value         float64
}

func (t TransferModel) Show() {
	fmt.Println(t.ID, t.ToAccountID, t.FromAccountID, t.Value)
}
