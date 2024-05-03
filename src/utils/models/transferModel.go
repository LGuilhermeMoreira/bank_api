package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Transfer struct {
	ID            uuid.UUID
	ToAccountID   uuid.UUID
	FromAccountID uuid.UUID
}

func (t Transfer) Show() {
	fmt.Println(t.ID, t.ToAccountID, t.FromAccountID)
}
