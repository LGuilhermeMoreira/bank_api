package models

import (
	"fmt"

	"github.com/google/uuid"
)

type LoginAccountModel struct {
	ID           uuid.UUID
	UserMail     string
	UserPassword string
}

func (l LoginAccountModel) Show() {
	fmt.Println(l.ID, l.UserMail, l.UserPassword)
}
