package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID `json:"id"`
	Owner   string    `json:"owner"`
	Balance float64   `json:"balance"`
}

func (a Account) Show() {
	fmt.Println(a.ID, a.Owner, a.Balance)
}
