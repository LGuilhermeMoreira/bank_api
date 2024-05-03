package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Entrie struct {
	Account_ID uuid.UUID
	Amount     float64
}

func (e Entrie) Show() {
	fmt.Printf("%v %v", e.Account_ID, e.Amount)
}
