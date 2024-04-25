package dto

import "github.com/google/uuid"

type DtoAccount struct {
	ID      uuid.UUID `json:"id"`
	Owner   string    `json:"owner"`
	Balance string    `json:"balance"`
}
