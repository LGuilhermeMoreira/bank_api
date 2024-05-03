package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID
	Owner   string
	Balance float64
}

func (a Account) Show() {
	fmt.Println(a.ID, a.Owner, a.Balance)
}
