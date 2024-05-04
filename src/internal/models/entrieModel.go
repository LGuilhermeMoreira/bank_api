package models

import (
	"fmt"

	"github.com/google/uuid"
)

type EntrieModel struct {
	ID         uuid.UUID
	Account_ID uuid.UUID
	Amount     float64
}

func (e EntrieModel) Show() {
	fmt.Printf("%v %v", e.Account_ID, e.Amount)
}
