package models

import (
	"fmt"

	"github.com/google/uuid"
)

type EntrieModel struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	Amount    float64
}

func (e EntrieModel) Show() {
	fmt.Printf("%v %v", e.AccountID, e.Amount)
}
