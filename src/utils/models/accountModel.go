package models

import (
	"fmt"

	"github.com/google/uuid"
)

type AccountModel struct {
	ID             uuid.UUID
	Owner          string
	Balance        float64
	LoginAccountID uuid.UUID
}

func (a AccountModel) Show() {
	fmt.Println(a.ID, a.Owner, a.Balance)
}
